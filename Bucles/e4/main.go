//Ejercicio cuatro - Bucles
//Author: Melina Griffo

package main

import "fmt"

func main() {

	var name string
	var age int
	var names []string
	var ages []int

	for {

		fmt.Println("Ingrese su nombre")
		fmt.Scanln(&name)

		if name == "EXIT" {
			break
		}
		names = append(names, name)
		fmt.Println("Ingrese su edad")
		fmt.Scanln(&age)
		ages = append(ages, age)

	}

	min_pos, max_pos := 0, 0

	for index, a := range ages {
		if ages[min_pos] > a {
			min_pos = index
		}
		if ages[max_pos] < a {
			max_pos = index

		}

	}
	fmt.Printf("La persona de menor edad es: %v. \nY el de mayor edad es: %v\n", names[min_pos], names[max_pos])
}
