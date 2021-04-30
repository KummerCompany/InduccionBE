//Ejercicio cuatro - Arreglos
//Author: Melina Griffo

package main

import "fmt"

func main() {

	//a

	a := [3]int{1, 2, 3}
	b := [3]int{4, 5, 6}
	c := [6]int{a[0], a[1], a[2], b[0], b[1], b[2]}

	fmt.Println(c)

	//b

	A := [2]int{2, 3}
	B := [2]int{2, 3}
	C := [2]int{A[0] + B[0], A[1] + B[1]}

	fmt.Println(C)

	//c

	A = [2]int{2, 3}
	B = [2]int{2, 3}
	C = [2]int{A[0] - B[0], A[1] - B[1]}

	fmt.Println(C)

}
