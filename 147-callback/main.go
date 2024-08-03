package main

import "fmt"

func main() {

	x := doMath(12, 24, sub)
	fmt.Println(x)
	y := doMath(12, 24, add)
	fmt.Println(y)

	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", sub)
	fmt.Printf("%T\n", doMath)

}

func add(a int, b int) int {
	return a + b
}
func sub(a int, b int) int {
	return a - b
}
func doMath(a int, b int, f func(a int, b int) int) int {
	return f(a, b)
}
