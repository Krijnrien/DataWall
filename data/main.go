package data

import (
	"DataWall/cassandra"
	"fmt"
	//"DataWall/distributer"
	//"DataWall/distributer"
)

type DeviceSetArr []cassandra.Device

// Not pointer because if system slows down and takes longer than 20 seconds to serialize we suddenly run on new data instead of finishing the old batch
func SerializeData(DevicesSet *[]cassandra.Device) {
	var apfel *[]cassandra.Device = DevicesSet
	fmt.Println(len(*apfel))
	//GetDisconnects(apfel)
	//distributer.CreateNodes(apfel)
	fmt.Println(len(*apfel))
}
