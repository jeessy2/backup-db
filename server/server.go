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
	port := ":" + strconv.Itoa(util.GetConfig().Server.ServerPort)
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
	} else {
		log.Panicln("Start server with error: ", err)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	remoteAddr := conn.RemoteAddr().String()
	log.Println(remoteAddr, "connected success!")

	// file name
	outFileName, err := util.ConnReceiveString(conn)
	if err != nil {
		return
	}
	if !util.ConnSendString(conn, "ok") {
		return
	}

	// receive file size
	fileSizeStr, err := util.ConnReceiveString(conn)
	if err != nil {
		return
	}
	if !util.ConnSendString(conn, "ok") {
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
	os.MkdirAll(newFileName[0 : strings.LastIndex(newFileName, "/")], 0755)
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

	for {
		buffer := make([]byte, 1024)
		len, err := conn.Read(buffer)
		if err != nil {
			log.Printf("Read from %s : %s , error: %s\n", remoteAddr, newFileName, err)
			progress.StopDisplay = true
			break
		}

		if len > 0 {
			writeLen, err := file.Write(buffer[:len])
			currentReceivedLen += writeLen
			if err != nil || writeLen != len {
				log.Printf("Write file %s : %s, error: %s\n", remoteAddr, newFileName, err)
				progress.StopDisplay = true
				break
			}
			// equals, received file success
			if currentReceivedLen == fileSize {
				log.Printf("Write file success %s : %s\n", remoteAddr, newFileName)
				progress.StopDisplay = true
				break
			}
		} else {
			break
		}
	}

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
