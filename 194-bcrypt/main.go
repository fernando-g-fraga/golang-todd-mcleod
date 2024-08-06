package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	s := "Fe159753"
	s2 := "fe159753"

	bc, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
	fmt.Println(bc)

	ok := bcrypt.CompareHashAndPassword(bc, []byte(s2))

	fmt.Println(ok)
}
