package bus

import (
	"net"
	"log"
    "golem/messages/command"
)

const (
    busAddress = "/tmp/golem_bus.sock"
)

/*
EventBusListener Will be the backend bus which messages will be passed over
Using a unix socket to make it more efficient.  Does
not have any TCP overhead.
*/
func EventBusListener() {
	l, err := net.Listen("unix", busAddress)
    if err != nil {
        log.Fatal("listen error:", err)
    }
	
	for {
        fd, err := l.Accept()
        if err != nil {
            log.Fatal("accept error:", err)
        }

        go processConnection(fd)
    }
	
	
}

func processConnection(c net.Conn) {
	
}

/*
EventBusClient Will push commands onto the event bus.
*/
func EventBusClient(command Command) {
    
}
