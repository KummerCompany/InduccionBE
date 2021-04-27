//Ejercicio cinco - Arreglos
//Author: Melina Griffo

package main

import "fmt"

func main() {

	var ages [5]int

	fmt.Println("Escriba la edad de su primer alumno:")
	fmt.Scanln(&ages[0])

	fmt.Println("Escriba la edad de su segundo alumno:")
	fmt.Scanln(&ages[1])

	fmt.Println("Escriba la edad de su tercer alumno:")
	fmt.Scanln(&ages[2])

	fmt.Println("Escriba la edad de su cuarto alumno:")
	fmt.Scanln(&ages[3])

	fmt.Println("Escriba la edad de su quinto alumno:")
	fmt.Scanln(&ages[4])

	fmt.Println("Las edades de sus alumnos son:", ages)

}
