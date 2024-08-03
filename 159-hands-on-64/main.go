package main

import "fmt"

func main() {
	fmt.Printf("%T\n", Add)
	fmt.Printf("%T\n", Subtract)
	fmt.Printf("%T\n", DoMath)
	x := DoMath(42, 16, Add)
	fmt.Println(x)
	y := DoMath(42, 16, Subtract)
	fmt.Println(y)
}

//DoMath gets two integer values plus one operator
func DoMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

//Add sums up two integers and returns the result as integer
func Add(a int, b int) int {
	return a + b
}

//Subtract deduct two integers and returns the diference as integer
func Subtract(a int, b int) int {
	return a - b
}
