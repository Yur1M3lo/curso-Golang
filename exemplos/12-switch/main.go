package main

import "fmt"

func diaSemana(numero int8) string {
	switch numero {
	case 1:
		return "Segunda-feira"
	case 2:
		return "Terça-feira"
	case 3:
		return "Quarta-feira"
	case 4:
		return "Quinta-feira"
	case 5:
		return "Sexta-feira"
	case 6:
		return "Sabado"
	case 7:
		return "Domingo"
	default:
		return "Número invalido"
	}
}

func diaSemana2(numero int8) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Segunda"
	case numero == 2:
		diaDaSemana = "Terça"
	case numero == 3:
		diaDaSemana = "Quarta"
	case numero == 5:
		diaDaSemana = "Quinta"
	case numero == 6:
		diaDaSemana = "Sexta"
	case numero == 7:
		diaDaSemana = "Sabado"
	case numero == 8:
		diaDaSemana = "Domingo"
	default:
		diaDaSemana = "Numero invalido"
	}

	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	dia := diaSemana(3)
	fmt.Println(dia)
	
	dia2 := diaSemana2(5)
	fmt.Println(dia2)
}