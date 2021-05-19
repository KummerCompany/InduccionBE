//Ejercicio cuatro - Mapas
//Author: Melina Griffo

package main

import "fmt"

func main() {
	var nombre, apellido, calle, ciudad, idioma string

	var a string = "español"
	var b string = "ingles"
	var c string = "portugues"

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

	fmt.Println("Ingrese su idioma")
	fmt.Scanln(&idioma)

	if idioma == a || idioma == b || idioma == c {
		m["language"] = idioma
		fmt.Println(m)
	} else {
		fmt.Println("err")
	}
}
