//Ejercicio nueve - Mapas
//Author: Melina Griffo

package main

import "fmt"

func main() {
	var datos int
	var nombre string
	nombres := make(map[string]int)

	fmt.Println("Los nombres de cuántas personas desea cargar:")
	fmt.Scanln(&datos)

	for i := 0; i < datos; i++ {

		fmt.Printf("Ingrese nombre %d:", i+1)
		fmt.Scanln(&nombre)
		// Comentario y resolucion de Santi:
		// analizamos si el nombre ya existia en el mapa
		// si existe, le aumentamos en uno el valor
		// si no existe, lo creamos y le guardamos un 1
		// haciendo referencia a que se nombro una vez
		if val, ok := nombres[nombre]; ok {
			nombres[nombre] = val + 1

		} else {
			nombres[nombre] = 1

		}

	}
	fmt.Println(nombres)
}
