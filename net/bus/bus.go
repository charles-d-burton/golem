package bus

/*
This package is intended to start the socket connections that
the system uses. It also will bind the clients and group the necessary connections
together.  This means that when a client connects it will be connected to the control
bus chan, will be given a message chan, and a chan to send file data through.  Each of these
chans will be connected to a socket.
*/

import (
	"net"
	"log"
    "github.com/golem/net/bus/containers"
    "strconv"
)

const (
    busAddress = "/tmp/golem_bus.sock"
)

var hostPool []containers.Host

/*
EventBusListener Will be the backend bus which messages will be passed over
Using a unix socket to make it more efficient.  Does
not have any TCP overhead.
*/
func EventBusListener() {
    log.Println("Starting Listener")
	l, err := net.Listen("unix", busAddress)
    if err != nil {
        log.Fatal("listen error:", err)
        panic(err)
    }
	
	for {
        fd, err := l.Accept()
        if err != nil {
            log.Fatal("accept error:", err)
            //panic(err)
        }

        go processConnection(fd)
    }
    log.Println("Exiting Listener")	
}

/*
This opens the communication CommSocketListener
peons will connect to and use this socket
//TODO:  Connect the socket bus here with the control bus
*/
func SocketListener(ip string, port int, role string) {
    listener := initListener(ip, port)
    if role == "comm" {
        startCommPort(listener)
    } else {
        go startDataPort(listener)
    }
}

/*
Creates the socket listener
*/
func initListener(ip string, port int) (net.Listener) {
    log.Println("Listening for Comm on socket: ", strconv.Itoa(port))
    l, err := net.Listen("tcp", ip + ":" + strconv.Itoa(port))
    if err != nil {
        log.Fatal(err)
        panic(err)
    }
    return l
}

func startCommPort(l net.Listener) {
    for {
        fd, err := l.Accept()
        if err != nil {
            log.Println("Something went wrong: ", err)
        }
        go processConnection(fd)
    }
}

func startDataPort(l net.Listener) {
    for {
        fd, err := l.Accept()
        if err != nil {
            log.Println("Something went awry: ", err)
        }
        go attachDataPort(fd)
    }    
}

func processConnection(c net.Conn) {
    ip, _, err := net.SplitHostPort(c.RemoteAddr().String())
    if err != nil {
        log.Println(err)
        return
    }
    for _, peon := range hostPool {
        peonIP, _, err := net.SplitHostPort(peon.RemoteHost.RemoteAddr().String())
        if err != nil {
            log.Println("Connecetion error")
            return
        }
        if peonIP == ip {
            log.Println("Already accepted a host with this IP address\n", "Closing connections")
            c.Write([]byte("Already Accepted a connection from this host\n"))
            c.Write([]byte("Closing..."))
            defer c.Close()
            return
        }
    }
    log.Println("Processing Connection from: ", ip)
	hostChannel := make (chan string)
    host := containers.Host{c, nil, "test", 10000, hostChannel}
    hostPool = append(hostPool, host)
    log.Println("Finished Processing Connection")
}

func attachDataPort(c net.Conn) {
    ip, _, err := net.SplitHostPort(c.RemoteAddr().String())
    if err != nil {
        log.Println("Connection Error")
        return
    }
    log.Println("Initializing Data port from: ", ip)
    for _, peon := range hostPool {
        peonIP, _, err := net.SplitHostPort(peon.RemoteHost.RemoteAddr().String())
        if err != nil {
            log.Println("Connection error")
            return
        }
        if peonIP == ip {
            peon.DataPort = c
            break //Found a match, no need to continue further
        }
    }
}

/*
FindClients this function will search the structs of associated peons
and return a slice of all the matching peons.
*/
func FindClients(pattern string) ([]containers.Host,  error) {
    if hostPool != nil {
        for _, value := range hostPool {
            name := value.HostName
            log.Println("Hostname of Search: " + name)
        }
    }
    return nil, nil
}

/*
EventBusClient Will push commands onto the event bus.
func EventBusClient(command Command) {
    
}*/
