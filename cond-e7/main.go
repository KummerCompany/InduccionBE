//Ejercicio siete - Condicionales
//Author: Melina Griffo

package main

import "fmt"

func main() {

	var x, t float64
	count_max_days := float64(10)
	count_max_km := float64(500)
	price_km := 0.19
	discount := 0.80

	fmt.Println("Ingrese la distancia en km que hay entre las 2 ciudades")
	fmt.Scanln(&x)

	fmt.Println("Ingrese los días de estancia en la ciudad")
	fmt.Scanln(&t)

	d := x * 2
	p1 := d * price_km * discount
	p2 := d * price_km

	if d > count_max_km && t > count_max_days {

		fmt.Println("El precio de su billete es: $", p1)

	} else {

		fmt.Println("El precio de su billete es: $", p2)

	}

}
