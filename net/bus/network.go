package bus

import (
    "strconv"
    "log"
    "net"
    "crypto/tls"
    "golem/secure"
    
)

/*
This opens the communication CommSocketListener
peons will connect to and use this socket
//TODO:  Connect the socket bus here with the control bus
*/
func SocketListener(ip string, port int, role string) {
    //tlsCfg.InsecureSkipVerify = true    
    //tlsListener, err := tls.Listen("tcp4", ip + ":" + strconv.Itoa(port), &tlsCfg)

    
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
    cert, err := tls.LoadX509KeyPair(secure.MasterPubCert, secure.MasterPrivateKey)
    if err != nil {
        log.Fatal(err)
        panic(err)
    }
    
    tlsCfg := &tls.Config{Certificates: []tls.Certificate{cert}}
    tlsCfg.InsecureSkipVerify = true
    l, err := tls.Listen("tcp4", ip + ":" + strconv.Itoa(port), tlsCfg)
    
    
    //l, err := net.Listen("tcp", ip + ":" + strconv.Itoa(port))
    if err != nil {
        log.Fatal(err)
        panic(err)
    }
    defer l.Close()
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