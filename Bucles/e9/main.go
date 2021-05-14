//Ejercicio nueve - Bucles
//Author: Melina Griffo

package main

import "fmt"

func main() {
	for stock := 1000; stock >= 200; stock-- {
		fmt.Printf("Pedido realizado, quedan %d unidades en stock\n ", stock)
		if stock == 200 {
			fmt.Println("Quedan pocas unidades")
		}
	}
}
