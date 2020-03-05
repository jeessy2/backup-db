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

		// send file name
		util.ConnSendString(conn, fileName)

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
		log.Println("send file...")

		fileSize := len(fileAllBytes)

		currentSendLen := 0
		go util.ProgressDisplay("Send", &currentSendLen, fileSize, fileName, serverAddr)

		for i := 0; i < fileSize; i++ {
			firstIndex := i * 1024
			nextIndex := (i + 1) * 1024
			if firstIndex > fileSize {
				break
			}
			if nextIndex >= fileSize {
				// can't over
				conn.Write(fileAllBytes[firstIndex:fileSize])
				currentSendLen = fileSize
			} else {
				conn.Write(fileAllBytes[firstIndex:nextIndex])
				currentSendLen = nextIndex
			}
		}

	}

}
