package main

import (
	"fmt"

	"digital-envelope/common/box"
	"digital-envelope/common/envelope"
	"digital-envelope/common/secretbox"
)

func main() {
	secretKey := secretbox.GenerateSecretKey()
	fmt.Println(secretKey)

	sessionPub, sessionPri, err := box.GenerateKeyPair()
	fmt.Println(sessionPub, sessionPri, err)

	msg := "f*ck envelope seal and open ?"
	fmt.Println(msg)

	cipher, encryptedSecretKey, tempPub := envelope.Seal(msg, secretKey, sessionPub)
	fmt.Println(cipher, encryptedSecretKey, tempPub)

	plain, ok := envelope.Open(cipher, encryptedSecretKey, tempPub, sessionPri)
	fmt.Println(plain, ok)
}
