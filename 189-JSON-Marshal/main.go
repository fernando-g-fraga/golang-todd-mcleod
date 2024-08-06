package main

import (
	"encoding/json"
	"fmt"
)

type pessoas struct {
	First string
	Last  string
	Age   int
}

func main() {

	p1 := pessoas{
		First: "Fernando",
		Last:  "Fraga",
		Age:   28,
	}

	p2 := pessoas{
		First: "Marshall",
		Last:  "Douglas",
		Age:   2,
	}

	compilado := []pessoas{p1, p2}

	b, err := json.Marshal(compilado)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))

}
