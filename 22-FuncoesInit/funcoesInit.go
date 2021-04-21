package main

import "fmt"

var n int

func main() {
	fmt.Println("executando a funcao main")
	fmt.Println(n)
}

func init() {
	fmt.Println("executando a funcao init") // sera iniciada antes da funcao main
	n = 10
}
