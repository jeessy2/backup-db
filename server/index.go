package server

// Run runner
func Run() {
	// delete old backup
	go DeleteOldBackup()
	// inspect backup
	go InspectBackup()
}
