package config

import (
	"encoding/json" // Serializing JSON configuration file to struct
	"os"            // Open configuration file
	"path/filepath" // Define config file path
	"sync"          // Config get() func only required to run once

	//log "github.com/sirupsen/logrus" // Logging errors
)

/**
 * Struct for global application configuration
 */
type Configuration struct {
	IpAddress string // IP addresses of cassandra database
	Keyspace  string // Keyspace (DB name) of cassandra database
	ApiPort   uint8  // The port on which the API will run. Always positive and below 65535
	Logging   bool   // Whether to display logs or not
	Token     string // Auth token from Fontys API

	ApiDomain      string // Domain of fontys API
	ApiProtocol    string // Protocol to reach Fontys API
	ApiDevicesPath string // Path from ApiDomain to Fontys ConnectedDevices API

	Ripple struct {
		IP   string `json:"ip"`
		Port int    `json:"port"`
	} `json:"cluster"`
}

const configPath = "../DataWall/config/config.json"

var conf *Configuration // Predeclared global configuration struct, accessible by Get() func returning conf pointer.
var once sync.Once
var s string = `{
			"IpAddress": "127.0.0.1",
				"Keyspace": "data",
				"ApiPort": 8081,
				"Token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6ImdyQWk2cnJRU0JiVVItY01ZOHpRTHE2aGdVQSIsImtpZCI6ImdyQWk2cnJRU0JiVVItY01ZOHpRTHE2aGdVQSJ9.eyJpc3MiOiJodHRwczovL2lkZW50aXR5LmZoaWN0Lm5sIiwiYXVkIjoiaHR0cHM6Ly9pZGVudGl0eS5maGljdC5ubC9yZXNvdXJjZXMiLCJleHAiOjE1MDk1MzQ5MzIsIm5iZiI6MTUwOTUyNzczMiwiY2xpZW50X2lkIjoiYXBpLWNsaWVudCIsInVybjpubC5maGljdDp0cnVzdGVkX2NsaWVudCI6InRydWUiLCJzY29wZSI6WyJvcGVuaWQiLCJwcm9maWxlIiwiZW1haWwiLCJmaGljdCIsImZoaWN0X3BlcnNvbmFsIiwiZmhpY3RfbG9jYXRpb24iXSwic3ViIjoiY2ViNTA3ZDMtZGMxOC00NDdiLTkxNTEtZjNiNDZlOGRhMTkzIiwiYXV0aF90aW1lIjoxNTA5NTI3NzMxLCJpZHAiOiJmaGljdC1zc28iLCJyb2xlIjpbInVzZXIiLCJzdHVkZW50Il0sInVwbiI6IkkzNTg3MTdAZmhpY3QubmwiLCJuYW1lIjoiU2F2b3YsTWFydGluIE0uWi4iLCJlbWFpbCI6Im1hcnRpbi5zYXZvdkBzdHVkZW50LmZvbnR5cy5ubCIsInVybjpubC5maGljdDpzY2hlZHVsZSI6ImNsYXNzfERlbHRhIC8gRWkzUzQgLyBTTTQxIiwiZm9udHlzX3VwbiI6IjM1ODcxN0BzdHVkZW50LmZvbnR5cy5ubCIsImFtciI6WyJleHRlcm5hbCJdfQ.LTegINiv-yOyHfwRwi-LXb0Yx_lv_mqeuzKLc1BXzKwP6UfMjgJa5E62HCBn5yIHUyeb-I6pFoJoRW6MXk28ZDo3VG-A-K4IoJT43bSuh4zZ5kpYssmHNjp_t6xcrDrv0mfga4RULGj5BSeaCvRVne5Coa4r8Ik-I-TestALHTW-QnIdZY5-6oa3WlHg5nI3lV0at_rk5bMQwNOqWfR_q6RmN-aFgmWT9HgsDzIwpXy6geZrLR_tn0Jm8SagrRj3zvPOJOG_SnCBxFJG2MX6OOeJ2QBZw3244BmQsheBcj6sD95faX0yuqtvYD4GcENCuPbpK5WlPhhPOdQF392dhg",
				"ApiProtocol": "https://",
				"ApiDomain": "api.fhict.nl",
				"ApiDevicesPath": "/location/devices",
				"Ripple": {
			"ip": "localhost",
			"port": 8084
			}
		}`

/** Config.Get
 * Opens json config file. Decodes content to json into var conf of Configuration struct.
 * @return Configuration pointer
 */
func Get() *Configuration {
	once.Do(func() {
		//// Set file path to be serialized
		//absPath, _ := filepath.Abs(configPath)
		//
		//// All logs by contextLogger now include file path
		//contextLogger := log.WithFields(log.Fields{
		//	"path": absPath,
		//})
		//contextLogger.Info("Retrieving and setting configuration from file...")
		//
		//// Open & close file of given path
		//file, fileErr := os.Open(absPath)
		//defer file.Close()
		//
		//// Handling possible errors
		//if fileErr != nil {
		//	log.WithFields(log.Fields{
		//		"Error": fileErr.Error(),
		//	}).Fatal("Failed opening file!")
		//}

		// Decoding (serializing) JSON content of file to Configuration struct
		json.Unmarshal([]byte(s), &conf)

		//contextLogger.Info("Configuration initialized!")
	})

	return conf
}

func fileExists() bool {
	absPath, _ := filepath.Abs(configPath)
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return false
	}
	return true
}
