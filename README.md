#Go

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