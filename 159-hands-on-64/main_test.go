package main

import (
	"log"
	"testing"
)

// func testAdd(t *testing.T) {
// 	total := Add(5, 5)
// 	if total != 10 {
// 		t.Errorf("Sum was incorret, got: %d, want: %d", total, 10)
// 	}
// }

func TestAdd(t *testing.T) {
	got := Add(7, 5)
	want := 12
	if got != want {
		log.Fatalf("error - wants %v and got %v", want, got)
	}
}

func TestSubtract(t *testing.T) {
	got := Subtract(10, 5)
	want := 5

	if got != want {
		log.Fatalf("Error - Wants %v and got %v", want, got)
	}
}

func TestDoMath(t *testing.T) {
	got := DoMath(5, 10, Add)
	want := 15

	if got != want {
		log.Fatalf("Error - Wants %v and got %v", want, got)
	}
}
