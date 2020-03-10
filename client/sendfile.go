package client

import (
	"backup-db/util"
	"io/ioutil"
	"log"
	"net"
	"strconv"
)

// SendFile send file to server
func SendFile(fileName string) {
	log.Println("Starting send file to server: ", util.GetConfig().Server.IP, ":", util.GetConfig().Server.ServerPort)
	fileAllBytes, err := ioutil.ReadFile(fileName)
	if err == nil {
		sendFileInner(fileName, fileAllBytes)
	} else {
		log.Println("Read file \"", fileName, "\" error: ", err)
	}

}

func sendFileInner(fileName string, fileAllBytes []byte) {
	serverAddr := util.GetConfig().Server.IP + ":" + strconv.Itoa(util.GetConfig().Server.ServerPort)
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		log.Printf("Resolve %s with error: %s \n", serverAddr, err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Println("Connect server error: ", err)
	} else {
		defer conn.Close()

		// send file name, should add parentSavePath + ProjectName
		util.ConnSendString(conn, parentSavePath+"/"+util.GetConfig().ProjectName+"/"+fileName)

		// it's ok?
		ok, err := util.ConnReceiveString(conn)
		if err != nil || "ok" != ok {
			return
		}

		// send file size
		util.ConnSendString(conn, strconv.Itoa(len(fileAllBytes)))

		// it's ok?
		ok, err = util.ConnReceiveString(conn)
		if err != nil || "ok" != ok {
			return
		}

		// send file
		fileSize := len(fileAllBytes)

		currentSendLen := 0
		progress := util.Progress{
			CurrentReceivedLen: &currentSendLen,
			ReceiveOrSend:      "Send",
			FileSize:           fileSize,
			FileName:           fileName,
			RemoteAddr:         serverAddr,
		}
		go util.ProgressDisplay(&progress)

		for i := 0; i < fileSize; i++ {
			firstIndex := i * 1024
			nextIndex := (i + 1) * 1024
			if firstIndex > fileSize {
				progress.StopDisplay = true
				break
			}
			if nextIndex >= fileSize {
				// can't over
				len, err := conn.Write(fileAllBytes[firstIndex:fileSize])
				if err != nil || len != fileSize-firstIndex {
					log.Printf("Write file to server %s : %s with error: %s\n", serverAddr, fileName, err)
					progress.StopDisplay = true
					break
				}
				currentSendLen = fileSize
			} else {
				len, err := conn.Write(fileAllBytes[firstIndex:nextIndex])
				if err != nil || len != nextIndex-firstIndex {
					log.Printf("Write file to server %s : %s with error: %s\n", serverAddr, fileName, err)
					progress.StopDisplay = true
					break
				}
				currentSendLen = nextIndex
			}
		}

	}

}
