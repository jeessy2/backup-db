package server

// RunCycle runner
func RunCycle() {
	// delete old backup
	go DeleteOldBackup()
	// inspect backup
	go InspectBackup()
}
