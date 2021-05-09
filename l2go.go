package main

import (
	"./config"
	"./gameserver"
	"./loginserver"
	"flag"
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var mode, gameServerId int
	flag.IntVar(&mode, "mode", 0, "Set to 0 to run the Login Server or 1 to run the Game Server")
	flag.IntVar(&gameServerId, "server", 1, "Set the id of the Game Server you want to run")
	flag.Parse()

	// Load the global configuration object
	globalConfig := config.Read()

	if mode == 0 {
    server := loginserver.New(globalConfig)
    server.Init()
    server.Start()
	} else {
		// Try to load the Game Server configuration
		if gameServerId >= 1 && len(globalConfig.GameServers) >= gameServerId {
      srvConfig := config.GameServerConfigObject{}
			srvConfig.LoginServer = globalConfig.LoginServer
			srvConfig.GameServer = globalConfig.GameServers[gameServerId-1]
      server := gameserver.New(srvConfig)
      server.Init()
      server.Start()
		} else {
			fmt.Println("No configuration found for the specified server.")
		}

	}

	fmt.Println("Server stopped.")
}
