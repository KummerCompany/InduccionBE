//Ejercicio dos - Mapas
//Author: Melina Griffo

package main

import "fmt"

func main() {
	var nombre, apellido, calle, ciudad string

	m := make(map[string]string)

	fmt.Println("Ingrese su nombre:")
	fmt.Scanln(&nombre)
	fmt.Println("Ingrese su apellido:")
	fmt.Scanln(&apellido)
	fmt.Println("Ingrese su domicilio:")
	fmt.Scanln(&calle)
	fmt.Println("Ingrese su ciudad:")
	fmt.Scanln(&ciudad)

	m["name"] = nombre
	m["surname"] = apellido
	m["address"] = calle
	m["city"] = ciudad

	fmt.Println(m)
}
