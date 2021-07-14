//Ejercicio tres - Funciones
//Author: Melina Griffo

package main

import "fmt"

func main() {
	y := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	z := bar(y)
	fmt.Println(z)
}

func bar(x []int) int {
	suma := 0

	for _, v := range x {
		suma += v
	}
	return suma
}
