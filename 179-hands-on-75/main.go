package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {

	fmt.Printf("Memory Address of a: %v \n", a)
	fmt.Printf("Memory Address of b: %v \n", b)
	fmt.Printf("Memory Address of c: %v \n", c)
	fmt.Printf("Memory Address of d: %v \n", d)

	fmt.Printf("Type of a: %T \n", a)
	fmt.Printf("Type of b: %T \n", b)
	fmt.Printf("Type of c: %T \n", c)
	fmt.Printf("Type of d: %T \n", d)

	fmt.Printf("Values of a: %v \n", *a)
	fmt.Printf("Values of b: %v \n", *b)
	fmt.Printf("Values of c: %v \n", *c)
	fmt.Printf("Values of d: %v \n", *d)

}
