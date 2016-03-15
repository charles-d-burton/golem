package containers

import (
	"net"
	"sync"
)

/*
Container for a host object
Has a channel with a network connection.
We'll use the channel to send messages over
the network to the remote host.
*/
type Host struct {
	RemoteHost net.Conn
	DataPort net.Conn
	HostName string
	Port int
	Message chan string
    PubKey [32]byte
	*sync.RWMutex
}

/*
Job... top level container for jobs
*/
type Job interface {
	Lock() sync.RWMutex
	JobID() string
}