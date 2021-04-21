package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for i < 10 {
		i++
		fmt.Println("Incrementando i ")
		time.Sleep(time.Second)
	}
	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando J: ", j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"joão", "jose", "matheus"}

	for i, elemento := range nomes {
		fmt.Println(i, elemento)
	}

	for i, letra := range "PALAVRA" {
		fmt.Println(i, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
