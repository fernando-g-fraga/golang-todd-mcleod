package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (ba ByAge) Len() int           { return len(ba) }
func (ba ByAge) Swap(i, j int)      { ba[i], ba[j] = ba[j], ba[i] }
func (ba ByAge) Less(i, j int) bool { return ba[i].Age < ba[j].Age }

type ByLast []user

func (bl ByLast) Len() int           { return len(bl) }
func (bl ByLast) Swap(i, j int)      { bl[i], bl[j] = bl[j], bl[i] }
func (bl ByLast) Less(i, j int) bool { return bl[i].Last < bl[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)
	fmt.Println("---------Sort: Users---------")
	for i, v := range users {
		fmt.Printf("Person: #%v | Name: %v | Last: %v | Age: %v \n", i, v.First, v.Last, v.Age)
		sort.Strings(v.Sayings)
		for _, saying := range v.Sayings {
			fmt.Println("\t ", saying)
		}
	}
	fmt.Println("---------Sort: ByAge---------")
	sort.Sort(ByAge(users))
	for i, v := range users {
		fmt.Printf("Person: #%v | Name: %v | Last: %v | Age: %v \n", i, v.First, v.Last, v.Age)
		for _, saying := range v.Sayings {
			fmt.Println("\t ", saying)
		}
	}
	sort.Sort(ByLast(users))
	fmt.Println("---------Sort: ByLast---------")
	for i, v := range users {
		fmt.Printf("Person: #%v | Name: %v | Last: %v | Age: %v \n", i, v.First, v.Last, v.Age)
		for _, saying := range v.Sayings {
			fmt.Println("\t ", saying)
		}
	}

}
