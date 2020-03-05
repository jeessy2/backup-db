package util

import (
	"log"
	"time"
)

// ProgressDisplay display progress
func ProgressDisplay(progress *Progress) {
	time.Sleep(time.Second * 1)

	for {
		log.Printf("%s file: %s, percent: %.2f %%, remote: %s\n",
			progress.ReceiveOrSend,
			progress.FileName,
			100-float64((progress.FileSize-*progress.CurrentReceivedLen))/float64(progress.FileSize)*100,
			progress.RemoteAddr)

		if *progress.CurrentReceivedLen >= progress.FileSize || progress.StopDisplay {
			break
		}
		time.Sleep(time.Second * 3)
	}
}

// Progress progress
type Progress struct {
	StopDisplay        bool
	CurrentReceivedLen *int
	ReceiveOrSend      string
	FileSize           int
	FileName           string
	RemoteAddr         string
}
