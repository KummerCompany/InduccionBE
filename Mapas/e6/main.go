//Ejercicio seis - Mapas
//Author: Melina Griffo

package main

import "fmt"

func main() {
	var (
		nombre, apellido, calle string
		edad                    uint8
		numero, cod_postal      int16
	)
	calle_interes := "calderon"

	m := make(map[string]interface{})

	fmt.Println("Ingrese su nombre:")
	fmt.Scanln(&nombre)
	fmt.Println("Ingrese su apellido:")
	fmt.Scanln(&apellido)
	fmt.Println("Ingrese su edad:")
	fmt.Scanln(&edad)
	fmt.Println("Ingrese el nombre de su calle:")
	fmt.Scanln(&calle)
	fmt.Println("Ingrese el número de su calle:")
	fmt.Scanln(&numero)
	fmt.Println("Ingrese el código postal:")
	fmt.Scanln(&cod_postal)

	m["name"] = nombre
	m["surname"] = apellido
	m["age"] = edad
	m["address"] = calle
	m["number_address"] = numero
	m["postal_code"] = cod_postal

	if calle == calle_interes || edad >= 18 {
		fmt.Println("Cumple con el requisito de mayoria de edad y/o vive en calle Calderon\n", m)
	}
}
