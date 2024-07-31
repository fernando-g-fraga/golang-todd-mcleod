package main

import "fmt"

type Person struct {
	first string
}

type SecretAgent struct {
	Person
	ltk bool
}

type Human interface {
	speak()
}

func (p Person) speak() {
	fmt.Println("I am", p.first)
}

func (sa SecretAgent) speak() {
	fmt.Println("I am a secret agent,", sa.first)
}

func main() {

	p1 := Person{first: "Fernando"}
	p1.speak()

	sa1 := SecretAgent{
		Person: Person{first: "James"},
		ltk:    true,
	}

	sa1.speak()

}
