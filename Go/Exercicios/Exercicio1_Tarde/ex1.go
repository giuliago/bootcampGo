package main

import (
	"fmt"
)

func main() {
	palavra := "teste"
	numLetras := len(palavra)

	fmt.Printf("A palavra %s contém %d letras.\n", palavra, numLetras)

	fmt.Println("Letras da palavra:")
	for _, letra := range palavra {
		fmt.Printf("%c\n", letra)
	}

}
