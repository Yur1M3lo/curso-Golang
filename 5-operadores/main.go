package main

import "fmt"

func main() {
	// ARITMETICOS
	soma := 1 + 5
	sub := 10 - 5
	divisao := 40 / 4
	multi := 3 * 10
	resto := 10 % 2
	fmt.Println(soma, sub, divisao, multi, resto)

	// ATRIBUIÇAO
	var variavel string = "texto"
	varivavel2 := "texto2"
	fmt.Println(variavel, varivavel2)

	// RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// LOGICOS
	fmt.Println("-------------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// UNARIOS
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)
	
	numero--
	numero -= 5
	
	numero *= 3
	numero /= 2
	fmt.Println(numero)

}