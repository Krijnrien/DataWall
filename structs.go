package main

import "time"

// Struct to use in the authentication process
type TokenSource struct {
	AccessToken string `json:"access_token"`
}

// Struct for Fontys API return result
type Device struct {
	X         float32   `json:"x"`        // X axis of device location. Generated from single geolocation.
	Y         float32   `json:"y"`        // Y axis of device location. Generated from single geolocation.
	Z         int8      `json:"z"`        // Z-index determines floor with one digit. Not making unsigned as future might make basement negative?
	UserType  uint      `json:"userType"` // User type associated with device.
	Hash      string    `json:"hash"`     // UserHash from fontys is a combination of Ip address and datetime (presumed of IP handout).
	CreatedAt time.Time // CreatedAt field is added and assigned instantly at return.
}
