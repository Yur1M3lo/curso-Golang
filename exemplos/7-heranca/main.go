/*
	Golang não é orientada a objetos então não possui o conceito de herança,
	mas oference uma ferramenta para trabalhar com structs
*/
package main

import "fmt"

type pessoa struct {
	nome 		string
	sobrenome 	string 
	idade 		uint8
}

type aluno struct {
	pessoa
	curso 		string 
	faculdade 	string 
}


func main() {
	fmt.Println("herança")

	p1 := pessoa{"yuri", "melo", 23}
	fmt.Println(p1)

	a1 := aluno{p1, "Engenharia", "UFRJ"}
	fmt.Println(a1)
	fmt.Println(a1.nome)
	fmt.Println(a1.faculdade)


}