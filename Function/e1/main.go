//Ejercicio uno - Funciones
//Author: Melina Griffo
package main

import "fmt"

func main() {
	x := foo()

	y, z := bar()
	fmt.Println(x, y, z)
}

func foo() int {
	return 46
}

func bar() (int, string) {
	return 27, "Hola Mundo"

}
