package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	Raio float64
}
type ScretAgent struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func (c *Circulo) Escala(fator float64) {
	c.Raio *= fator
}

func main() {

	c1 := Circulo{Raio: 5}

	fmt.Println(c1.Area())

	c2 := Circulo{Raio: 3.5}

	c2.Escala(1.5)

	fmt.Println(c2.Raio)
}
