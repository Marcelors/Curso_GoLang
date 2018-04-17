package main

import (
	"fmt"
)

func main() {

	// var aprovados map[int]string
	// map devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[12345678900] = "Pedro"
	aprovados[95135745682] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (Cpf: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[95135745682])
	delete(aprovados, 95135745682)
	fmt.Println(aprovados[95135745682])
}
