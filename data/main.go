package data

import (
	"DataWall/cassandra"
	"fmt"
)

type DeviceSetArr []cassandra.Device

// Not pointer because if system slows down and takes longer than 20 seconds to serialize we suddenly run on new data instead of finishing the old batch
func SerializeData(DevicesSet *[]cassandra.Device) {
	var deviceSet = DevicesSet
	fmt.Println(len(*deviceSet))
	//GetDisconnects(deviceSet)
	fmt.Println(len(*deviceSet))
}
