package main

import (
	"errors"
	"fmt"
)

func main()  {
	// --------------numeros inteiros--------------
	// int8, int16, int32, int64
	// uint unsyg 

	var numero int16 = 100
	fmt.Println(numero)
	
	// alias
	var numero2 rune = 1234
	fmt.Println(numero2)

	var numero3 byte = 12 // BYTE UINTS
	fmt.Println(numero3)



	// --------------numeros reais--------------
	// float32, float64 

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1240000000.432442
	fmt.Println(numeroReal2)



	// --------------string--------------
	var str string = "Texto"
	fmt.Println(str)

	// --------------booleanos--------------
	var booleano bool = true
	fmt.Println(booleano)


	//--------------error--------------
	var erro error = errors.New("ERRO INTERNO")
	fmt.Println(erro)
}