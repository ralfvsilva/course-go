package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

// Google I/O 0012 - Go Concurrency Patterns

// <-chan - canal somente de leitura (apenas recebe dados)

func Titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := io.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)</title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}

func main() {
	t1 := Titulo("https://www.cod3r.com.br", "https://www.google.com.br")

	fmt.Println("Primeiro:", <-t1, "\nSegundo:", <-t1)
}
