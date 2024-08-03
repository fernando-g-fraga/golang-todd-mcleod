package main

import "fmt"

/*
Create a user defined struct with
○ the identifier “person”
○ the fields:
■ first
■ age
● attach a method to type person with
○ the identifier “speak”
○ the method should have the person say their name and age
● create a value of type person
● call the method from the value of type person

*/

type person struct {
	first string
	age   int
}

func (p person) speak() string {
	return "Hello, my name is " + p.first
}
func main() {

	p1 := person{
		first: "fernando",
		age:   19,
	}

	fmt.Println(p1.speak())

}
