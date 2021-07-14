//Ejercicio seis - Bucles
//Author: Melina Griffo

package main

import "fmt"

func main() {

	var number int
	var posicion1 int
	var posicion2 int

	num := make([]int, 5)

	for i := 0; i < 5; i++ {

		fmt.Println("Ingrese un número:")
		fmt.Scanln(&number)
		num[i] = number

	}
	fmt.Println("los números ingresados son:", num)
	fmt.Println("Ingrese primera posición:")
	fmt.Scanln(&posicion1)
	fmt.Println("Ingrese segunda posición:")
	fmt.Scanln(&posicion2)
	posiciona := num[posicion1]
	posicionb := num[posicion2]

	posiciona, posicionb = posicionb, posiciona

	num[posicion1] = posiciona
	num[posicion2] = posicionb

	fmt.Println("Su nuevo arreglo es:", num)
}
