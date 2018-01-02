package main

// Constants
const ApiProtocol = "https://"
const ApiDomain = "api.fhict.nl"
const ApiDevicesPath = "/location/devices"
const TokenUrl = "https://identity.fhict.nl/connect/token"

// Values determined by flags
var IpAddress string
var Keyspace string
var Port int
var UseDatabase bool
var ClientId string
var ClientSecret string
