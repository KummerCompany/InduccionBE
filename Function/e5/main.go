//Ejercicio cinco - Funciones
//Author: Melina Griffo

package main

import (
	"fmt"
	"math"
)

type circulo struct {
	radio float64
}

type cuadrado struct {
	lado float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.radio, 2)
}

func (c cuadrado) area() float64 {
	return math.Pow(c.lado, 2)
}

type forma interface {
	area() float64
}

func info(f forma) {
	fmt.Println(f.area())
}
func main() {

	circ_1 := circulo{
		radio: 12.3,
	}

	cuad_1 := cuadrado{
		lado: 25.5,
	}

	info(circ_1)
	info(cuad_1)
}
