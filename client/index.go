package client

// Run run
func Run() {
	// delete old backup
	go DeleteOldBackup()
	// start client
	go StartBackup()
}
