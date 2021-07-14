//Ejercicio cinco - Structs
//Author: Melina Griffo

package main

import "fmt"

type person struct {
	nombre, apellido, calle string
	edad                    uint8
	numero, cod_postal      int16
}

func main() {

	var (
		value string
		val   uint8
		data  int16
	)

	datos := new(person)

	fmt.Println("Ingrese su nombre:")
	fmt.Scanln(&value)
	datos.nombre = value

	fmt.Println("Ingrese su apellido:")
	fmt.Scanln(&value)
	datos.apellido = value

	fmt.Println("Ingrese su edad:")
	fmt.Scanln(&val)
	datos.edad = val

	fmt.Println("Ingrese el nombre de su calle:")
	fmt.Scanln(&value)
	datos.calle = value

	fmt.Println("Ingrese el número de su calle:")
	fmt.Scanln(&data)
	datos.numero = data

	fmt.Println("Ingrese el código postal:")
	fmt.Scanln(&data)
	datos.cod_postal = data

	fmt.Println(datos)
}
