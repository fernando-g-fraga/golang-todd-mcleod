package main

import "fmt"

type person struct {
	first string
}

func changeName(p person, s string) person {
	p.first = s
	return p
}

func changeNameP(p *person, s string) {
	p.first = s
}

func main() {

	p1 := person{"Fernando"}
	fmt.Println(p1.first)
	p2 := person{"Marcos"}
	fmt.Println(p2.first)

	p1 = changeName(p1, "Gabriel")
	fmt.Println(p1.first)
	changeNameP(&p2, "Gilberto")
	fmt.Println(p2.first)

}
