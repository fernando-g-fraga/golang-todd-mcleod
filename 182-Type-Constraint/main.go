package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type myNumbers interface {
	constraints.Integer | constraints.Float
}

type myAlias int

func addT[T myNumbers](a, b T) T {
	return a + b
}
func main() {
	var n myAlias = 42
	fmt.Println(addT(n, 2))
	fmt.Println(addT(2, 2.5))
	fmt.Println(addT(2, 2))

}
