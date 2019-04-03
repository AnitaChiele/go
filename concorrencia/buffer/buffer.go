package main

import (
	f "fmt"
	"time"
)

func rotina(ch chan int) {
	ch <- 1
	f.Println("Enviou 01 para o channel!")
	ch <- 2
	f.Println("Enviou 02 para o channel!")
	ch <- 3
	f.Println("Enviou 03 para o channel!")
	ch <- 4
	f.Println("Enviou 04 para o channel!")
	ch <- 5
	f.Println("Enviou 05 para o channel!")
	ch <- 6
	f.Println("Enviou 06 para o channel!")
}

func main() {
	// Channel com buffer de 3 posições.
	ch := make(chan int, 3)
	go rotina(ch)

	// deixar de ler dados do channel não gera deadlock porém,
	// deixar de enviar dado no channel e ter alguém esperando
	// é o que gera deadlock.

	/*
		Para elucidar melhor o código, descomente o time.Sleep(time.Second)
		Quando esse trecho está descomentado ele espera um segundo antes de
		exibir.

		Explicando o que acontece:
		Como o channel tem buffer 3, ele só consegue preencher até 3 valores
		int, por isso ele só irá imprimir até a msg:
		"Enviou 03 para o channel!".

		A partir do valor 03 enviado pro channel, os envios ficarão pausados
		até que uma parte do chaneel seja consumida, liberando espaço para
		preencher os outros.
	*/

	time.Sleep(time.Second)
	f.Println(<-ch)

	/*
		O <-ch já faz a leitura do primeiro buffer do channel e já remove o conteúdo,
		do buffer, liberando espaço para a função rotina armazenar um int nesse espaço
		que ficou 'vago'.

		O time.Sleep(time.Scond) da linha baixo é necessário pq o método main terminou
		antes de exibir o valor do buffer do channel não dando tempo pra executar a última
		linha de código comentada (f.Println(<-ch)). A linha do time.Sleep foi adicionada
		apenas para ser visível o que aconteceu.

		Descomente as duas linhas abaixo para visualizar o que foi explicado nos
		comentários das linhas 33-44.
	*/

	// time.Sleep(time.Second)
	// f.Println(<-ch)
}
