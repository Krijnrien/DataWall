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
	"net/http"
	"fmt"
)

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

var Devices *[]cassandra.Device
const interval time.Duration = 20000 * time.Millisecond // Time with 20 seconds interval

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

func serveApi(w http.ResponseWriter, h *http.Request) {

	log.WithFields(log.Fields{
		"Start time": time.Now(),
	}).Debug("Retrieving data from Fontys API")

	// Retrieve configuration for Fontys Devices API url
	cfg := *config.Get()
	devicesEndpointUrl := cfg.ApiProtocol + cfg.ApiDomain + cfg.ApiDevicesPath // Fontys endpoint url

	// TODO Should this variable be predefined?

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

	structuredjson, err := json.MarshalIndent(Devices,"", " ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(structuredjson))
	fmt.Println("end")

}
/** getDevicesLocationsData
 * Get Fontys authentication token. Connect to devices location endpoint, read & serialize response.
 * currentTime //TODO Unused parameter? NO!
 */
func getDevicesLocationsData(currentTime time.Time) {
	http.HandleFunc("/array",serveApi)
	http.ListenAndServe(":3000",nil)
}
