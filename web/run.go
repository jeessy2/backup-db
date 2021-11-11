package web

import (
	"backup-db/client"
)

// Run run
func Run() {
	client.RunCycle()
}

// RunOnce run
func RunOnce() {
	client.RunOnce()
}
