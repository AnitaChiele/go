package main

import (
	f "fmt"
	"time"
)

func fale(pessoa, texto string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		f.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	/*
		go é uma palavra reservada que pode ser utilizada para rodar
		a concorrência.
	*/

	f.Println("=============================================")
	f.Println("Chamada default:")
	fale("Maria", "Pq vc não fala comigo?", 3)
	fale("João", "Só posso falar depois de vc!", 1)
	f.Println("=============================================")

	/*
		O código abaixo não exibirá nenhuma msg de output.
		Isso acontece pq a go routine dura enquanto a thread
		principal estiver sendo executado. Ou seja, o código
		da função fale nem foi executado porque o método main
		terminou mais rápido do que daria tempo para executar
		a função fale - lembrando que a função fale tem um
		time.Sleep.
	*/
	// go fale("Maria", "Ei...", 500)
	// go fale("João", "Opa...", 500)
	/*
		Para forçar a execução da função, por hora, utilizaremos
		o time.Sleep no método main também, dando o tempo necessário
		para a função fale executar.
	*/
	// time.Sleep(time.Second * 5)

	go fale("Maria", "Entendi!!", 10)
	fale("João", "Parabéns!", 5)
	/*
		Lembrando que devido a concorrência, a ordem das falas das
		linhas (45-46) serão sempre aleatórias a cada execução.
	*/

	f.Println("Fim!")
}
