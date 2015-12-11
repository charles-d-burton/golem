package main

import (
	"log"
	"flag"
	"golem-server/master"
	"os"
	"os/signal"
	"stringify"
	
)

func main() {
	var (
		mode = flag.String("mode", "master", "The operational mode of the server")
		commPort = flag.Int("commport", 10000, "The listening communication port")
		dataPort = flag.Int("dataport", 10001, "The port for data transfers")
		
	)
	
	flag.Parse()
	
	
	if *mode == "master" {
		log.Println("master")
		master.StartMaster(*commPort, *dataPort)
	} else if *mode == "overlord" {
		log.Println("overlord")
	} else if *mode == "peon" {
		log.Println("peon")
	} else {
		log.Println("Invalid mode setting")
	}
	
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			log.Println(stringify(sig))
		}
	}()
}
