package main

import (
	"fmt"
	pointers "my-first-go-project/10_pointers"
)

func main() {
	age := 25
	fmt.Println("Original age:", age)

	pointers.ChangeAge(&age)
	fmt.Println("New age:", age)

	pointer := pointers.CreatePointer()
	fmt.Println("Value at pointer:", *pointer)

	// Avoid nil pointer dereference by checking if the pointer is nil
	var nilPointer *int

	if nilPointer != nil {
		fmt.Println("Value at nil pointer:", *nilPointer)
	} else {
		fmt.Println("nilPointer is nil, cannot dereference.")
	}
}
