//Ejercicio cinco
//Author: Melina Griffo

package main

import (
	"fmt"
)

func main() {
	var lado float64

	fmt.Println("Ingrese la medida de un lado del cuadrado: ")
	fmt.Scanf("%f\n", &lado)

	c := lado * 4

	fmt.Println("El Perímetro de su cuadrado es:", c)

	d := lado * lado

	fmt.Println("El Área de su cuadrado es:", d)
}
