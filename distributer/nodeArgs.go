package distributer

import (
	"DataWall/cassandra"
	"net/rpc"
)

type Args struct {
	A *[]cassandra.Device
}

type Node struct {
	client *rpc.Client
}
