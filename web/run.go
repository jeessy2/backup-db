package web

import (
	"backup-db/client"
	"backup-db/server"
	"backup-db/util"
)

// Run run
func Run() {
	switch util.GetEnvType() {
	case "server":
		server.RunCycle()
	default:
		client.RunCycle()
	}
}

// RunOnce run
func RunOnce() {
	switch util.GetEnvType() {
	case "client":
		client.RunOnce()
		break
	}
}
