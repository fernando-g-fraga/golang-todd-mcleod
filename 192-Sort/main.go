package main

import (
	"fmt"
	"slices"
)

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	fmt.Println(xi)
	slices.Sort(xi)
	fmt.Println(xi)
	slices.Sort(xs)
	fmt.Println(xs)
}
