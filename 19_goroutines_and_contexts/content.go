package goroutines_and_contexts

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func Execute() {
	// EXAMPLE WITHOUT GOROUTINES
	// printMessage("Hello, world!")
	// printMessage("Hi, world!")
	// printMessage("Hello, universe!")

	// EXAMPLE WITH GOROUTINES
	/*
		Nesse segundo exemplo, o `main()` pode terminar antes das Goroutines executarem. Quando a função `main()` termina, o programa inteiro é encerrado, mesmo que outras Goroutines ainda estejam trabalhando.

		Para resolver isso, podemos utilizar `WaitGroup`.
	*/
	// go printMessage("Hello, world!")
	// go printMessage("Hi, world!")
	// go printMessage("Hello, universe!")

	// EXAMPLE WITH GOROUTINES AND WAITGROUP
	// var wg sync.WaitGroup

	// wg.Add(3) // Adiciona 3 tasks ao contador do WaitGroup

	// go printMessageWithWaitGroup("Task 1", &wg) // executa a função em uma goroutine
	// go printMessageWithWaitGroup("Task 2", &wg)
	// go printMessageWithWaitGroup("Task 3", &wg)

	// wg.Wait() // Aguarda todas as goroutines terminarem

	// fmt.Println("All tasks completed.")

	// concurrentExample()

	// CONTEXT EXAMPLE
	// contextExample()

	// SIMULATE A TIMEOUT ON REQUEST
	// simulateATimeoutOnRequest()

	// CONTEXT WITH GOROUTINES EXAMPLE
	contextWithGoroutinesExample()
}

func printMessage(message string) {
	fmt.Println(message)
}

// WAITGROUP EXAMPLE

func printMessageWithWaitGroup(message string, wg *sync.WaitGroup) {
	defer wg.Done() // Decrementa o contador do WaitGroup quando a função terminar

	fmt.Println(message)
}

// TRANSFORM SEQUENTIAL CODE INTO CONCURRENT

func fetchData(service string) {
	time.Sleep(2 * time.Second)

	fmt.Println("Data received:", service)
}

func sequentialExample() {
	start := time.Now()

	fetchData("API 1")
	fetchData("API 2")
	fetchData("API 3")

	fmt.Println("Time taken (sequential):", time.Since(start))
}

func concurrentExample() {
	var wg sync.WaitGroup

	start := time.Now()

	services := []string{"API 1", "API 2", "API 3"}

	for _, service := range services {
		wg.Add(1)

		go func(service string) {
			defer wg.Done()

			fetchData(service)
		}(service)
	}

	wg.Wait()

	fmt.Println("Time taken (concurrent):", time.Since(start))
}

// CONTEXT

func contextExample() {
	/*
		Context Definition:

		Eles ajudam a controlar a execução de operações concorrentes, principalmente quando precisamos trabalhar com timeouts, cancelamentos e comunicação entre Goroutines.

		O Context é muito utilizado em aplicações que realizam requisições HTTP, consultas ao banco de dados e outras operações que podem demorar para terminar.
	*/

	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)

	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"https://jsonplaceholder.typicode.com/posts",
		nil,
	)

	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	response, err := http.DefaultClient.Do(req)

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("Request timed out.")
			return
		}

		fmt.Println("Error making request:", err)
		return
	}

	defer response.Body.Close()

	fmt.Println("Status:", response.Status)
	fmt.Println("Status code:", response.StatusCode)
}

func simulateATimeoutOnRequest() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)

	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"https://httpbin.org/delay/5",
		nil,
	)

	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	_, err = http.DefaultClient.Do(req)

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("Request timed out.")
			return
		}

		fmt.Println("Error making request:", err)
		return
	}

	fmt.Println("Request completed successfully.")
}

// CONTEXT WITH GOROUTINES

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker cancelled.")
			return
		default:
			fmt.Println("Worker is working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func contextWithGoroutinesExample() {
	ctx, cancel := context.WithCancel(
		context.Background(),
	)

	go worker(ctx)

	time.Sleep(2 * time.Second)

	cancel() // Cancela o contexto, sinalizando para a goroutine parar

	time.Sleep(500 * time.Millisecond) // Aguarda um pouco para que a goroutine termine
}
