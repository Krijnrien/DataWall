package main

import (
	"DataWall/api"
	"DataWall/polling"

	//log "github.com/sirupsen/logrus" // Logging library
)

/** Datawall main
 * // TODO
 */
func main() {
	/**
	 * Starting device location polling first.
	 * If the server crashes it should asap start delivering data to waiting clients.
	 */
	go polling.EndpointPolling()

	/**
	 * Starting API listenandserve
	 */
	api.Server()
}
