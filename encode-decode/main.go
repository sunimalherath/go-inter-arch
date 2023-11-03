package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	msg := "Assume this is the email message that needs to be encoded with Base64"
	encodedMsg := encode(msg)

	fmt.Println("Encoded Message:", encodedMsg)

	decodedMsg, err := decode(encodedMsg)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Decoded Message:", decodedMsg)
}

func encode(msg string) string {
	return base64.URLEncoding.EncodeToString([]byte(msg))
}

func decode(encodedMessage string) (string, error) {
	decoded, err := base64.URLEncoding.DecodeString(encodedMessage)
	if err != nil {
		return "", fmt.Errorf("couldn't decode the encoded message: %w", err)
	}

	return string(decoded), nil
}
