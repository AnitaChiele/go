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
