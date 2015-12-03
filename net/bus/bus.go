package bus

import (
	"net"
	"log"
    "golem-server/net/bus/containers"
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
func CommSocketListener(ip string, port int, role string) {
    log.Println("Listening for Comm on socket: " + strconv.Itoa(port))
    if ip == "" {
        ip = "localhost"
    }
    if port == 0 {
        port = 10000
    }
    l, err := net.Listen("tcp", ip + ":" + strconv.Itoa(port))
    if err != nil {
        log.Fatal(err)
        panic(err)
    }
    for {
        fd, err := l.Accept()
        if err != nil {
            log.Fatal(err)
            //panic(err)
        }
        go processConnection(fd)
    }
}

func processConnection(c net.Conn) {
    log.Println("Processing Connection")
	hostChannel := make (chan string)
    host := containers.Host{c, "test", 10000, hostChannel}
    hostPool = append(hostPool, host)
    log.Println("Finished Processing Connection")
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
