package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("The title of the book is: ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("This is the number: ", strconv.Itoa(int(c)))
}

func logInfo(s fmt.Stringer) {
	log.Println("Log from 139", s.String())
}

func main() {

	b := book{
		title: "Capitães da Areia",
	}

	var c count = 10

	logInfo(b)
	logInfo(c)

}
