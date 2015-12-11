package master

import (
	"golem-server/net/bus"
)

/*
StartMaster starts the Master processes
*/
func StartMaster(commPort int, dataPort int) {
	go bus.EventBusListener()
	bus.SocketListener("localhost", 20000, "test")
}