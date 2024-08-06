package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Color struct {
	Id     int
	Name   string
	Colors []string
}

func main() {

	var jsonBlob = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll",    "Order": "Dasyuromorphia"}
		]`)

	type Animal struct {
		Name  string
		Order string
	}

	var animals []Animal

	err := json.Unmarshal(jsonBlob, &animals)

	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%+v", animals)

	g1 := Color{
		Id:     1,
		Name:   "Grupo 1",
		Colors: []string{"Red", "Green", "Blue", "Brown", "Yellow", "Black"},
	}

	b, err := json.Marshal(g1)

	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(b)

}
