# Go

Curso udemy: [Curso go](http://www.udemy.com/curso-go).

# Verificando as informações do GO
`$ go env`
`$ go env GOROOT`
`$ go env GOPATH`

# Verifica se há problemas de compilação
`$ go vet <nome-arquivo.go>`

# Compila e executa
`$ go build <nome-arquivo.go>`
`$ go run <nome-arquivo.go>`

# Instalando pacotes
`$ go get -u <nome-pacote>`

Exemplo:
`$ go get -u github.com/go-sql-driver/mysql`
`$ go get -u github.com/AnitaChiele/go_src`

# Rodando pacotes
## Linux
Acessar o diretório dos pacotes e executar:
`$ go run *.go`

## Windows
Acessar o diretório dos pacotes e executar:
`$ go run arquivo01.go arquivo02.go`
> No caso do windows é necessário especificar arquivo por arquivo.

# Diretório pasta GO
Está acessível a partir da pasta do usuário em uma pasta chamada: go.

## Pastas:
**bin**: tem os arquivos executáveis.
**pkg**: os compilados (separado por arquiteturas).
**src**: arquivos de desenvolvimento.

# Rodando os testes
Acessar o diretório onde está o teste e rodar o comando:
`$ go test`

## Rodando todos os testes do projeto
Acessar o diretório onde está o teste e rodar o comando:
`$ go test ./...`

## Rodando os testes no modo verboso
`$ go test -v`

## Rodando o coverage
Só funciona dentro do diretório da pasta GO.
`$ go test -cover`

### Para gerar um arquivo com o resultado
`$ go test -coverprofile=<nome-arquivo>.out`
Ex:
`$ go test -coverprofile=resultado.out`

#### Para ler o arquivo do resultado do teste que foi gerado
`$ go tool cover -func=<nome-arquivo-gerado>.out`
Ex:
`$ go tool cover -func=resultado.out`

#### Para gerar um arquivo .html com o resultado do teste
`$ go tool cover -html=<nome-arquivo-gerado>.out`
Ex:
`$ go tool cover -html=resultado.out`

# Servidor HTTP
Para executar o servidor http corretamente é necessário que seja feito pelo terminal.
Vá até o diretório que tem o arquivo e execute:
`$ go run <nome-arquivo>.go`

