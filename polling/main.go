package polling

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"DataWall/cassandra"
	"DataWall/config"

	log "github.com/sirupsen/logrus" // Logging library
	"golang.org/x/oauth2"            // Authentication library
)

const interval time.Duration = 20000 * time.Millisecond // Time with 20 seconds interval

//TODO
/** endpointPolling
*
*/
func EndpointPolling() {
	log.Info("Starting data-gatherer application!")
	getDevicesLocationsData(time.Now())
	// call getDevicesLocationsData func every tick predefined by interval var.
	doEvery(interval, getDevicesLocationsData)
}

/** do Every //TODO Func name not clear enough
* Timer to repeat func every given amount of time. //TODO Does this have to be a seperate func? Can it not be recursive?
* @param interval in whole seconds.
8 @param function name to repeat every interval tick
*/
func doEvery(interval time.Duration, repeatFunction func(time.Time)) {
	for currentTime := range time.Tick(interval) {
		repeatFunction(currentTime)
	}
}

/** getDevicesLocationsData
 * Get Fontys authentication token. Connect to devices location endpoint, read & serialize response.
 * currenTime //TODO Unused parameter? NO!
 */
func getDevicesLocationsData(currentTime time.Time) {
	log.WithFields(log.Fields{
		"Start time": time.Now(),
	}).Debug("Retrieving data from Fontys API")

	// Retrieve configuration for Fontys Devices API url
	cfg := *config.Get()
	devicesEndpointUrl := cfg.ApiProtocol + cfg.ApiDomain + cfg.ApiDevicesPath // Fontys endpoint url

	// TODO Should this variable be predefined?
	var devices []cassandra.Device

	// Retrieve Token from Config and set in proper struct
	tokenSource := &TokenSource{
		AccessToken: config.Get().Token,
	}

	// TODO DEPRECATED? NO!
	// Create oauth2 client with inserted token to proceed GET request and read the response
	resp, _ := oauth2.NewClient(oauth2.NoContext, tokenSource).Get(devicesEndpointUrl)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	// Serialize JSON response to device struct.
	err := json.Unmarshal([]byte(string(body)), &devices)
	if err != nil {
		// TODO Handle error more gracefully!
		log.WithFields(log.Fields{
			"End time": err.Error(),
		}).Error("Could not serialize JSON response to device struct")
	}

	// Send devices list to insert func to be inserted into the DB.
	go cassandra.InsertDevices(devices)

	// Split data set into each floor
	// Send data to pusher

	log.WithFields(log.Fields{
		"End time": time.Now(),
	}).Debug("Finished retrieving data from Fontys API")
}