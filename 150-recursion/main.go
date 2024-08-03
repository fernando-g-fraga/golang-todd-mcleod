package main

import "fmt"

func main() {

	fmt.Println(factorial(4))
	fmt.Println(loop(4))

}

func factorial(a int) int {
	fmt.Println("this is the value of A now: ", a)
	if a == 0 {
		return 1
	}
	return a * factorial(a-1)
}

func loop(a int) int {
	total := 1
	for a > 0 {
		total *= a
		a--
	}
	return total
}
