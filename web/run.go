package web

import (
	"backup-db/client"
	"backup-db/server"
	"backup-db/util"
)

// Run run
func Run() {
	conf, err := util.GetConfig()
	if err == nil {
		switch conf.Type {
		case "server":
			server.Run()
		default:
			client.Run()
		}
	}
}

// RunOnce run
func RunOnce() {
	conf, err := util.GetConfig()
	if err == nil {
		switch conf.Type {
		case "server":
			// todo runonce
			server.Run()
		default:
			client.RunOnce()
		}
	}
}
