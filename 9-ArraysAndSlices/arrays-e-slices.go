package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array1 [5]int
	fmt.Println(array1)

	array2 := [5]string{"posicao 1", "posicao 2", "posicao 3", "posicao 4", "posicao 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	fmt.Println(array3)

	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice)

	slice = append(slice, 7)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "posicao alterada"
	fmt.Println(slice2)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	//Arrays Internos
	fmt.Println("---------Arrays Internos-----------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4)) //length
	fmt.Println(cap(slice4)) //capacidade
}
