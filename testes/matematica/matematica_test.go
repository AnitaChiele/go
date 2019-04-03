package matematica

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."

/*
	Para rodar os comandos de testes e ser um teste válido é necessário seguir
	as convenções:
		Os arquivos de testes terminam com _test.go
		O nome das funções de testes devem começar com Test
	Se essas duas convenções não forem seguidas, o GO não entende que são funções
	ou, arquivos de testes e ao rodar o comando de teste, ele não executa o teste.
*/

func TestMedia(t *testing.T) {
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}
