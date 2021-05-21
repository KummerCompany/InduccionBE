//Ejercicio ocho - Mapas
//Author: Melina Griffo

package main

import "fmt"

func main() {

	var (
		nombre, apellido, calle    string
		edad                       uint8
		numero, cod_postal         int16
		datos                      int
		nombres, apellidos, calles []string
		edades                     []uint8
		numeros, cod_postales      []int16
	)

	fmt.Println("Los datos de cuántas personas desea cargar:")
	fmt.Scanln(&datos)

	for i := 0; i < datos; i++ {

		fmt.Printf("Ingrese nombre %d:", i+1)
		fmt.Scanln(&nombre)
		nombres = append(nombres, nombre)

		fmt.Printf("Ingrese apellido %d:", i+1)
		fmt.Scanln(&apellido)
		apellidos = append(apellidos, apellido)

		fmt.Printf("Ingrese edad %d:", i+1)
		fmt.Scanln(&edad)
		edades = append(edades, edad)

		fmt.Printf("Ingrese nombre de calle %d:", i+1)
		fmt.Scanln(&calle)
		calles = append(calles, calle)

		fmt.Printf("Ingrese número de calle %d:", i+1)
		fmt.Scanln(&numero)
		numeros = append(numeros, numero)

		fmt.Printf("Ingrese código postal %d:", i+1)
		fmt.Scanln(&cod_postal)
		cod_postales = append(cod_postales, cod_postal)

	}
	min_pos, max_pos := 0, 0

	for index, a := range edades {
		if edades[min_pos] > a {
			min_pos = index
		}
		if edades[max_pos] < a {
			max_pos = index

		}

	}
	fmt.Println(nombres, apellidos, edades, calles, numeros, cod_postales)
	fmt.Printf("La persona de menor edad es: %v. \nY el de mayor edad es: %v\n", nombres[min_pos], nombres[max_pos])
}
