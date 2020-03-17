package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"log"
)

// AesGcmEncrypt aes gcm encrypt
func AesGcmEncrypt(secretKey string, plainBytes []byte) []byte {
	// Load your secret key from a safe place and reuse it across multiple
	// Seal/Open calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	// When decoded the key should be 16 bytes (AES-128) or 32 (AES-256).
	key, err := hex.DecodeString(secretKey)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	// GetConfig().Server.SecretKey as nonce
	ciphertext := aesgcm.Seal(nil, GetConfig().Server.SecretKey, plainBytes, nil)
	return ciphertext
}

// AesGcmDecrypt aes gcm decrypt
func AesGcmDecrypt(secretKey string, ciphertext []byte) []byte {
	// Load your secret key from a safe place and reuse it across multiple
	// Seal/Open calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	// When decoded the key should be 16 bytes (AES-128) or 32 (AES-256).
	key, _ := hex.DecodeString(secretKey)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, GetConfig().Server.SecretKey, ciphertext, nil)
	if err != nil {
		log.Println(err.Error())
	}

	return plaintext
}
