package util

import (
	"log"
	"net"
)

const maxStringLen = 4096

// SendRSAPublicKey Send public key to client, return private key
func SendRSAPublicKey(conn net.Conn) (randomKey string, err error) {
	// send public key
	prvkey, pubkey := GenRsaKey()
	_, err = conn.Write(pubkey)
	if err != nil {
		log.Println("SendRSAPublicKey with error: ", err)
		return
	}
	buffer := make([]byte, maxStringLen)
	// decrypt random key
	keyLen, err := conn.Read(buffer)
	if err != nil {
		log.Println("Receive random key with error: ", err)
		return
	}
	randomKey = string(RsaDecrypt(buffer[:keyLen], prvkey))
	return
}

// ReceiveRSAPublicKey Send public key to client, return private key
func ReceiveRSAPublicKey(conn net.Conn) (randomKey string, err error) {
	// receive public key
	buffer := make([]byte, maxStringLen)
	recvLen, err := conn.Read(buffer)
	if err != nil {
		log.Println("ReceiveRSAPublicKey with error: ", err)
		return
	}
	pubKey := buffer[:recvLen]
	// encrypt random key
	randomKey = GetRandomString(32)
	_, err = conn.Write(RsaEncrypt([]byte(randomKey), pubKey))
	if err != nil {
		log.Println("Send random key with error: ", err)
		return
	}
	return
}

// ConnSendString send string to connection
func ConnSendString(conn net.Conn, str string, randomKey string) bool {
	remoteAddr := conn.RemoteAddr().String()
	// aes encrypt
	bytes := AesGcmEncrypt(randomKey, []byte(str))
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
func ConnReceiveString(conn net.Conn, randomKey string) (str string, err error) {
	buffer := make([]byte, maxStringLen)

	recvLen, err := conn.Read(buffer)
	if err == nil {
		// aes decrypt
		str = string(AesGcmDecrypt(randomKey, buffer[:recvLen]))
	}
	return
}
