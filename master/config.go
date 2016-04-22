package master

import (
    "log"
    "os"
    "github.com/pelletier/go-toml"
    "golem/config"
)

/*
masterConfig ... 
The master configurations
*/
type masterConfig struct {
    CommPort int
	DataPort int
	BindIP   string
}


func GetConfig() masterConfig {
    options := &config.Config
    configExists := checkConfig(options.ConfigFile)
    if !configExists {
        log.Println("Config Data not found... Exiting")
        os.Exit(1)
    }
    config, err := toml.LoadFile(options.ConfigFile)
    if err != nil {
        log.Println(err.Error)
        os.Exit(1)
    }
    
    commPort := config.Get("master.commport").(int)
    dataPort := config.Get("master.dataport").(int)
    bindIP := config.Get("master.bindip").(string)
    
    masterConfig := masterConfig{commPort, dataPort, bindIP}
    return masterConfig
    
}

func checkConfig(name string) bool {
	if _, err := os.Stat(name); err == nil {
		log.Println(name, "exist!")
		return true
	}

	return false
}