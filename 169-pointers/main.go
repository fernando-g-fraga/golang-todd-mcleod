package main

import "fmt"

func intDelta(n *int) {
	*n = 43
}

func xiDelta(ii []int) {
	ii[0] = 99

}

func mapDelta(m map[string]int, s string) {
	m[s] = 99
}

func main() {
	a := 42

	fmt.Println(a)
	intDelta(&a)
	fmt.Println(a)

	xi := []int{1, 2, 3, 4}
	fmt.Println(xi)
	xiDelta(xi)
	fmt.Println(xi)

	m := make(map[string]int)
	m["James"] = 32
	fmt.Println(m["James"])
	mapDelta(m, "James")
	fmt.Println(m["James"])

}
