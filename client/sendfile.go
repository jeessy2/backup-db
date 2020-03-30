package client

import (
	"backup-db/util"
	"bufio"
	"log"
	"net"
	"os"
	"strconv"
)

// SendFile send file to server
func SendFile(fileName string) error {
	log.Printf("Starting send file to server: %s:%d", util.GetConfig().Server.IP, util.GetConfig().Server.ServerPort)

	serverAddr := util.GetConfig().Server.IP + ":" + strconv.Itoa(util.GetConfig().Server.ServerPort)
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		log.Printf("Resolve %s with error: %s \n", serverAddr, err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Println("Connect server error: ", err)
		return err
	}
	defer conn.Close()

	randomKey, err := util.ReceiveRSAPublicKey(conn)
	if err != nil {
		return err
	}

	// send file name, should add parentSavePath + ProjectName
	util.ConnSendString(conn, ParentSavePath+"/"+util.GetConfig().ProjectName+"/"+fileName, randomKey)

	// it's ok?
	ok, err := util.ConnReceiveString(conn, randomKey)
	if err != nil || "ok" != ok {
		return err
	}

	return sendFileReal(fileName, serverAddr, conn, randomKey)
}

func sendFileReal(fileName string, serverAddr string, conn net.Conn, randomKey string) error {

	file, err := os.Open(fileName)
	fileInfo, err := file.Stat()
	if err != nil {
		log.Printf("Read file %s with error: %s\n", fileName, err)
		return err
	}
	defer file.Close()

	// send file size
	fileSize := int(fileInfo.Size())
	util.ConnSendString(conn, strconv.Itoa(fileSize), randomKey)

	// it's ok?
	ok, err := util.ConnReceiveString(conn, randomKey)
	if err != nil || "ok" != ok {
		return err
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

	// read file with bufio
	buffer := make([]byte, 1024)
	reader := bufio.NewReader(file)
	for {
		readLen, err := reader.Read(buffer)
		if err != nil {
			if err.Error() != "EOF" {
				log.Printf("Read file %s with error: %s\n", fileName, err)
			}
			break
		}

		currentSendLen += readLen

		encryptBytes := util.AesGcmEncrypt(randomKey, buffer[:readLen])

		writeLen, err := conn.Write(encryptBytes)
		if err != nil || writeLen != len(encryptBytes) {
			log.Printf("Write file to server %s:%s with error: %s\n", serverAddr, fileName, err)
			progress.StopDisplay = true
			break
		}
	}

	return err

}
