//Ejercicio seis
//Author: Melina Griffo

package main

import (
	"fmt"
	"strconv"
)

func main() {
	//EJERCICIO 6a
	a := "123421"

	b, _ := strconv.Atoi(a)

	fmt.Println("String a Int:", b)

	//EJERCICIO 6b

	c := 54.2
	d := int(c)

	fmt.Println("Float a Int", d)

	//EJERCICIO 6c

	e := []byte("14")
	f, _ := strconv.Atoi(string(e))
	fmt.Println("Byte a Int", f)

}
