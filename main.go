package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"flag"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"strconv"
	"strings"
)

const interval = 20000 * time.Millisecond // Time with 20 seconds interval

var Devices []Device

func main() {
	parseFlags()

	log.Info("Starting the datawall backend on port " + strconv.Itoa(Port))

	// Refresh the devices list every `interval`.
	go DoEvery(interval, refreshDevices)

	// Refresh now.
	go refreshDevices(time.Now())

	// Start the api server.
	apiServer(time.Now())
}

func parseFlags() {
	flag.StringVar(&IpAddress, "ip", "127.0.0.1", "The ip that the cassandra database will run on.")
	flag.StringVar(&Keyspace, "keyspace", "data", "The keyspace of the cassandra database.")
	flag.IntVar(&Port, "port", 3000, "The port that the api server will run on.")
	flag.BoolVar(&UseDatabase, "use-database", false, "Whether the application should store all fontys api responses in the database.")
	flag.StringVar(&ClientId, "client-id", "i361819-datawall", "ClientId to get authentication token")
	flag.StringVar(&ClientSecret, "client-secret", "Ryk4m-i07JhqHheSLYHvaN4eBqhw8WxZOo3sW7yr", "The hash to get authentication token")
	flag.Parse()

	// Check if the token is valid
	if !ValidToken(ClientId) {
		log.Error("Missing required parameter ClientId. For help see -h.")
		os.Exit(1)
	}

	// Check if the token is valid
	if !ValidToken(ClientSecret) {
		log.Error("Missing required parameter Client secret. For help see -h.")
		os.Exit(1)
		fmt.Println(ClientId)
	}

	// Check if the port is valid
	if !ValidPort(Port) {
		err := fmt.Sprintf("Port %d is already in use. For help see -h.", Port)
		log.Error(err)
		os.Exit(1)
	}
}

/**
 * The handler for the devices endpoint
 */
func devicesEndpoint(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(Devices)
}

/**
 * The API server that handles the requests.
 */
func apiServer(_ time.Time) {
	router := mux.NewRouter()
	router.HandleFunc("/devices", devicesEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(Port), router))
}

/**
 * The function that refreshes the cache and gets the response from the fontys api.
 */
func refreshDevices(_ time.Time) {
	log.WithFields(log.Fields{
		"Start time": time.Now(),
	}).Debug("Retrieving data from Fontys API")

	// Retrieve configuration for Fontys devices API url
	devicesEndpointUrl := ApiProtocol + ApiDomain + ApiDevicesPath

	// Values set for POST request
	v := url.Values{}
	v.Set("grant_type", "client_credentials")
	v.Set("scope", "fhict")
	v.Set("client_id", ClientId)
	v.Set("client_secret", ClientSecret)

	// Values.Encode() encodes the values into "URL encoded" form sorted by key
	s := v.Encode()

	// Create Post request
	req, err := http.NewRequest("POST", TokenUrl, strings.NewReader(s))
	if err != nil {
		log.Error("Could not create request")
		os.Exit(0)
	}

	// Set header of request
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Client executes request
	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		log.Error("Could not request server")
		os.Exit(0)
	}

	// Read body of request and close it
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	// Unmarshall body to the access token that will be used to contact the fontys api.
	var token *TokenSource
	err = json.Unmarshal([]byte(string(body)), &token)
	if err != nil {
		log.Error("Could not serialize JSON response to token struct. Are your credentials right?")
		os.Exit(0)
	}

	// Create oauth2 client with inserted token to proceed GET request and read the response
	resp, _ = oauth2.NewClient(oauth2.NoContext, token).Get(devicesEndpointUrl)
	body, _ = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	// Check if []Devices variable contains anything
	if len(Devices) > 0 {
		// Serialize JSON response to device struct.
		// And store in temporary devices variable
		var devices *[]Device
		err := json.Unmarshal([]byte(string(body)), &devices)
		if err != nil {
			log.Error("Could not serialize JSON response to device struct. Is your token up to date?")
			os.Exit(0)
		}
		// Use temporary devices variable to update main []Devices variable
		go update(*devices)

	} else {
		// Main []Devices variable does not contain anything, which means that this is the first run.
		err := json.Unmarshal([]byte(string(body)), &Devices)
		if err != nil {
			log.Error("Could not serialize JSON response to device struct. Is your token up to date?")
			os.Exit(0)
		}

		// sets times of creation
		for i := 0; i < len(Devices); i++ {
			Devices[i].CreatedAt = time.Now()
		}
	}

	if UseDatabase {
		// Insert the response in the database
		go InsertDevices(Devices)
	}
}

/**
 * The function that updates main []Devices variable
 * @param slice of new devices
 */
func update(newDevices []Device) {
	for i, oldDevice := range Devices {

		updateIndex := Contains(newDevices, oldDevice)
		if updateIndex >= 0 {
			// updates location of existing device
			oldDevice.X = newDevices[updateIndex].X
			oldDevice.Y = newDevices[updateIndex].Y
			oldDevice.Z = newDevices[updateIndex].Z
		} else {
			// removes logged out device
			if i == len(Devices) - 1 {
				Devices = Devices[:len(Devices)-1]
			} else{
				Devices = append(Devices[:i], Devices[i+1:]...)
			}
		}
	}

	for _, newDevice := range newDevices {

		if Contains(Devices, newDevice) == -1 {
			// adds new device
			newDevice.CreatedAt = time.Now()
			Devices = append(Devices, newDevice)
		}
	}
}

