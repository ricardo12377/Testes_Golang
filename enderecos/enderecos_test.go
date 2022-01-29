package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestQualquer(t *testing.T) {
	if 1 > 2 {
		t.Errorf("O teste quebrou")
	}
}

func TestTipoDeEndereco(t *testing.T) {

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praças das Rosas", "Tipo inválido"},
		{"Estrada qualquer", "Estrada"},
		{"RUA DOS BANCOS", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"", "Tipo inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do espereado %s",
				tipoDeEnderecoRecebido, cenario.retornoEsperado)
		}
	}
}

//para testar todos os pacotes, usar o codigo na pasta raiz go test ./...
//go test -v para um teste mais descritivo
//t.Parallel dentro de uma função de teste faz a função rodar em paralelo, tem que usar em mais de uma para fazer diferença
// go test --cover para saber a porcentagem do codigo coberto por teste
//go test --coverprofile cobertura.txt , irá criar um arquivo com as informações do teste
//para ler ele : go tool cover --func=resultado.txt
//saber linha que nao esta coberta go tool cover --html=cobertura.txt
