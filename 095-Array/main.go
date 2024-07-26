package main

import (
	"fmt"
)

func main() {

	h5()

}

func h1() {
	xi := [5]int{}

	for i := 0; len(xi) > i; i++ {
		xi[i] = i
	}

	for i, v := range xi {
		fmt.Printf("%v - %T - %v \n", i, v, v)
	}
}

func h2() {
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range xi {
		fmt.Printf("%v | %T | %v \n ", i, v, v)
	}
	fmt.Println("-----------")

	// [42 43 44 45 46]
	// [47 48 49 50 51]
	// [44 45 46 47 48]
	// [43 44 45 46 47]

	fmt.Printf("%v \n", xi[:5])
	fmt.Printf("%v \n", xi[5:])
	fmt.Printf("%v \n", xi[2:7])
	fmt.Printf("%v \n", xi[1:6])

}

func h3() {

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)

	// x = append(x, 53, 54, 55)
	// fmt.Println(x)

	// y := []int{56, 57, 58, 59, 60}
	// x = append(x, y...)
	// fmt.Println(x)

	//[42, 43, 44, 48, 49, 50, 51]

	x = append(x[:3], x[6:]...)
	fmt.Println(x)

}

func h4() {

	xs1 := make([]int, 50)
	fmt.Println(xs1)
	fmt.Println(len(xs1))
	fmt.Println(cap(xs1))

	xs2 := make([]int, 0, 50)
	fmt.Println(xs2)
	fmt.Println(len(xs2))
	fmt.Println(cap(xs2))

	xs1 = append(xs1, 98)
	xs2 = append(xs2, 99)

	fmt.Println(xs1)
	fmt.Println("--------")
	fmt.Println(xs2)

	xs := make([]string, 0, 50)

	xs = append(xs, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`,
		`Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`,
		`Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, `Mississippi`,
		` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`,
		` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`,
		` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, `Utah`, ` Vermont`, ` Virginia`, ` Washington`,
		` West Virginia`, ` Wisconsin`, ` Wyoming`)

	fmt.Println(xs)
	fmt.Println(len(xs))
	fmt.Println(cap(xs))

	for i := 0; i < len(xs); i++ {
		fmt.Print(xs[i], i)
	}
}

func h5() {

	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "I'm 008."}

	xs3 := [][]string{xs1, xs2}

	fmt.Println(xs3)

	// for i := range xs3 {
	// 	for j := range xs3[i] {
	// 		fmt.Println(xs3[i][j])
	// 	}
	// }

	for i, v := range xs3 {
		fmt.Println(i, v)
		for a, b := range v {
			fmt.Println(a, b)
		}
	}
}
