package main


import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func worker(id int, dataChannel <-chan int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case data, ok := <-dataChannel:
			if !ok {
				fmt.Printf("Worker %d received close signal, exiting\n", id)
				return
			}
			fmt.Printf("Worker %d received data: %d\n", id, data)
		case <-ctx.Done():
			fmt.Printf("Worker %d received cancel signal, exiting\n", id)
			return
		}
	}
}

func Run(numWorkers int) {
	dataChannel := make(chan int)

	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChannel, ctx, &wg)
	}

	
	go func() {
		for i := 0; i < 10; i++ {
			dataChannel <- i
		}
		close(dataChannel)
	}()

	
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-signalCh
	cancel() 

	
	wg.Wait()
}
