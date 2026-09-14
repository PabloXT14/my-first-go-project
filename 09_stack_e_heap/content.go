package stack_e_heap

import "fmt"

func Execute() {
	heapPointer := heap()

	fmt.Println("Heap Example:")
	fmt.Println("Value of number:", *heapPointer)
	fmt.Println("Address of number:", heapPointer)
}

func stack() {
	a := 10
	b := 20

	fmt.Println("Stack Example:")
	fmt.Println("Value of a:", a)
	fmt.Println("Value of b:", b)
	fmt.Println("Sum of a and b:", a+b)
}

// * = Pointer an address in memory of the pc
func heap() *int {
	number := 10

	// & = Get the address of the variable in memory
	return &number
}
