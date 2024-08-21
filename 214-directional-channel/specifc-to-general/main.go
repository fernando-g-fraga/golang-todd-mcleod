package main

import "fmt"

func main() {
	c := make(chan int)    //channel
	cr := make(<-chan int) //channel receive-only
	cs := make(chan<- int) //channel send-only

	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", cs)
	fmt.Printf("%T \n", cr)

	fmt.Println("-------------------")

	cr = c
	cs = c

	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", cs)
	fmt.Printf("%T \n", cr)

}
