package controller

import (
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus" // Logging library
	"DataWall/api/cache"
	"path/filepath"
)

const defaultLimit int = 50         // Default limit if none could be found in header.
const headerLimit string = "Limit"  // Field to search for in header.
const defaultResponse string = "{}" // Empty response body predefined.

/** RegisterEndPoints
 * Set all endpoints of webserver to listen and serve to.
 */
var storage Storage

type Storage interface {
	Get(key string) []byte
	Set(key string, content []byte, duration time.Duration)
}

func RegisterEndPoints() {
	// Instantiate type of cache storage. Using memory cache.
	storage = cache.NewMemCache()
	// GET: /
	http.Handle("/", caching("20s", root))
	// GET: /data
	http.HandleFunc("/data", getAllLocations)
}

/** getLimit
 * Gets header limit from HTTP requests and returns this limit. If no limit header cannot be retrieved it returns programmatically default set limit.
 * @param h:
 * @return limit : integer of maximum amount of records to return
 */
func getLimit(header *http.Header) int {
	// Get header limit from the HTTP request
	var headerLimit string = header.Get(headerLimit)
	log.WithFields(log.Fields{
		"limit": headerLimit,
	}).Debug("Header limit of HTTP request")
	limit, err := strconv.Atoi(headerLimit)
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err.Error(),
		}).Error("Could not retrieve header limit from HTTP request!")
		// Return default set limit if none could be retrieved from header.
		return defaultLimit
	}
	return limit
}

/** root
 * Serves static HTML file displaying endpoints by default.
 * @param writer
 * @param request
 */
func root(writer http.ResponseWriter, request *http.Request) {
	// Set header of request
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Get file path of static html file to serve as root.
	absPath, _ := filepath.Abs("../api/static/home.html")
	// Serve (return) static html file to request
	http.ServeFile(writer, request, absPath)
}
