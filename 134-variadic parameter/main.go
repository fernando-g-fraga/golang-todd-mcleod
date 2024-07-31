package main

import "fmt"

func main() {

	teste := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := foo(teste...)
	fmt.Println(sum)
}

func foo(ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T \n", ii)

	n := 0
	for _, v := range ii {
		n += v
	}
	return n

}

func bar(s string, ii ...int) (string, int) {

	n := 0
	for _, v := range ii {
		n += v
	}
	return s, n
}
