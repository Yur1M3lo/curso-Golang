package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalPonteiros(numero *int) {
	*numero = *numero * -1
}

func main() {

	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalPonteiros(&novoNumero)
	fmt.Println(novoNumero)


}