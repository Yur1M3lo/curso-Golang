package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

	var array[5] string
	array[0] = "Pos 1"
	fmt.Println(array)

	array2 := [5]string{"Pos 1", "Pos 2", "Pos 3", "Pos 4", "Pos 5"}
	fmt.Println(array2)

	array3 := [...]int8{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice)

	slice = append(slice, 11)
	fmt.Println(slice)
	
	slice2 := array2[1:3]
	fmt.Println(slice2)
	
	slice2[1] = "Posição alterada"
	fmt.Println(slice2)


	// Arrays internos
	fmt.Println("---------------------")
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	

}