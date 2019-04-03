# go

Habilitado a documentação offline (segura o terminal):
`$ godoc -htpp=:<PORTA>`

Exemplo:
`$ godoc -htpp=:1313`

Verificar infomrações do go:
`$ go env`
`$ go env GOROOT`
`$ go env GOPATH`

Verifica se há problemas de compilação:
`$ go vet <nome-arquivo.go>`

Compila e executa:
`$ go build <nome-arquivo.go>`
`$ go run <nome-arquivo.go>`

Instalando pacotes:
`$ go get -u <nome-pacote>`
Exemplo:
`$ go get -u github.com/go-sql-driver/mysql`

Rodando pacotes no linux:
Acessar o diretório dos pacotes e executar:
`$ go run *.go`

Rodando pacotes no windows:
Acessar o diretório dos pacotes e executar:
`$ go run arquivo01.go arquivo02.go`
>> No caso do windows é necessário especificar arquivo por arquivo.

Diretório GO:
Está acessível a partir da pasta do usuário em uma pasta chamada: go.
Pastas:
bin: tem os arquivos executáveis.
pkg: os compilados (separado por arquiteturas)
src: arquivos de desenvolvimento.

O restante do curso será utilizando outro repo onde terá versionado dentro
de src github.com