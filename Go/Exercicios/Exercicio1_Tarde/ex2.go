package main

import (
	"fmt"
)

var clienteIdade int = 23
var empregado bool = true
var anoAtividade int = 2
var salario int = 90000

func main() {
	if clienteIdade < 22 || empregado == false || anoAtividade < 2 {
		fmt.Println("Nao é possivel fazer o empréstimo")
	} else {
		if salario < 100000 {
			//juros := true
			fmt.Println("O empréstimo foi aprovado, mas está sujeito a juros")
		} else {
			fmt.Println("O empréstimo foi aprovado e nao está sujeito a juros")
		}
	}
}
