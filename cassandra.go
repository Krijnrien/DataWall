package main

import (
	"time"
	"fmt"
	"sync"

	log "github.com/sirupsen/logrus" // Logging errors
	"github.com/gocql/gocql"
)

var session *gocql.Session
var once sync.Once

/** GetDevices
 * Requests devices location from database with given maximum record limit.
 * @param limit: Amount of rows you want to limit your request to. Must be positive.
 * @return List of device locations by struct Locations.
 */
func GetDevices(limit int) []Device {
	var locList []Device          // Predeclared list of Device struct.
	m := map[string]interface{}{} // What is this?

	// Formatting database query. Selecting fields user_hash, createdAt, loc_x, loc_y, loc_y with maximum records given by limit variable.
	// TODO Test if limit uint is positive
	selectLocationsQuery := fmt.Sprintf(
		"%s %d",
		"SELECT user_hash, createdat, loc_x, loc_y, loc_z FROM locations LIMIT ",
		limit)
	// Create and assign variable of query iterator
	iterable := getSession().Query(selectLocationsQuery).Iter()

	// Iterate over DB query result and add to map.
	for iterable.MapScan(m) {
		locList = append(locList, Device{
			Hash:      m["user_hash"].(string),
			CreatedAt: m["createdat"].(time.Time),
			X:         m["loc_x"].(float32),
			Y:         m["loc_y"].(float32),
			Z:         m["loc_z"].(int8),
		})
		m = map[string]interface{}{}
	}
	return locList
}

/**
 * Creates new goroutine for each device in Device list to insert each device into database
 * @param devices. List of all serialized devices to insert.
 */
func InsertDevices(devices []Device) {
	for _, device := range devices {
		// Log for debugging start time of insert
		log.WithFields(log.Fields{
			"Start time": time.Now(),
		}).Debug("Start inserting list of devices into DB")
		// Create new Goroutine to execute insert method.
		go insert(device)

		// Log for debugging end time of insert
		log.WithFields(log.Fields{
			"End time": time.Now(),
		}).Debug("Finished inserting list of devices into DB")
	}
}

/**
 * Inserts struct parameter into database
 * @param device. Struct of device must have all fields populated or null will be inserted.
 */
func insert(device Device) {
	// Insert device struct into database
	if queryErr := getSession().Query(`INSERT INTO locations (loc_x, loc_y, loc_z, user_hash, createdat) VALUES \
										(?, ?, ?, ?, ?)`, device.X, device.Y, device.Z, device.Hash, time.Now()).Exec();
		queryErr != nil {
		log.WithFields(log.Fields{
			"Error": queryErr,
		}).Fatal("Failed inserting record into database! ")
	}
}

/**
 * Returns the cassandra session. If no session exists, create a new one.
 */
func getSession() *gocql.Session {
	// Block executes only once
	once.Do(func() {
		log.Info("Creating new cassandra instance")
		createSession()
	})
	return session
}

/**
 * Create a new cassandra session.
 */
func createSession() {
	var sessionErr error

	// Connect to cassandra database cluster and database keyspace (DB name).
	cluster := gocql.NewCluster(IpAddress)
	cluster.Keyspace = Keyspace

	// Set global getSession variable
	session, sessionErr = cluster.CreateSession()
	// Handling possible create getSession error
	if sessionErr != nil {
		// Logging & entering Panic state.
		log.WithFields(log.Fields{
			"Error": sessionErr.Error(),
		}).Panic("Failed to create Cassandra DB getSession!")
	}
}
