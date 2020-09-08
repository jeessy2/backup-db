package server

import (
	"backup-db/util"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

// Start server
func Start() {
	conf, err := util.GetConfig()
	if err == nil {
		port := ":" + strconv.Itoa(conf.Server.Port)
		listener, err := net.Listen("tcp", port)

		if err == nil {
			log.Println("Started server success!")
			//循环接收客户端的连接，创建一个协程具体去处理连接
			for {
				conn, err := listener.Accept()
				if err != nil {
					log.Println("Accept error", err)
					continue
				}

				// handle connection
				go handleConnection(conn)
			}
		}
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	remoteAddr := conn.RemoteAddr().String()
	log.Println(remoteAddr, "connected success!")
	randomKey, err := util.SendRSAPublicKey(conn)

	// file name
	outFileName, err := util.ConnReceiveString(conn, randomKey)
	if err != nil {
		return
	}
	if !util.ConnSendString(conn, "ok", randomKey) {
		return
	}

	// receive file size
	fileSizeStr, err := util.ConnReceiveString(conn, randomKey)
	if err != nil {
		return
	}
	if !util.ConnSendString(conn, "ok", randomKey) {
		return
	}
	fileSize, err := strconv.Atoi(fileSizeStr)
	if err != nil {
		return
	}
	log.Printf("Receive file: %s ,size: %d ,remote: %s\n", outFileName, fileSize, remoteAddr)

	// receive file bytes
	newFileName := getNewFileName(outFileName)

	// Creating parent dir
	os.MkdirAll(newFileName[0:strings.LastIndex(newFileName, "/")], 0755)
	file, err := os.Create(newFileName)
	if err != nil {
		log.Printf("Can't create file %s , error: %s\n", newFileName, err)
		return
	}
	file.Chmod(0744)

	defer file.Close()

	currentReceivedLen := 0

	progress := util.Progress{
		CurrentReceivedLen: &currentReceivedLen,
		ReceiveOrSend:      "Receive",
		FileSize:           fileSize,
		FileName:           newFileName,
		RemoteAddr:         remoteAddr,
	}
	go util.ProgressDisplay(&progress)

	// receive
	// aes gcm 1024 bytes encrypt to 1040 bytes
	buffer := make([]byte, 1040)
	for {
		readLen, err := conn.Read(buffer)
		if readError(err, &progress) != nil {
			break
		}

		// Read again when the receive is incomplete
		if readLen < 1040 {
			bufferAgain := make([]byte, 1040-readLen)
			readLenAgain, errAgain := conn.Read(bufferAgain)
			if errAgain != nil {
				// last bytes don't display message
				if errAgain.Error() != "EOF" && readError(errAgain, &progress) != nil {
					break
				}
			} else {
				for i := 0; i < readLenAgain; i++ {
					buffer[i+readLen] = bufferAgain[i]
				}
				readLen += readLenAgain
			}
		}

		// write to file
		writeLen, err := file.Write(util.AesGcmDecrypt(randomKey, buffer[:readLen]))
		currentReceivedLen += writeLen
		if err != nil {
			log.Printf("Write %s : %s, error: %s\n", remoteAddr, newFileName, err)
			progress.StopDisplay = true
			break
		}

		// equals break
		if currentReceivedLen == fileSize {
			progress.StopDisplay = true
			break
		}

	}

	// result
	if currentReceivedLen == fileSize {
		log.Printf("Write file success %s : %s, fileSize: %d bytes\n", remoteAddr, newFileName, fileSize)
	} else {
		log.Printf("Write file %s : %s with error! received size is not equals file size!", remoteAddr, newFileName)
	}

}

func readError(err error, progress *util.Progress) error {
	if err != nil {
		log.Printf("Read from %s : %s , error: %s\n", progress.RemoteAddr, progress.FileName, err)
		progress.StopDisplay = true
	}
	return err
}

// make new file name, while be auto add 1 when file exist
// max is 100
func getNewFileName(outFileName string) (newFileName string) {
	if util.PathExists(outFileName) {
		for i := 1; i <= 100; i++ {
			newFileName = outFileName + strconv.Itoa(i)
			if !util.PathExists(newFileName) {
				return
			}
		}
	} else {
		newFileName = outFileName
	}
	return
}
