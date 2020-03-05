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

		// send file name
		util.ConnSendString(conn, fileName)

		// it's ok?
		ok, err := util.ConnReceiveString(conn)
		if err != nil || "ok" != ok {
			return
		}

		// send file size
		util.ConnSendString(conn, strconv.Itoa(len(bytes)))

		// it's ok?
		ok, err = util.ConnReceiveString(conn)
		if err != nil || "ok" != ok {
			return
		}

		// send file
		log.Println("send file...")

		bytesLen := len(bytes)
		for i := 0; i < bytesLen; i++ {
			if i*1024 > bytesLen {
				break
			}
			if (i+1)*1024 > (bytesLen-1) {
				conn.Write(bytes[i*1024 : bytesLen-1])
			} else {
				conn.Write(bytes[i*1024 : (i+1)*1024])
			}
		}

	}

}
