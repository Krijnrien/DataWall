package api

import (
	"fmt"
	"time"
	"net/http"

	"DataWall/config"
	"DataWall/api/controller"

	log "github.com/sirupsen/logrus" // Logging library
)

/** API package main
 * Starts web server and listens on port for requests
 */
func Server() {
	// Set all routes for web server to listen to
	controller.RegisterEndPoints()
	// Set configuration var for Ip Address and port.
	cfg := *config.Get()

	// All logs from contextLogger will add ip and port as fields.
	contextLogger := log.WithFields(log.Fields{
		"Ip Address": cfg.IpAddress,
		"API port":   cfg.ApiPort,
		"Time":       time.Now(),
	})

	contextLogger.Info("Starting web server!")
	// Start web server with config port. No handler is set.
	http.ListenAndServe(fmt.Sprint(":", cfg.ApiPort), nil)
}