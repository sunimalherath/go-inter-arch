package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	msg := "This is the sample message that's going to be encoded and decoded with AES"

	fmt.Println("Original message:", msg)

	password := "SuperStr0ng"

	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Fatalln(err)
	}

	// take the first 16 since AES only ....
	bs = bs[:16]

	// encoding the message
	res, err := encodeDecode(bs, msg)
	if err != nil {
		log.Fatalln("error while encoding", err)
	}

	fmt.Println("AES encoded message:", string(res))

	// decoding the message
	res2, err := encodeDecode(bs, string(res))
	if err != nil {
		log.Fatalln("error while decoding", err)
	}

	fmt.Println("AES decoded message:", string(res2))
}

func encodeDecode(key []byte, input string) ([]byte, error) {
	blk, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// iv: initialization vector
	iv := make([]byte, aes.BlockSize)

	stream := cipher.NewCTR(blk, iv)

	buff := &bytes.Buffer{}

	sw := cipher.StreamWriter{
		S: stream,
		W: buff,
	}

	_, err = sw.Write([]byte(input))
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}
