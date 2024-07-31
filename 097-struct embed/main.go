package main

import (
	"fmt"
)

func main() {

	h2()

	// type engine struct {
	// 	electric bool
	// }

	// type vehicle struct {
	// 	engine
	// 	make  string
	// 	model string
	// 	doors int
	// 	color string
	// }

	// tesla := vehicle{
	// 	engine: engine{
	// 		electric: true},
	// 	make:  "Tesla",
	// 	model: "Model S",
	// 	doors: 4,
	// 	color: "blue",
	// }

	// volvo := vehicle{
	// 	engine: engine{
	// 		electric: false,
	// 	},
	// 	make:  "Volvo",
	// 	model: "EX30",
	// 	doors: 4,
	// 	color: "Space Gray",
	// }

	// fmt.Println(tesla)
	// fmt.Println(volvo)

	// fmt.Println(tesla.color)
	// fmt.Println(volvo.color)
	// fmt.Println(volvo.electric)
	// fmt.Println(volvo.engine.electric)

}

func h2() {

	f1 := struct {
		firstname string
		friends   map[string]int
		favDrinks []string
	}{
		firstname: "Fernando",
		friends:   map[string]int{"Marcos": 1, "Abigail": 1},
		favDrinks: []string{"Apple Juice", "MontainDew", "Beer"},
	}

	for k, v := range f1.friends {
		fmt.Printf(" - %v - %v \n", k, v)
	}

	for i, v := range f1.favDrinks {
		fmt.Printf("Favorite Drinks #%v : %v \n", i, v)
	}

}
