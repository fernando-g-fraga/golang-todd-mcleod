package main

import "fmt"

func main() {
	foo()
	bar("Fernando")
	s1 := aloha("Barnes")
	fmt.Println(s1)

	d1, d2 := dogYears("Douglas", 2)
	fmt.Println(d1, d2)
}

func foo() {
	fmt.Println("Im from foo")
}

func bar(s string) {
	fmt.Println("My Name is", s)
}

func aloha(s string) string {
	return fmt.Sprint("They call me Mr.", s)
}

func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is this old in dog years "), age
}
