package main

import (
	"fmt"
	"math/rand"
)

var x int = rand.Intn(250)

func init() {
	fmt.Println("This is where initialization for my program occurs")
}

func main() {
	hands39()

}

func hands24() {
	//SEQUENCIAL

	fmt.Printf("O valor de x é %v \n", x)

	y := rand.Intn(3)
	fmt.Printf("O valor de y é %v \n", y)

	//CONDITIONAL
	//0>100
	if x > 0 && x < 100 {
		fmt.Println("Entre 0 e 100")
	}
	//101>200
	if x > 100 && x < 200 {
		fmt.Println("Entre 100 e 200")
	}
	//201>250
	if x > 200 && x < 250 {
		fmt.Println("Entre 201 e 250")
	}
}

func hands24switch() {
	switch {
	case x > 0 && x < 100:
		fmt.Println("Entre 0 e 100 - Switch")
	case x > 100 && x < 200:
		{
			fmt.Println("Entre 100 e 200 - Switch")
		}
	default:
		{
			fmt.Println("Entre 201 e 250 - Switch")
		}

	}

}

func hands27() {
	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Printf("O valor de x é %v \n", x)
	fmt.Printf("O valor de x é %v \n", y)

	if x < 4 && y < 4 {
		fmt.Printf("x and y less than 4")
	} else if x > 6 && y > 6 {
		fmt.Printf("x and y less than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Printf("x greater or equal to four and less or equal to 6")
	} else if y != 5 {
		fmt.Println("y is not 5")
	} else {
		fmt.Printf("None of the previous cases were met")

	}
}

func hands27switch() {
	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Printf("O valor de x é %v \n", x)
	fmt.Printf("O valor de y é %v \n", y)

	switch {
	case x < 4 && y < 4:
		fmt.Printf("x and y less than 4")
	case x > 6 && y > 6:
		fmt.Printf("x and y greater than 6")
	case x >= 4 && x <= 6:
		fmt.Printf("○ x is greater than or equal to 4 and less than or equal to 6")
	case y != 5:
		fmt.Println("y is not 5")
	default:
		fmt.Println("○ none of the previous cases were met")
	}

}

func hands29() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}

	for i := 0; i <= 100; i++ {
		fmt.Println("---Contador: %v-------", i)
		hands27switch()
	}
}

func hands30() {

	for i := 0; i < 43; i++ {
		var x int = rand.Intn(5)
		switch x {
		case 1:
			fmt.Printf("A variável x tem valor de: %v", x)
		case 2:
			fmt.Printf("A variável x tem valor de: %v", x)
		case 3:
			fmt.Printf("A variável x tem valor de: %v", x)
		case 4:
			fmt.Printf("A variável x tem valor de: %v", x)
		case 5:
			fmt.Printf("A variável x tem valor de: %v", x)
		default:
			fmt.Printf("A variável x tem valor de: %v", x)
		}

		fmt.Printf("\n Iteração %v\n", i)
	}
}

func hands31() {
	y := -5
	for y > 0 {
		fmt.Println(y)

		y--
	}

}

func hands32() {
	i := 0
	for {
		if i < 10 {
			i++
			fmt.Println(i)
		} else {
			break
		}

	}
}

func hands33() {

	for i := 0; i < 100; i++ {

		if i%2 == 0 {
			continue
		} else {
			fmt.Println(i)
		}
	}
}

func hands34() {

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("Outer Loop %v \t InnerLoop %v \n", i, j)
		}
		fmt.Println("")
	}
}

func hands35() {
	xi := []int{42, 43, 44, 45, 46, 47}

	for key, value := range xi {
		fmt.Printf("O valor da slice é: %v e a posição é: %v \n", value, key)
	}

}

func hands36() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for key, value := range m {
		fmt.Printf("Key: %v Value: %v", key, value)
	}

}

func hands37() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for key, value := range m {
		fmt.Printf("Key: %v Value: %v \n", key, value)
	}

	age := m["James"]
	fmt.Println(age)

	value, ok := m["Q"]

	if !ok {
		fmt.Println("VALOR NAO OK")
	} else {
		fmt.Println(value)
	}

}

func hands38() {
	var cont int
	for i := 0; i < 101; i++ {
		if x := rand.Intn(5); x == 3 {
			println("X is 3")
			cont++
		}
	}
	fmt.Printf("Foram encontrados %v resultados", cont)
}

func hands39() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
