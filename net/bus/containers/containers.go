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
	UUID string
	Message chan string
	*sync.RWMutex
}

/*
Job... top level container for jobs
*/
type Job interface {
	Lock() sync.RWMutex
	JobID() string
}

/*
Mesh...
Struct to contain the mesh relay mapping
*/
type Mesh struct {
	Peers []string
	Nets []string
	
	
}

/* Peer ... container
to hold information about Peers
for distribution to all clients
to inform about the mesh
*/
type Peer struct {
	Host string
	Net string
	
}