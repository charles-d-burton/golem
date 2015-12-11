package main

import (
	"log"
	"flag"
	"golem-server/master"
	"os"
	"os/signal"	
	"syscall"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go handleCtrlC(c)
	var (
		mode = flag.String("mode", "master", "The operational mode of the server")
		commPort = flag.Int("commport", 10000, "The listening communication port")
		dataPort = flag.Int("dataport", 10001, "The port for data transfers")
		bindIP = flag.String("bind", "0.0.0.0", "The IP you want the server to bind to")		
	)
	
	flag.Parse()
	
	if *mode == "master" {
		log.Println("master")
		master.StartMaster(*commPort, *dataPort, *bindIP, "master")
	} else if *mode == "overlord" {
		log.Println("overlord")
	} else if *mode == "peon" {
		log.Println("peon")
	} else {
		log.Println("Invalid mode setting")
	}
	
	
	/*go func() {
		for sig := range c {
			
			log.Println(sig)
		}
	}()*/
}

func handleCtrlC(c chan os.Signal) {
	sig := <-c
	log.Println("\nReceived signal: ", sig)
	os.Remove("/tmp/golem_bus.sock")
	log.Println("Removed UNIX domain socket")
	os.Exit(0)
}
