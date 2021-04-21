package main

import "fmt"

func main() {
	//ARITMETICOS
	soma := 2 + 3
	subtracao := 3 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 4
	restoDivisao := 10 % 4

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)
	//FIM ARITMETICOS

	//ATRIBUICAO
	var variavel string = "String"
	variavel2 := "String2"
	fmt.Println(variavel, variavel2)
	//FIM DOS OPERADORES DE ATRIBUICAO

	//OPERADORES RELACIONAIS
	var n1 int = 1
	var n2 int = 2
	fmt.Println(n1 > n2)
	fmt.Println(n1 < n2)
	fmt.Println(n1 >= n2)
	fmt.Println(n1 <= n2)
	fmt.Println(n1 == n2)
	fmt.Println(n1 != n2)
	//FIM DOS RELACIONAIS

	//OPERADORES LOGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	//FIM DOS OPERADORES LOGICOS

	//OPERADORES UNARIOS
	numero := 10
	numero++
	numero += 15
	numero -= 15
	numero *= 15
	numero /= 15
	numero %= 15
	fmt.Println(numero)
	//FIM DOS UNARIOS

}
