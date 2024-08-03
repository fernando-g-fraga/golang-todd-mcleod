package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	v := openfile("teste.txt")
	fmt.Println(v)
}

func openfile(filename string) string {
	f, err := os.ReadFile(filename)

	if err != nil {
		log.Fatalln("There was an fatal error opening the file.", err)
	}
	return string(f)

}
