package master

import (
	"golem-server/net/bus"
)

/*
StartMaster starts the Master processes
*/
func StartMaster(commPort int, dataPort int, bindIP string, role string) {
	go bus.EventBusListener()
	bus.SocketListener(bindIP, commPort, role)
}