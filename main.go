package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"flag"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"strconv"
)

const interval = 20000 * time.Millisecond // Time with 20 seconds interval

var cache = "[]"

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
	flag.StringVar(&Token, "token", "", "The access token that will be used to contact the fontys api.")
	flag.BoolVar(&UseDatabase, "use-database", false, "Whether the application should store all fontys api responses in the database.")
	flag.Parse()

	// Check if the token is valid
	if !ValidToken(Token) {
		log.Error("Missing required parameter token. For help see -h.")
		os.Exit(1)
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
	fmt.Fprintf(w, cache)
}

/**
 * The API server that handles the requests.
 */
func apiServer(_ time.Time) {
	http.HandleFunc("/devices", devicesEndpoint)
	http.ListenAndServe(":"+strconv.Itoa(Port), nil)
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

	// Retrieve Token from Config and set in proper struct
	tokenSource := &TokenSource{
		AccessToken: Token,
	}

	// Create oauth2 client with inserted token to proceed GET request and read the response
	resp, _ := oauth2.NewClient(oauth2.NoContext, tokenSource).Get(devicesEndpointUrl)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	// Serialize JSON response to device struct.
	var devices *[]Device
	err := json.Unmarshal([]byte(string(body)), &devices)
	if err != nil {
		log.Error("Could not serialize JSON response to device struct. Is your token up to date?")
		os.Exit(0)
	}

	// Cache the response
	bytes, _ := json.MarshalIndent(devices, "", " ")
	cache = string(bytes)

	if UseDatabase {
		// Insert the response in the database
		go InsertDevices(*devices)
	}
}
