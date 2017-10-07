package data

import (
	"sync"
	"DataWall/cassandra"
)

type DevicesMutex struct {
	Devices cassandra.Device
	Mutex *sync.RWMutex
}