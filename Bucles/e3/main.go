//Ejercicio tres - Bucles
//Author: Melina Griffo

package main

import "fmt"

func main() {
	var value int
	var longitud int

	fmt.Print("¿Cuántos números del 1 al 10 desea ingresar? ")
	fmt.Scanf("%d\n", &longitud)

	val := make([]int, longitud)

	for i := 0; i < longitud; i++ {
		fmt.Printf("\nIngrese su número %d: ", i+1)
		fmt.Scanf("%d\n", &value)
		val[i] = value

		if value == 21 {
			sumatoria := 0
			for _, val := range val {
				sumatoria += val
			}

			promedio := sumatoria / len(val)
			fmt.Printf("El promedio de los números ingresados es: %d\n", promedio)
		}
	}

}
