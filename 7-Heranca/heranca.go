package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint
	altura    uint
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"Matheus Felipe", "de Souza Araújo", 20, 178}
	fmt.Println(p1)

	e1 := estudante{p1, "engenharia de software", "Universidade Federal do Ceará"}
	fmt.Println(e1)
	fmt.Println("nome:" + e1.nome)
}
