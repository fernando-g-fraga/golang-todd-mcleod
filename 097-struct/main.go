package main

import "fmt"

func main() {
	h1()
}

type person struct {
	firstName string
	lastName  string
	fIC       []string
}

func h1() {

	jack := person{
		firstName: "Jackeline",
		lastName:  "Duler",
		fIC: []string{
			"Mousse de limão",
			"Brigadeiro",
			"Maracuja"},
	}

	fernando := person{
		firstName: "Fernando",
		lastName:  "Fraga",
		fIC: []string{
			"Ceu Azul",
			"Brigadeiro",
			"Maracuja"},
	}

	fmt.Println(jack.firstName)
	for _, v := range jack.fIC {
		fmt.Printf("Ice cream: %v \n", v)
	}

	fmt.Println(fernando.firstName)
	for _, v := range fernando.fIC {
		fmt.Printf("Ice cream: %v \n", v)
	}

	m1 := map[string][]string{}

	m1[fernando.lastName] = fernando.fIC
	m1[jack.lastName] = jack.fIC

	println()
	for k, v := range m1 {
		fmt.Printf("LastName: %v", k)
		for k1, v2 := range v {
			fmt.Printf("%v | Ice Cream: %v", k1, v2)
		}
	}

}
