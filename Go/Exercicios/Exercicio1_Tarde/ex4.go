package main

import "fmt"

var employees = map[string]int{
	"Benjamin": 20,
	"Manuel":   26,
	"Brenda":   19,
	"Dario":    44,
	"Pedro":    30,
}
var func20 int = 0

func main() {
	employees["Federico"] = 25
	delete(employees, "Pedro")
	for nome, idade := range employees {
		if nome == "Benjamin" {
			fmt.Printf("A idade de %s é %d \n", nome, idade)
		}
		if idade >= 20 {
			func20++
		}
	}
	fmt.Printf("A empresa possui %d funcionários com mais de 20 anos \n", func20)
}
