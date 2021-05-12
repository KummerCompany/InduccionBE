//Ejercicio 13 - Bucles
//Author: Melina Griffo

package main

import "fmt"

func main() {

	year_birth := 1994

	for {
		if year_birth > 2021 {
			break
		}
		fmt.Println(year_birth)
		year_birth++
	}

}
