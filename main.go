package main

import "fmt"

// SIMPLE FUNCTION
func sum(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

// MULTIPLE RETURN VALUES
func divide(a int, b int) (int, int) {
	return a / b, a % b
}

// VARIADIC FUNCTION
func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

// HIGHER-ORDER FUNCTION
func applyOperation(a int, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// CLOSURE
func counter() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

// NAKED RETURN
func calculate(a int, b int) (sum int, product int) {
	sum = a + b
	product = a * b

	return
}

func main() {
	sum, product := calculate(5, 10)

	fmt.Println("Sum:", sum)
	fmt.Println("Product:", product)
}
