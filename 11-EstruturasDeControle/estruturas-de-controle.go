package main

import "fmt"

func main() {
	numero := 10
	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 { //criar uma variavel no ifInit limita ela ao escopo do If
		fmt.Println("Numero é maior que 0")
	} else {
		fmt.Println("Numero é menor que 0")
	}
}
