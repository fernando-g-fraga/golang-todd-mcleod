package main

import (
	"fmt"
	"math"
)

/*
● create a type SQUARE
○ length float64
○ width float64
● create a type CIRCLE
○ radius float64
● attach a method to each that calculates AREA and returns it
○ circle area= π r 2
■ math.Pi
■ math.Pow
○ square area = L * W
● create a type SHAPE that defines an interface as anything that has the AREA method
● create a func INFO which takes type shape and then prints the area
● create a value of type square
● create a value of type circle
● use func info to print the area of square
● use func info to print the area of circle
*/

type Square struct {
	length float64
	width  float64
}

type Circle struct {
	radius float64
}

func (s Square) Area() float64 {
	return s.length * s.width
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	Area() float64
}

func info(t shape) float64 {
	return t.Area()
}

func main() {

	s1 := Square{
		length: 10,
		width:  10,
	}

	c1 := Circle{
		radius: 1.5,
	}

	fmt.Println(info(c1))
	fmt.Println(info(s1))

}
