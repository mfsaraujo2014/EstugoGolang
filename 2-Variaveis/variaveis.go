package main

import "fmt"

func main() {
	var variavel string = "Variável 1"
	variavel2 := "Variável 2"
	fmt.Println(variavel)
	fmt.Println(variavel2)

	var (
		variavel3 string = "Variável 3"
		variavel4 string = "Variável 4"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variável5", 15
	fmt.Println(variavel5, variavel6)

	//inverter valor de variaves sem aux
	variavel4, variavel5 = variavel5, variavel4
	fmt.Println(variavel4, variavel5)
}
