//Ejercicio ocho - Bucles
//Author: Melina Griffo

package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	var nombreIngresado, contraseñaIngresada string

	nombreUsuario := "joaquin"
	contraseñaPlana := "boss123"
	contraseñaPlanaComoByte := []byte(contraseñaPlana)

	hash, err := bcrypt.GenerateFromPassword(contraseñaPlanaComoByte, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	hashComoCadena := string(hash)

	fmt.Println("Ingrese el nombre de Usuario:")
	fmt.Scanln(&nombreIngresado)
	fmt.Println("Ingrese la contraseña:")
	fmt.Scanln(&contraseñaIngresada)
	if nombreUsuario == nombreIngresado || contraseñaIngresada == hashComoCadena {
		fmt.Println("Usuario logueado corrrectamente")
	} else {
		fmt.Println("Error")
	}
}
