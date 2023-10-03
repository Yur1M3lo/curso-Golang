package main

import "fmt"

type usuario struct {
	nome string
	idade int8
}

func main() {
	fmt.Println("Arquivo Struct")

	var u1 usuario
	u1.nome = "Yuri"
	u1.idade = 23
	fmt.Println(u1)

	u2 := usuario{"Shzam", 23}
	fmt.Println(u2)
}