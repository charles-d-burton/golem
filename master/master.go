package master

import (

	"golem/net/bus"
    "crypto/tls"

)

/*
StartMaster starts the Master processes
*/
func StartMaster(commPort int, dataPort int, bindIP string, role string) {
	log.Println("Loading Master keys....")
    cert, err := tls.LoadX509KeyPair("/etc/golem/pki/master/master_key.pem", "/etc/golem/pki/master/master_pub.pem" )
    if err != nil {
        panic(err)
    }
    log.Println("Keys loaded successfully!")
    tlsCfg := &tls.Config{Certificates: []tls.Certificate{cert}}
    tlsCfg.InsecureSkipVerify = true
	
	go bus.EventBusListener()
	go bus.SocketListener(bindIP, dataPort, *tlsCfg, "data")
	bus.SocketListener(bindIP, commPort, *tlsCfg, "comm")
}