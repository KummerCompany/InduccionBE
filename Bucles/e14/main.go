//Ejercicio 14 - Bucles
//Author: Melina Griffo

package main

import "fmt"

func main() {
	constant := 4
	for i := 10; i <= 100; i++ {
		fmt.Printf(" El resto de %v dividido 4 es %v\n", i, i%constant)
	}

}
