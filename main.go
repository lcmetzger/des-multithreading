package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	link1 := make(chan string)
	link2 := make(chan string)
	cep := "89010-904"

	// APICep
	go func() {
		res, err := http.Get("https://cdn.apicep.com/file/apicep/" + cep + ".json")
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		link1 <- string(body)

	}()

	// ViaCep
	go func() {
		res, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		link2 <- string(body)
	}()

	select {
	case res1 := <-link1:
		println("Resposta de APICEP:\n" + res1)

	case res2 := <-link2:
		println("Resposta de VIACEP:\n" + res2)

	case <-time.After(time.Second):
		println("ocorreu timeout de 1 segundo")
	}

}
