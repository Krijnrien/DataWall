package data

import (
	"fmt"
	"DataWall/cassandra"
)


func GetDisconnects(DevicesSet *[]cassandra.Device) {
	for _, device := range *DevicesSet {
	fmt.Println(device.Hash)

	}
}
