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
    "golem-server/net/bus/containers"
    "strconv"
    "errors"
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
    
    peon, err := findConnection(c)
    if err != nil {
        log.Println("Connection errored out")
        return
    }
    if peon != nil {
        remote, _, _ := net.SplitHostPort(peon.RemoteHost.RemoteAddr().String())
        log.Println("Found a host already connected with that address", remote)
        c.Write([]byte("Already Accepted a connection from this host\n"))
        defer c.Close()
        return            
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
    peonConn, err := findConnection(c)
    if err != nil {
        log.Println("Connection errored out")
        defer c.Close()
        return
    }
    if peonConn != nil {
        log.Println("Initializing Data port from: ", ip)
        peonConn.DataPort = c
    } else {
        log.Println("Did not find a communication port")
        c.Write([]byte("Did not find a communication port for you, closing"))
        defer c.Close()
    }
}

func findConnection(c net.Conn) (*containers.Host, error) {
    ip, _, err := net.SplitHostPort(c.RemoteAddr().String())
    if err != nil {
        log.Println("Unable to resolve ip address from connection")
        return nil, errors.New("Unable to resolve ip address from connections")
    }
    for _, peon := range hostPool {
        peonIP, _, err := net.SplitHostPort(peon.RemoteHost.RemoteAddr().String())
        if err != nil {
            log.Println("Unable to resolve peon ip address")
            return nil, errors.New("Unable to resolve ip address from connection")
        }
        if peonIP == ip {
            return &peon, nil
        }
    }
    return nil, nil
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
