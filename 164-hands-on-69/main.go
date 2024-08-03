package main

import "fmt"

func main() {
	x()

	v := func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}

	v()

}

var x = func() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
