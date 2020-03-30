package main

import (
	"backup-db/client"
	"backup-db/server"
	"backup-db/util"
)

func main() {

	if util.GetConfig().Server.IP == "" {
		// delete old backup
		go server.DeleteOldBackup()
		// inspect backup
		go server.InspectBackup()
		// start server
		server.Start()
	} else {
		// delete old backup
		go client.DeleteOldBackup()
		// start client
		client.StartBackup()
	}

}
