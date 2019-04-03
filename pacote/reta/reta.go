package main

import "math"

/*
	A inicialização de métodos, variáveis e etc com a primeira letra maiúscula ou minúscula possuem significados
	diferentes:

	Primeira letra maiúscula: Público - visível para fora do pacote.
	Primeira letra minúscula: Privado - visível apenas dentro do pacote.

	Ex:
	NomeQualquer - gerará algo público.
	nomeQualquer - gerará algo privado.
	_Ponto 			 - gerará algo privado.

	Cada função deve ser única por pacote.
	Um pacote pode ter N arquivos.
*/

// Representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// função pública - visível fora do pacote.
// calcula a distância linear entre dois pontos.
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
