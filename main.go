package main

import (
	"log"
	"golem/master"
	"os"
	"os/signal"	
	"syscall"
	"github.com/codegangsta/cli"
)

var (
	mode = "cli"
	commPort = 10000
	dataPort = 10001
	bindIP = "0.0.0.0"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go handleCtrlC(c)
	golem := cli.NewApp()
	golem.Name = "golem"
	golem.Author = "Charles Burton"
	golem.Email = "charles.d.burton@gmail.com"
	golem.Usage = "A fast remote execution engine written in Go"
	golem.Flags = []cli.Flag {
		cli.StringFlag {
			Name: "m, mode",
			Value: "cli",
			Usage: "Set the application mode to peon, master, or overlord.  Defaults to command line",
			Destination: &mode,
		},
		cli.IntFlag {
			Name: "d, dataport",
			Value: 10001,
			Usage: "Set the server port that data will transfer over. Default: 10001",
			Destination: &dataPort,
		},
		cli.IntFlag {
			Name: "c, commport",
			Value: 10000,
			Usage: "Set the port that golem will communicate with. Default: 10000",
			Destination: &commPort,
		},
		cli.StringFlag {
			Name: "b, bind",
			Value: "0.0.0.0",
			Usage: "Sets the IP golem will bind to.",
			Destination: &bindIP,
		},
	}
	
	golem.Commands = []cli.Command {
		
	}
	golem.Run(os.Args)
	
	if mode == "master" {
		master.StartMaster(commPort, dataPort, bindIP, mode)
	} else if mode == "peon" {
		
	} else if mode == "overlord" {
		
	} else {
		runCommandLine(*golem)
	}
}

func runCommandLine(golem cli.App) {
	
}


func handleCtrlC(c chan os.Signal) {
	sig := <-c
	log.Println("\nReceived signal: ", sig)
	os.Remove("/tmp/golem_bus.sock")
	log.Println("Removed UNIX domain socket")
	os.Exit(0)
}
