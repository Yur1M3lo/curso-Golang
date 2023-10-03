package main

import "fmt"

func calculosMatematicos(n1, n2 int8) (soma int8, sub int8) {
	soma = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	soma, sub := calculosMatematicos(10, 20)
	fmt.Println(soma, sub)

}