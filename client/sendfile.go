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
		conn.Write(bytes)

	}

}
