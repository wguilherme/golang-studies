package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	withoutGoRoutine()

	withGoRoutine()

}

func withGoRoutine() {
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(10)
	for range 10 {
		go func() {
			defer wg.Done()
			res, err := http.Get("https://google.com")
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()
			fmt.Println(res.Status)
		}()
	}
	wg.Wait()
	fmt.Println("Time taken (gorountine):", time.Since(start))
}

func withoutGoRoutine() {
	start := time.Now()

	for range 10 {
		res, err := http.Get("https://google.com")
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
		fmt.Println(res.Status)
	}
	fmt.Println("Time taken:", time.Since(start))
}
