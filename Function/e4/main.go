//Ejercicio cuatro - Funciones
//Author: Melina Griffo

package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	edad     int
}

func (p persona) presentar() {
	fmt.Println("Me llamo", p.nombre, p.apellido, "y tengo", p.edad, "años")
}

func main() {
	p1 := persona{
		nombre:   "Maria",
		apellido: "Perez",
		edad:     23,
	}
	p1.presentar()
}
