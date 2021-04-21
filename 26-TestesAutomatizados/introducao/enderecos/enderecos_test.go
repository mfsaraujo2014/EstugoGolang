package enderecos_test

import (
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()

	cenarioDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Rua DEF", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rua Dos Bobos", "Rua"},
		{"Avenida Dantas", "Avenida"},
		{"Estrada Inacio", "Estrada"},
		{"Avenida Principal", "Avenida"},
		{"", "Tipo inválido"},
		{"Rua Matheus", "Rua"},
		{"Praça Zé", "Tipo inválido"},
	}

	for _, cenario := range cenarioDeTeste {
		tipoEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("Tipo recebido %s é diferente do tipo esperado %s", tipoEnderecoRecebido, cenario.retornoEsperado)
		}
	}
}

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Errorf("Teste Quebrou")
	}
}

/*COMANDOS PARA TESTE
go test
go test -v
go test cover
go test coverprofile relatorio.txt
go tool cover --func=relatorio.txt
go tool cover --html=relatorio.txt
*/
