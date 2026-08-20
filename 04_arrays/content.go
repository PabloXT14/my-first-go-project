package arrays

import "fmt"

func Execute() {
	// ARRAYS - BASIC DECLARATION
	/*
		arr1 := [3]int{} // Array of integers with a fixed size of 3
		arr1[0] = 10
		arr1[1] = 20
		arr1[2] = 30

		fmt.Println("Array 1:", arr1)
	*/

	// SHORT STATEMENT
	/*
		arr2 := [3]int{10, 20, 30} // Array of integers with a fixed size of 3
		fmt.Println("Array 2:", arr2)
	*/

	// START SPECIFIC INDEX
	/*
		arr3 := [5]int{0: 10, 3: 40}
		fmt.Println("Array 3:", arr3)
	*/

	// SIZE FIXED (IN COMPILATION TIME)
	/*
		const size = 3
		arr4 := [size]int{10, 20, 30}
		fmt.Println("Array 4:", arr4)
	*/

	// ACCESSING ELEMENTS
	/*
		arr5 := [3]int{30, 20, 10}
		fmt.Println("First element:", arr5[0])
		fmt.Println("Second element:", arr5[1])
		fmt.Println("Third element:", arr5[2])
	*/

	// SLICES - HAVE DYNAMIC SIZE
	slice1 := []int{10, 20, 30}
	fmt.Println("Slice 1:", slice1)

	// APPEND
	slice2 := []int{10, 20, 30}
	slice2 = append(slice2, 40)
	fmt.Println("Slice 2:", slice2)
}
