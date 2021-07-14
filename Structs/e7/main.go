//Ejercicio siete - Structs
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
		value           string
		val             uint8
		data            int16
		number          int
		datos           person
		datos_completos []person
	)

	fmt.Println("Los datos de cuántas personas desea cargar:")
	fmt.Scanln(&number)

	for i := 0; i < number; i++ {

		fmt.Printf("Ingrese el nombre %d:", i+1)
		fmt.Scanln(&value)
		datos.nombre = value

		fmt.Printf("Ingrese el apellido %d:", i+1)
		fmt.Scanln(&value)
		datos.apellido = value

		fmt.Printf("Ingrese la edad %d:", i+1)
		fmt.Scanln(&val)
		datos.edad = val

		fmt.Printf("Ingrese el nombre de la calle %d:", i+1)
		fmt.Scanln(&value)
		datos.calle = value

		fmt.Printf("Ingrese el número de la calle %d:", i+1)
		fmt.Scanln(&data)
		datos.numero = data

		fmt.Printf("Ingrese el código postal %d:", i+1)
		fmt.Scanln(&data)
		datos.cod_postal = data

		datos_completos = append(datos_completos, datos)

	}

	for i, v := range datos_completos {
		fmt.Printf("\t Los datos de la persona %d son: % v\n", i+1, v)
	}
}
