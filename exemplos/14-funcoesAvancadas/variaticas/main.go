package main

import "fmt"

func soma(numeros ...int) int {
	var resultado int

	for _, valor := range numeros {
		resultado += valor
	}

	return resultado
}

func main() {
	result := soma(1, 1, 1, 1, 1)
	fmt.Println("Resultado: ", result)

}