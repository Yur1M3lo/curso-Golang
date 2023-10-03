package main

import "fmt"

func main() {
	/*i := 0

	for i < 10 {
		i++
		fmt.Println("incrementando i")
		fmt.Println(i)
	}*/

	/*for j := 0; j < 10; j++ {
		fmt.Println("Segundo for")
	}*/

	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, nome:= range nomes {
		fmt.Println(indice, nome)
	}

	usuario := map[string]string {
		"nome": "exemplo",
		"sobrenome": "sobrenome",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}


}