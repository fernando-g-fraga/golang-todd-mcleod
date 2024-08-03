package main

import "fmt"

/*
defer” multiple functions in main
○ show that a deferred func runs after the func containing it exits.
○ determine the order in which the multiple defer funcs run

*/

func main() {

	defer fmt.Println(foo())
	fmt.Println(bar())
	fmt.Println(oneMore())

}

func foo() string {
	return "Numero 1"
}
func bar() string {
	return "Numero 2"
}
func oneMore() string {
	return "Numero 3"
}
