package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"DataWall/cassandra"

	log "github.com/sirupsen/logrus" // Logging library
)

func getAllLocations(writer http.ResponseWriter, request *http.Request) {
	log.Debug("Fetching and attempting to serve all devices!")
	// Get limit for maximum records to fetch
	var limit int = getLimit(&request.Header)
	// Fetch list of all devices from database with given limit
	data := cassandra.GetDevices(limit)

	// Serialize list to JSON format
	b, serializeErr := json.Marshal(data)

	// Handle possible error
	if serializeErr != nil {
		log.WithFields(log.Fields{
			"Error": serializeErr,
		}).Error("Failed to serialize devices list to JSON!")
		// Write error to webpage
		// TODO writing error to user is potentially dangerous!
		fmt.Fprint(writer, defaultResponse)
	} else {
		// Write devices endpoint on webpage
		fmt.Fprintf(writer, string(b))
		log.Debug("Successfully fetched devices list and printed on screen!")
	}
}
