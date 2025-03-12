package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, name string) {
	ticker := time.NewTicker(250 * time.Millisecond).C
	for {
		select {
		case <-ticker:
			fmt.Println(name, "tick")
		case <-ctx.Done():
			fmt.Println("Finishing due to:", ctx.Err().Error())
			fmt.Println(name, "Received a signal to finish... exiting...")
			return
		}
	}

}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go worker(ctx, "worker 1")
	go worker(ctx, "worker 2")

	time.Sleep(12 * time.Second)

	fmt.Println("main exiting")

}
