package main

import "strconv"

const IpAddress = "127.0.0.1"
const Keyspace = "data"
const Port = 3000
const Logging = true
const ApiProtocol = "https://"
const ApiDomain = "api.fhict.nl"
const ApiDevicesPath = "/location/devices"

var ServeAddres = ":" + strconv.Itoa(Port)
var Token = ""
