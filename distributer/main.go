package distributer

import (
	//"DataWall/config"
	"net"
	"fmt"
	"net/rpc"
	log "github.com/sirupsen/logrus"
	"DataWall/cassandra"
	//"DataWall/config"
)

func CreateNodes(Devices *[]cassandra.Device) {
	// CFG host string ins't properly formatted, creates wrong address.
	// cfg := *config.Get()
	// cfg.Ripple.IP+":"+fmt.Sprint(cfg.Ripple.Port)
	conn, err := net.Dial("tcp", "localhost:8084")
	if err != nil {
		log.Fatal("Connecting: ", err)
	}

	fmt.Println("testing")
	node := &Node{client: rpc.NewClient(conn)}
	fmt.Println("" + fmt.Sprint(node.Multiply(Devices)))
}

func (t *Node) Multiply(Devices *[]cassandra.Device) []cassandra.Device {
	args := &Args{Devices}
	var reply []cassandra.Device
	err := t.client.Call("Ripple.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	return reply
}
