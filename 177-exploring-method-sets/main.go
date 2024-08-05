package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("i'm ", d.first, "and i'm walking!")
}

func (d *dog) run() {
	fmt.Println("i'm Rover", "and i'm running!")
}

type yougin interface {
	walk()
	run()
}

func youngRun(y yougin) {
	y.run()
}

func main() {

	d1 := dog{"Douglas"}
	d1.walk()
	d1.run()
	// youngRun(d1) - Not going to work (run is a pointer method, d1 is a value sematics.)

	d2 := &dog{"Marshall"}
	d2.walk()
	d2.run()
	youngRun(d2)

}
