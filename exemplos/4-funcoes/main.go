package main

import "fmt"

func somar(numero1 int8, numero2 int8) int8 {          // retorno simples
	return numero1 + numero2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {    // multiplos retornos
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}


func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("texto de teste")

	resultadoSoma, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma)

}