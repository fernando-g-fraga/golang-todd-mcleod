package main

import (
	"fmt"

	"github.com/fernando-g-fraga/puppy"
)

var outsiderV string = "Vim do exterior"

const outsiderC string = ("Vim de exterior, mas sou constante")

func main() {

	inner := "Vim de dentro"

	fmt.Println(outsiderV)
	fmt.Println(outsiderC)
	fmt.Println(inner)

	fmt.Println(puppy.Bark())

}
