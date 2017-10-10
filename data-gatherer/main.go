package polling

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"DataWall/config"
	"DataWall/data"
	log "github.com/sirupsen/logrus" // Logging library
	"golang.org/x/oauth2"            // Authentication library
	"DataWall/cassandra"
)
const interval time.Duration = 20 * time.Second // Time with 20 seconds interval
//TODO
/** endpointPolling
*
*/
func EndpointPolling() {
	log.Info("Starting data-gatherer application!")
	getDevicesLocationsData(time.Now())

	for{
		getDevicesLocationsData()
		time.Sleep(interval)
	}

}

var Devices *[]cassandra.Device




/** getDevicesLocationsData
 * Get Fontys authentication token. Connect to devices location endpoint, read & serialize response.
*/
func getDevicesLocationsData() {

	log.WithFields(log.Fields{
		"Start time": time.Now(),
	}).Debug("Retrieving data from Fontys API")

	// Retrieve configuration for Fontys Devices API url
	cfg := *config.Get()
	devicesEndpointUrl := cfg.ApiProtocol + cfg.ApiDomain + cfg.ApiDevicesPath // Fontys endpoint url


	// Retrieve Token from Config and set in proper struct
	tokenSource := &TokenSource{
		AccessToken: config.Get().Token,
	}

	// TODO DEPRECATED? NO!
	// Create oauth2 client with inserted token to proceed GET request and read the response
	resp, _ := oauth2.NewClient(oauth2.NoContext, tokenSource).Get(devicesEndpointUrl)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

//	DevicesSet.Mutex.RLock()
	// Serialize JSON response to device struct.
	err := json.Unmarshal([]byte(string(body)), &Devices)
	if err != nil {
		// TODO Handle error more gracefully!
		log.WithFields(log.Fields{
			"End time": err.Error(),
		}).Error("Could not serialize JSON response to device struct")
	}
	//DevicesSet.Mutex.RUnlock()

	go data.SerializeData(Devices)

	log.WithFields(log.Fields{
		"End time": time.Now(),
	}).Debug("Finished retrieving data from Fontys API")
}
