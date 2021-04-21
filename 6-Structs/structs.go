package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint
}

func main() {
	var u usuario
	u.nome = "Matheus"
	u.idade = 20
	fmt.Println(u)

	endereco2 := endereco{"rua dos bobos", 0}
	u2 := usuario{"inacio", 17, endereco2}
	fmt.Println(u2)

	u3 := usuario{nome: "vitoria"}
	fmt.Println(u3)
}
