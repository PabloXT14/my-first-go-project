package pointers

import "fmt"

func Execute() {
	intPointer()
}

func intPointer() {
	age := 25
	pointer := &age

	fmt.Println("Value of age:", age)
	fmt.Println("Address of age:", pointer)
	fmt.Println("Value at address:", *pointer) // DEREFERENCING the pointer to get the value stored at that address
}

// Using pointers to change original value in memory
func ChangeAge(age *int) {
	*age = 30
}

func CreatePointer() *int {
	x := 42
	return &x // Returning the address of x
}
