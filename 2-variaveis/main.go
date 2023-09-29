package main

import "fmt"

func main() {
	var variavel1 string = "exemplo 1" // declaração explicita
	variavel2 := "exemplo 2" // declaração implicita

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var ( 								// multipla declaração explicita
		variavel3 string = "exemplo 3"
		variavel4 string = "exemplo 4"
	) 
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "exemplo 5", "exemplo 6" // multipla declaração implicita
	fmt.Println(variavel5, variavel6)

	const constante1 string = "constante 1" // declaração de constante
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5 // invertendo valores das variaveis 
	fmt.Println(variavel5, variavel6)
	
}