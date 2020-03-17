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

		randomKey, err := util.ReceiveRSAPublicKey(conn)
		if err != nil {
			return
		}

		// send file name, should add parentSavePath + ProjectName
		util.ConnSendString(conn, ParentSavePath+"/"+util.GetConfig().ProjectName+"/"+fileName, randomKey)

		// it's ok?
		ok, err := util.ConnReceiveString(conn, randomKey)
		if err != nil || "ok" != ok {
			return
		}

		sendFileReal(fileName, serverAddr, conn, randomKey)
	}
}

func sendFileReal(fileName string, serverAddr string, conn net.Conn, randomKey string) {

	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println("Read file \"", fileName, "\" error: ", err)
		return
	}
	// send file
	fileSize := len(bytes)
	// send file size
	util.ConnSendString(conn, strconv.Itoa(fileSize), randomKey)
	// it's ok?
	ok, err := util.ConnReceiveString(conn, randomKey)
	if err != nil || "ok" != ok {
		return
	}

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
			encryptBytes := util.AesGcmEncrypt(randomKey, bytes[firstIndex:fileSize])
			writeLen, err := conn.Write(encryptBytes)
			if err != nil || writeLen != len(encryptBytes) {
				log.Printf("Write file to server %s : %s with error: %s\n", serverAddr, fileName, err)
				progress.StopDisplay = true
				break
			}
			currentSendLen = fileSize
		} else {
			encryptBytes := util.AesGcmEncrypt(randomKey, bytes[firstIndex:nextIndex])
			writeLen, err := conn.Write(encryptBytes)
			if err != nil || writeLen != len(encryptBytes) {
				log.Printf("Write file to server %s : %s with error: %s\n", serverAddr, fileName, err)
				progress.StopDisplay = true
				break
			}
			currentSendLen = nextIndex
		}
	}

}
