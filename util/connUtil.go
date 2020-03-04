package util
import (
	"log"
	"net"
)

const maxStringLen = 4096

// ConnSendString send string to connection
func ConnSendString (conn net.Conn, str string) bool {
	remoteAddr := conn.RemoteAddr().String()
	bytes := []byte(str)
	if len(bytes) > maxStringLen {
		log.Fatalln("Max send string length is ", maxStringLen)
		return false
	}
	_, err := conn.Write(bytes)
	if err != nil {
		log.Println("No response ", remoteAddr)
	} else {
		return true
	}
	return false
}


// ConnReceiveString receive string from connection
func ConnReceiveString (conn net.Conn) (str string, err error) {
	buffer := make([]byte, maxStringLen)

	recvLen, err := conn.Read(buffer)
	if err == nil {
		str = string(buffer[:recvLen])
	}
	return
}