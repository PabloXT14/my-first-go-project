package generics

import "fmt"

func Execute() {
	// WITHOUT GENERICS
	// sumIntResult := sumInt(1, 2)
	// println("Sum of integers:", sumIntResult)

	// sumFloatResult := sumFloat(1.5, 2.5)
	// println("Sum of floats:", sumFloatResult)

	// WITH GENERICS
	// sumIntResult := sum(1, 2)
	// println("Sum of integers:", sumIntResult)

	// sumFloatResult := sum(1.5, 2.5)
	// println("Sum of floats:", sumFloatResult)

	// sumResult := sum(1, 2.5)
	// println("Sum:", sumResult)

	// COMPARABLE CONSTRAINTS
	// fmt.Println(equal(10, 10))
	// fmt.Println(equal("Go", "Go"))
	// fmt.Println(equal(true, false))

	// GENERIC STRUCTS
	// intBox := Box[int]{Value: 100}
	// stringBox := Box[string]{Value: "Hello, Go!"}

	// fmt.Println("Integer Box Value:", intBox.Value)
	// fmt.Println("String Box Value:", stringBox.Value)

	// CONTAINS FUNCTION WITH GENERICS
	numbers := []int{1, 2, 3, 4, 5}
	names := []string{"Alice", "Bob", "Charlie"}

	fmt.Println(Contains(numbers, 3))
	fmt.Println(Contains(names, "Bob"))
	fmt.Println(Contains(names, "Eve"))
}

// WITHOUT GENERICS (we need to create a function for each type we want to use)

func sumInt(a, b int) int {
	return a + b
}

func sumFloat(a, b float64) float64 {
	return a + b
}

// WITH GENERICS (we can create a single function that works with any type)

type Number interface {
	int | float64
}

// func sum[T int | float64](a, b T) T {
// 	return a + b
// }

// OU

func sum[T Number](a, b T) T {
	return a + b
}

// COMPARABLE CONSTRAINTS (we can create a function that works with any type that implements the comparable interface, which includes all types that can be compared using the == and != operators)

func equal[T comparable](a, b T) bool {
	return a == b
}

// GENERIC STRUCTS (we can create a struct that works with any type)

type Box[T any] struct {
	Value T
}

// Get method for the Box struct that returns the value of the box

func (b Box[T]) Get() T {
	return b.Value
}

// CONTAINS FUNCTION WITH GENERICS

func Contains[T comparable](items []T, target T) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}

	return false
}
