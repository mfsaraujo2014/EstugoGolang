package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a funcao1")
}

func funcao2() {
	fmt.Println("Executando a funcao2")
}

func alunoAprovado(n1, n2 int) bool {
	defer fmt.Println("Media calculada. Resultado retornado!")
	fmt.Println("Entrando na funcao para ver se o aluno foi aprovado")

	media := n1 + n2/2

	if media >= 7 {
		return true
	}
	return false
}

func main() {
	//DEFER = adiar
	//defer funcao1()
	//funcao2()
	//com o defer na frente ele executa primeiro a funcao2

	fmt.Println(alunoAprovado(7, 8))
}
