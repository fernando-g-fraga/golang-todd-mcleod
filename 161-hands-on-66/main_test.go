package main

import (
	"log"
	"testing"
)

func TestOpenfile(t *testing.T) {
	got := openfile("teste.txt")
	want := "asdasasdasd"

	if got != want {
		log.Fatalln("Ocorreu um erro ao abrir o arquivo.")
	}
}
