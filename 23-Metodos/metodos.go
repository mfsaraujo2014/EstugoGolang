package main

import "fmt"

type usuario struct {
	nome  string
	idade uint
}

func (u usuario) salvar() {
	fmt.Printf("Salvando dados do usuário %s no banco de dados!\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	if u.idade >= 18 {
		return true
	}
	return false
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"matheus", 20}
	usuario1.salvar()
	maiorDeIdade := usuario1.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2 := usuario{"inacio", 17}
	usuario2.salvar()
	maiorDeIdade2 := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade2)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
