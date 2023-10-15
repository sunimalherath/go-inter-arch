package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	Name string
}

func main() {
	p1 := person{
		Name: "Jeny",
	}

	p2 := person{
		Name: "Lada",
	}

	xp := []person{p1, p2}

	bs, err := json.Marshal(xp)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("JSON:", string(bs))

	xp2 := []person{}

	err = json.Unmarshal(bs, &xp2)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Go data structure:", xp2)
}
