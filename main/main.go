package main

import (
	"time"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"net/http"

	"golang.org/x/oauth2"
	log "github.com/sirupsen/logrus" // Logging library
)

const interval = 20000 * time.Millisecond // Time with 20 seconds interval

var Devices *[]Device
var cache = "Waiting for first request"

func main() {
	// Check if the token is valid
	if len(Token) == 0 {
		log.WithFields(log.Fields{
			"Error": "Invalid token!",
		}).Error()
		os.Exit(0)
	}

	log.Info("Starting data-gatherer application!")

	// Refresh the devices list every `interval`.
	go doEvery(interval, refreshDevices)

	// Refresh now.
	go refreshDevices(time.Now())

	// Start the api server.
	apiServer(time.Now())
}

/**
* Timer to repeat func every given amount of time.
* @param interval in whole seconds.
* @param function name to repeat every interval tick
*/
func doEvery(interval time.Duration, repeatFunction func(time.Time)) {
	for currentTime := range time.Tick(interval) {
		repeatFunction(currentTime)
	}
}

/**
 * The handler for the devices endpoint
 */
func devicesEndpoint(w http.ResponseWriter, h *http.Request) {
	fmt.Fprintf(w, cache)
}

/**
 * The API server that handles the requests.
 */
func apiServer(currentTime time.Time) {
	http.HandleFunc("/devices", devicesEndpoint)
	http.ListenAndServe(ServeAddres, nil)
}

/**
 * The function that refreshes the cache and gets the response from the fontys api.
 */
func refreshDevices(currentTime time.Time) {
	log.WithFields(log.Fields{
		"Start time": time.Now(),
	}).Debug("Retrieving data from Fontys API")

	// Retrieve configuration for Fontys Devices API url
	devicesEndpointUrl := ApiProtocol + ApiDomain + ApiDevicesPath // Fontys endpoint url

	// Retrieve Token from Config and set in proper struct
	tokenSource := &TokenSource{
		AccessToken: Token,
	}

	// Create oauth2 client with inserted token to proceed GET request and read the response
	resp, _ := oauth2.NewClient(oauth2.NoContext, tokenSource).Get(devicesEndpointUrl)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	// Serialize JSON response to device struct.
	err := json.Unmarshal([]byte(string(body)), &Devices)
	if err != nil {
		log.WithFields(log.Fields{
			"Error":   err.Error(),
			"Comment": "Could not serialize JSON response to device struct. Is your token up to date?",
		}).Error()
	}

	// Cache the response
	bytes, _ := json.MarshalIndent(Devices, "", " ")
	cache = string(bytes)
}
