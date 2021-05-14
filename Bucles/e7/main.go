//Ejercicio siete - Bucles
//Author: Melina Griffo

package main

import "fmt"

func main() {

	var number int

	num := make([]int, 5)

	for i := 0; i < 5; i++ {

		fmt.Println("Ingrese un número:")
		fmt.Scanln(&number)
		num[i] = number
	}
	fmt.Println("Los números ingresados son:", num)

	for i := 0; i < len(num); i++ {
		for j := 0; j < len(num)-i-1; j++ {
			if num[j] > num[j+1] {
				new := num[j]
				num[j] = num[j+1]
				num[j+1] = new

			}
		}
	}
	fmt.Println("Los números ingresados en orden son:", num)
}
