package main

import (
	"log"
	"net"
	"net/rpc"
)

func registerNode(server *rpc.Server, node Node) {
	// registers Node interface with name `Ripple`.
	server.RegisterName("Ripple", node)
}

func main() {
	//Creating an instance of struct which implement Node interface
	noder := new(Rpc)

	// Register a new rpc server
	server := rpc.NewServer()
	registerNode(server, noder)

	// Listen for incoming tcp packets on specified port.
	l, e := net.Listen("tcp", ":8084")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	server.Accept(l)
}
