//Ejercicio dos - Structs
//Author: Melina Griffo

package main

import "fmt"

type User struct {
	nombre, apellido, calle, ciudad string
}

func main() {

	var value string

	datos := new(User)

	fmt.Println("Ingrese su nombre:")
	fmt.Scanln(&value)
	datos.nombre = value

	fmt.Println("Ingrese su apellido:")
	fmt.Scanln(&value)
	datos.apellido = value

	fmt.Println("Ingrese su domicilio:")
	fmt.Scanln(&value)
	datos.calle = value

	fmt.Println("Ingrese su ciudad:")
	fmt.Scanln(&value)
	datos.ciudad = value

	fmt.Println(datos)

}
