package containers

import (
	"net"
)

/*
Container for a host object
Has a channel with a network connection.
We'll use the channel to send messages over
the network to the remote host.
*/
type Host struct {
	RemoteHost net.Conn
	RemoteAddr net.Addr
	HostName string
	Port int
	Message chan string
}