//Ejercicio uno - Structs
//Author: Melina Griffo

package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	act_fav  []string
}

func main() {
	p1 := persona{
		nombre:   "Pedro",
		apellido: "Lopez",
		act_fav: []string{
			"montañismo",
			"computadoras",
			"caminar",
		},
	}
	p2 := persona{
		nombre:   "Carlos",
		apellido: "Uribe",
		act_fav: []string{
			"leer",
			"comprar",
			"musica",
		},
	}
	p3 := persona{
		nombre:   "Mariela",
		apellido: "Nicolas",
		act_fav: []string{
			"surf",
			"pintar",
			"cantar",
		},
	}
	fmt.Println(p1.nombre)
	fmt.Println(p1.apellido)

	for i, v := range p1.act_fav {
		fmt.Println("\t", i, v)
	}

	fmt.Println(p2.nombre)
	fmt.Println(p2.apellido)

	for i, v := range p2.act_fav {
		fmt.Println("\t", i, v)
	}

	fmt.Println(p3.nombre)
	fmt.Println(p3.apellido)

	for i, v := range p3.act_fav {
		fmt.Println("\t", i, v)
	}
}
