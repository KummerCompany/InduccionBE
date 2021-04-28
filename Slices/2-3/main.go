//Ejercicio tres - Slices
//Author: Melina Griffo

package main

import "fmt"

func main() {

	var name string
	var last_name string

	fmt.Println("Ingrese su nombre: ")
	fmt.Scanln(&name)

	fmt.Println("Ingrese su apellido: ")
	fmt.Scanln(&last_name)

	fmt.Println(name + last_name)

}

//Ejercicio dos - Slices
//Author: Melina Griffo

/*Array es un tipo de dato que permite el almacenamiento de elementos de un mismo tipo. Además
 ésta estructura tiene un tamaño predefinido, es decir que una vez definido no se puede modificar.

 En cambio un slice permite almacenar un conjunto de datos del mismo tipo, pero la cantidad de
 elementos de la misma puede crecer a lo largo de la ejecución del programa.

 Por ultimo tenemos el caso de listas, donde la diferencia más evidente es que sus elementos no se
 encuentran en espacios contiguos de memoria (como los arreglos), sin embargo cada elemento o "nodo"
 conoce la ubicación de su próximo valor. Una ventaja de las listas sobre los arreglos es que,
 al igual que los slices, pueden cambiar su tamaño de forma dinámica permitiendo aumentar o disminuir
su tamaño.*/

//Fuente: https://autonomia.digital/es/2020/06/01/introduction-data-structures.html
