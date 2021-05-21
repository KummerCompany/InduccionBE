//Ejercicio uno - Mapas
//Author: Melina Griffo

package main

import "fmt"

func main() {

	m := map[string][]string{
		"pedro_lopez":     {"montañismo", "computadoras", "caminar"},
		"carlos_uribe":    {"leer", "comprar", "musica"},
		"mariela_nicolas": {"surf", "pintar", "cantar"},
	}
	for key, value := range m {
		fmt.Println(key)

		for i, v := range value {
			fmt.Println(i, v)
		}
	}

}
