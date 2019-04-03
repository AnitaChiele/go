package main

import (
	f "fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

/*
	Google I/O 2012 - Go Concurrency Patterns:
	https://www.youtube.com/watch?v=f6kdp27TYZs
*/

/*
	Essa função é do tipo generator porque retorna um channel somente-leitura.
	Ou seja, a função que irá chamar essa função não precisará se preocupar em
	criar o channel nem go routines.
*/
func titulo(urls ...string) <-chan string {
	c := make(chan string)

	for _, url := range urls {
		// função anônima
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title[^>]*>([^<]+)</title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return c
}

func main() {
	t1 := titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := titulo("https://www.amazon.com", "https://www.youtube.com")

	f.Println("Primeiros: ", <-t1, "|", <-t2)
	f.Println("Segundos: ", <-t1, "|", <-t2)
}
