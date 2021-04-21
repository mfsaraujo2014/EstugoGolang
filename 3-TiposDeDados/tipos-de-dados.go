package main

import (
	"errors"
	"fmt"
)

func main() {
	//numeros inteiros
	var numero int64 = 1000000000
	fmt.Println(numero)

	var numero2 uint32 = 100000
	fmt.Println(numero2)

	//alias
	//INT32 = RUNE
	var numero3 rune = 12345
	fmt.Println(numero3)

	//BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)
	//Fim numeros inteiros

	//Numeros reais
	var numeroReal1 float32 = 123.78
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123333.78
	fmt.Println(numeroReal2)

	numeroReal3 := 8328.4343
	fmt.Println(numeroReal3)
	//Fim numeros reais

	//Strings
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)
	//FIM STRINGS

	//Booleano
	var booleano1 bool = true
	fmt.Println(booleano1)

	var booleano2 bool = false
	fmt.Println(booleano2)

	var booleano3 bool
	fmt.Println(booleano3)
	//Fim booleano

	//Erro
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
