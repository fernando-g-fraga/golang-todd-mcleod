package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {

	s := `[{"First":"Fernando","Last":"Fraga","Age":28},{"First":"Marshall","Last":"Douglas","Age":2}]`
	bs := []byte(s)

	var people []Person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("all of the data", people)

	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i+1)
		fmt.Println(v.First, v.Last, v.Age)
	}

}
