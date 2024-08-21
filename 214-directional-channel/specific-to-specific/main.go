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

	cs = cr //doesnt work
	cr = cs //doesnt work

	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", cs)
	fmt.Printf("%T \n", cr)

}