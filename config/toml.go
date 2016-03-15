package config

import (
    "sync"
)

//Set the configuration options
type masterConfiguration struct {
    DataPort    int
    CommPort    int
    PrivateKey  [32]byte
    BindIP      string
}

type peonConfiguration struct {
    Master          []string
    MasterDataPort  int
    MasterCommPort  int
    PrivateKey      [32]byte
        
}

var (
    masterConfig *masterConfiguration
    peonConfig *peonConfiguration
    
    //These are one-time synchronization locks
    masterOnce sync.Once
    peonOnce sync.Once
)

//GetMasterConfig ... load and retrieve the master configuration
func GetMasterConfig() *masterConfiguration {
    //Thread safe, fast, one-time initialization of config
    masterOnce.Do(func() {
        masterConfig = &masterConfiguration{}
    })
    return masterConfig
}

//GetPeonConfig ... load and retrieve the peon configuration
func GetPeonConfig() *peonConfiguration {
    //Thread safe, fast, one-time initialization of config
    peonOnce.Do(func() {
        peonConfig = &peonConfiguration{}
    })
    return peonConfig
}
