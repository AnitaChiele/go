# Concorrência VS Paralelismo
Curso udemy: [Curso go](http://www.udemy.com/curso-go).
Aulas: S09A66-68

> Go é uma linguagem concorrrente.

## Paralelismo:
Utiliza mais de um processador (ou núcleos) o que possibilita que sejam executados mais de uma coisa ao mesmo tempo.
Processamento simultâneos, executa códigos simultaneamente.

## Concorrência:
Capacidade de administrar múltiplas tarefas e não depende da quantidade de processadores ou núcleos de processadores.
Na concorrência ocorre a divisão de tempo de utilização do processor por tarefa, ou seja, cada tarefa tem um tempo para executar e passar para a próxima, passando a sensação de se tratar de um paralelismo. Na prática, o processador executa um pouco cada tarefa por vez (escalonadores).

## Paralelismo vs Concorrência:
O paralelismo é mais custoso devido ao overhead de ter que juntar todas as informações que foram "distribuídas". Dependendo da situação, paralelizar não é a melhor solução e pode ser até mais lento para obter uma resposta devido esse overhead.

Paralelismo vem como consequência de planejar bem o sistema concorrente.