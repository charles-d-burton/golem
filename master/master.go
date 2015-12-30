package master

import (
	"golem/net/bus"
)

/*
StartMaster starts the Master processes
*/
func StartMaster(commPort int, dataPort int, bindIP string, role string) {
	go bus.EventBusListener()
	go bus.SocketListener(bindIP, dataPort, "data")
	bus.SocketListener(bindIP, commPort, "comm")
}