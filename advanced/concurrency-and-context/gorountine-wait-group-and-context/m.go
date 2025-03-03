package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"time"
)

func main() {

	start := time.Now()
	var wg sync.WaitGroup
	var n int = 10
	wg.Add(n)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// dessa forma, o servidor vai demorar 10 segundos para responder
			// e vai fazer com que o contexto expire, cancelando a requisição
			// e retornando um erro de request timed out
			// para testar o sucesso, basta comentar a linha abaixo ou  diminuir o tempo de sleep
			// para um valor menor que o tempo de timeout do contexto (5 segundos, nesse caso)
			time.Sleep(10 * time.Second)
			fmt.Fprintln(w, "Hello, client")
		}),
	)

	for range n {
		go func() {

			defer wg.Done()

			req, err := http.NewRequestWithContext(
				ctx,
				"GET",
				server.URL,
				nil,
			)
			if err != nil {
				panic(err)
			}

			resp, err := http.DefaultClient.Do(req)

			if err != nil {
				if errors.Is(err, context.DeadlineExceeded) {
					fmt.Println("Request timed out")
					return
				}
				panic(err)
			}

			defer resp.Body.Close()
		}()
	}
	wg.Wait()
	fmt.Println("Time taken (gorountine):", time.Since(start))

}
