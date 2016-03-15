package main

import (
	"github.com/codegangsta/cli"
	"golem/config"
	"golem/master"
	"golem/secure"
	"log"
	"os"
	"os/signal"
	"os/user"
	"syscall"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	user, err := user.Current()
	if user.Uid != "0" || err != nil {
		panic("You must run this program as root!")
		os.Exit(0)
	} else {
		
		log.Println(user.Uid)
		config.MakeYamlFile()
		config.OpenYaml()
		
		options := &config.Config
	
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		go handleCtrlC(c)
		golem := cli.NewApp()
		golem.Name = "golem"
		golem.Author = "Charles Burton"
		golem.Email = "charles.d.burton@gmail.com"
		golem.Usage = "A fast remote execution engine written in Go"
		golem.Flags = []cli.Flag {
			cli.StringFlag{
				Name:        "m, mode",
				Value:       "cli",
				Usage:       "Set the application mode to peon, master, or overlord.  Defaults to command line",
				Destination: &options.Mode,
			},
			cli.IntFlag{
				Name:        "d, dataport",
				Value:       10001,
				Usage:       "Set the server port that data will transfer over. Default: 10001",
				Destination: &options.DataPort,
			},
			cli.IntFlag{
				Name:        "c, commport",
				Value:       10000,
				Usage:       "Set the port that golem will communicate with. Default: 10000",
				Destination: &options.CommPort,
			},
			cli.StringFlag{
				Name:        "b, bind",
				Value:       "0.0.0.0",
				Usage:       "Sets the IP golem will bind to.",
				Destination: &options.BindIP,
			},
            cli.StringFlag{
                Name:       "C, config",
                Value:      "/etc/golem/golem.conf",
                Usage:      "Load alternate config file",
                Destination: &options.ConfigFile,  
            },
		}
	
		golem.Commands = []cli.Command{}
		golem.Run(os.Args)
		
		/*if mode == "master" {
			master.StartMaster(commPort, dataPort, bindIP, mode)
		} else if mode == "peon" {
			
		} else if mode == "overlord" {
			
		} else {
			runCommandLine(*golem)
		}*/
		if options.Mode == "master" {
            err := secure.SetupKeys(options.Mode)
		    check(err)
		    master.StartMaster(options.CommPort, options.DataPort, options.BindIP, options.Mode)
            
		}
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


