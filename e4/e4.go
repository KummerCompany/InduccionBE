package main

import "fmt"

func main() {
	var nombre string
	fmt.Println("Ingrese su nombre: ")
	fmt.Scanf("%s\n", &nombre)
	fmt.Printf("Hola %s\n", nombre)

}
