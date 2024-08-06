package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByName []Person

func (bn ByName) Len() int           { return len(bn) }
func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByName) Less(i, j int) bool { return bn[i].Name < bn[j].Name }

func main() {

	p1 := Person{Name: "Fernando", Age: 20}
	p2 := Person{Name: "Gilberto", Age: 19}
	p3 := Person{Name: "Roberto", Age: 17}
	p4 := Person{Name: "Alberto", Age: 45}
	p5 := Person{Name: "Pietro", Age: 78}

	people := []Person{p1, p2, p3, p4, p5}

	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)

}
