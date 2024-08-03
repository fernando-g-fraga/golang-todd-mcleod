package main

import "fmt"

func main() {

	y := outer()
	fmt.Println(y())

}

func outer() func() int {
	return func() int {
		return 42
	}
}
