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
	bytes, err := ioutil.ReadFile(fileName)
	if err == nil {
		sendFileInner(fileName, bytes)
	} else {
		log.Println("Read file \"", fileName, "\" error: ", err)
	}

}

func sendFileInner(fileName string, bytes []byte) {
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
		conn.Write([]byte(fileName))
		buffer := make([]byte, 1024)

		// it's ok?
		okLen, err := conn.Read(buffer)
		if err != nil {
			log.Printf("Can't read \"ok\" from server %s, err: %s", serverAddr, err)
			return
		}

		ok := string(buffer[:okLen])
		if ok != "ok" {
			return
		}

		// send file
		log.Println("send file...")
		conn.Write(bytes)

	}

}
