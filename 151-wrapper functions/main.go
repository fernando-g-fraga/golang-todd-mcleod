package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	xb, err := readfile("poem.txt")
	if err != nil {
		log.Fatalf("readfile in main: %s", err)
	}

	fmt.Println(xb)
	fmt.Println(string(xb))

}

func readfile(fn string) ([]byte, error) {

	xb, err := os.ReadFile(fn)
	if err != nil {
		return nil, fmt.Errorf("Error in readfile: %s", err)
	}
	return xb, nil

}
