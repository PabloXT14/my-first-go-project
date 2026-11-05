package goroutines_and_contexts

import (
	"fmt"
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

	concurrentExample()
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
