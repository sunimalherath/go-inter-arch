package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Name string
}

func main() {

	// p2 := person{
	// 	Name: "Lada",
	// }

	// xp := []person{p1, p2}

	// bs, err := json.Marshal(xp)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// fmt.Println("JSON:", string(bs))

	// xp2 := []person{}

	// err = json.Unmarshal(bs, &xp2)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// fmt.Println("Go data structure:", xp2)

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		Name: "Jeny",
	}

	p2 := person{
		Name: "Lada",
	}

	xp := []person{p1, p2}

	err := json.NewEncoder(w).Encode(xp)
	if err != nil {
		log.Println("Bad data to encode", err)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	var xp []person

	err := json.NewDecoder(r.Body).Decode(&xp)
	if err != nil {
		log.Println("Bad data to decode", err)
	}

	log.Println("Person", xp)
}
