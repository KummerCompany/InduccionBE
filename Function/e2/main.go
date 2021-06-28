//Ejercicio dos - Funciones
//Author: Melina Griffo

package main

import "fmt"

func main() {

	y := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	z := foo(y...)
	fmt.Println(z)

}

func foo(x ...int) int {
	suma := 0

	for _, v := range x {
		suma += v
	}
	return suma
}
