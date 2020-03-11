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
		// delete older backup
		go server.DeleteOlderBackup()
	} else {
		// client
		client.StartBackup()
		// delete older backup
		go client.DeleteOlderBackup()
	}

}
