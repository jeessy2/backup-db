package util

import (
	"log"
	"time"
)

// ProgressDisplay display progress
func ProgressDisplay(receiveOrSender string, currentReceivedLen *int, fileSize int, fileName string, remoteAddr string) {
	time.Sleep(time.Second * 1)

	for {
		log.Printf("%s file: %s, percent: %.2f %%, remote: %s\n",
			receiveOrSender,
			fileName,
			100-float64((fileSize-*currentReceivedLen))/float64(fileSize)*100,
			remoteAddr)

		if *currentReceivedLen >= fileSize {
			break
		}
		time.Sleep(time.Second * 3)
	}
}
