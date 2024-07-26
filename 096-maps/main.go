package main

import "fmt"

func main() {
	am := map[string]int{
		"Douglas":  2,
		"Jack":     27,
		"Fernando": 28,
	}

	fmt.Println(am["Douglas"])
	fmt.Printf("%#v\n", am)
	fmt.Println(len(am))

	am1 := make(map[string]int)
	am1["Lucas"] = 20
	am1["Marcelo"] = 25

	fmt.Println(am1)
	fmt.Printf("%#v\n", am1)
	fmt.Println(len(am1))

	delete(am1, "Lucas")
	fmt.Println(am1)

	if v, ok := am["Douglas"]; !ok {
		fmt.Println("Key didnt exists")
	} else {
		fmt.Println(v)
	}

}
