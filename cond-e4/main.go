//Ejercicio 4 - Condicionales
//Author: Melina Griffo

/*"message": "cannot compare a > b (mismatched types int64 and float64)"
debemos tener los mismos tipos de datos
transformar el float a int o viceversa
como se muestra a continuacion: */

package main

import "fmt"

func main() {

	var a int64
	var b float64

	fmt.Println("Ingrese un numero entero: ")
	fmt.Scanf("%d/n", &a)

	fmt.Println("Ingrese un numero decimal: ")
	fmt.Scanf("%f/n", &b)

	c := int64(b)

	if a > c {
		fmt.Println("El numero mas grande es: ", a)
	} else {
		fmt.Println("El numero mas grande es:", c)
	}

}
