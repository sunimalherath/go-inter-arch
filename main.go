package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := "pass123456"

	hashedPass, err := hashPassword(pass)
	if err != nil {
		panic(err)
	}

	err = comparePassword(pass, hashedPass)
	if err != nil {
		log.Fatalln("Not logged in")
	}

	log.Println("Logged in!")
}

func hashPassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error while generating hash from password, %w", err)
	}

	return bs, nil
}

func comparePassword(password string, hashPassword []byte) error {
	err := bcrypt.CompareHashAndPassword(hashPassword, []byte(password))
	if err != nil {
		return fmt.Errorf("invalid password, %w", err)
	}

	return nil
}
