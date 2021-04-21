package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int) (int, int, int, int) {
	soma := n1 + n2
	subtracao := n1 - n2
	multiplicacao := n1 * n2
	divisao := n1 / n2

	return soma, subtracao, multiplicacao, divisao
}

func main() {
	soma := somar(10, 32)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	//retornando todos os valores da funcao
	/* resultadoSoma, resultadoSubtracao, resultadoMultiplicacao, resultadoDivisao := calculosMatematicos(4, 2)
	fmt.Println(resultadoSoma, resultadoSubtracao, resultadoMultiplicacao, resultadoDivisao) */

	//retornando apenas subtracao e multiplicacao usando underline_ para esconder soma e divisao
	_, resultadoSubtracao, resultadoMultiplicacao, _ := calculosMatematicos(4, 2)
	fmt.Println(resultadoSubtracao, resultadoMultiplicacao)

}
