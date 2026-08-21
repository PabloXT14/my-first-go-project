package loops

import (
	"fmt"
	"sync"
)

func Execute() {
	// BASIC FOR LOOP
	/*
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
	*/

	// FOR WITH WHILE
	/*
		count := 0

		for count < 5 {
			fmt.Println(count)
			count++
		}
	*/

	// INFINITE LOOP
	/*
		for {
			fmt.Println("This is an infinite loop")
		}
	*/

	// FOR RANGE LOOP
	/*
		names := []string{"Alice", "Bob", "Charlie"}

		// we can ignore the index by using an underscore (_) if we don't need it
		for index, name := range names {
			fmt.Printf("Index: %d, Name: %s\n", index, name)
		}
	*/

	// ITERATE OVER A INTEGER (obs: accept only one variable of iteration)
	/*
		for i := range 10 {
			fmt.Println("Item ", i)
		}
	*/

	// GOROUTINES (obs: this works properly only on versions 1.22+ of Go)
	const numGoroutines = 3

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}

	wg.Wait()
}
