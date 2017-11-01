package main

import (
	"DataWall/api"
	"DataWall/polling"
)

func main() {
	// Starting device location polling first.
	// If the server crashes it should start delivering data to waiting clients asap.
	go polling.EndpointPolling()

	// Start the API
	api.StartServer()
}
