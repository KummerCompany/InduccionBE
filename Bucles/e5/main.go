//Ejercicio cinco - Bucles
//Author: Melina Griffo

package main

import "fmt"

func main() {
	var value int
	val := make([]int, 10)

	for i := 0; i < 10; i++ {
		fmt.Printf("\nIngrese su número %d: ", i+1)
		fmt.Scanf("%d\n", &value)
		val[i] = value
	}

	max_pos := val[0]
	for index, v := range val {
		if val[max_pos] < v {
			max_pos = index

		}

	}
	fmt.Printf("El número mas grande es:%v que está en la posición %v\n", val[max_pos], max_pos+1)
}
