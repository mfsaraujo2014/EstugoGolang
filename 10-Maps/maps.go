package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	//Aninhando maps
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiroNome": "José",
			"ultimoNome":   "Souza",
		},
		"curso": {
			"nome":   "engenharia",
			"campus": "Quixada",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome") //apagando uma chave
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{ //adicionando chave
		"nome": "Gêmeos",
	}
	fmt.Println(usuario2)
}
