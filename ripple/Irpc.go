package main

import (
	"DataWall/cassandra"
	"DataWall/distributer"
)

type Node interface {
	Multiply(args *distributer.Args, reply *[]cassandra.Device) error
}