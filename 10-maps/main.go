package main

import "fmt"

func main() {

	usuario := map[string]string {
		"nome": "Yuri",
		"sobrenome": "Melo",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string {
		"nome": {
			"primeiro": "Yuri",
			"ultimo": "Melo",
		},
		"curso": {
			"nome": "Matemática",
			"campus": "Campus III",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "curso")
	fmt.Println(usuario2)
	
	usuario2["endereco"] = map[string]string{
		"rua": "abc",
		"numero": "123",
	}
	fmt.Println(usuario2)
}