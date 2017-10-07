package main

import (
	"DataWall/cassandra"
	"DataWall/distributer"
)

type Rpc int

func (t *Rpc) Multiply(args *distributer.Args, reply *[]cassandra.Device) error {
	*reply = *args.A
	return nil
}