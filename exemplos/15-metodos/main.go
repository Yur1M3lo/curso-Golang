package main

import "fmt"

type usuario struct {
	nome  string
	idade int8
}

func (u usuario) salvar() {
	fmt.Println("Dentro do metodo salvar")
}

func main() {
	usuario1 := usuario{"usuario1", 23}
	fmt.Println(usuario1)
	usuario1.salvar()


}