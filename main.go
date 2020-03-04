package main

import (
	"backup-db/client"
	"backup-db/server"
	"backup-db/util"
)

func main() {

	if util.GetConfig().Server.IP == "" {
		// server
		server.Start()
	} else {
		// client
		client.StartBackup()
	}

}
