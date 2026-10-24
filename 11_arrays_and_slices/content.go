package arrays_and_slices

import "fmt"

func Slices() {
	// Full slice expression
	/*
		arr := []int{1, 2, 3, 4, 5}

		slice := arr[:2:2] // full slice expression [start:end:capacity] - slice of arr from index 0 to 2 with capacity 2

		fmt.Println("Array:", arr)
		fmt.Println("Slice:", slice)
		fmt.Println("Slice capacity:", cap(slice))
	*/

	// Boundary check = go make a check to ensure that the index is within the bounds of the slice. If you try to access an index that is out of range, Go will panic and terminate the program.
	/*
		slice := []int{1, 2, 3, 4}
		fmt.Println("Slice:", slice)
		fmt.Println("Slice element:", slice[5]) // This will cause a runtime panic: index out of range
	*/

	// slice := []int{1, 2, 3, 4}

	// BoundaryCheck(slice)

	// changeFirst(slice) // This will not change the original slice, because the function receives a copy of the slice, not the original slice.
	slice := [4]int{1, 2, 3, 4} // This is an array, not a slice. Arrays are passed by value, so the function will receive a copy of the array, not the original array.
	changeFirst(slice)

	fmt.Println("Slice1:", slice)

	slice2 := []int{1, 2, 3, 4}
	changeFirstSlice(slice2) // This will change the original slice, because the function receives a reference/pointer to the slice, not a copy of the slice.

	fmt.Println("Slice2:", slice2)
}

func BoundaryCheck(slice []int) {
	_ = slice[3] // Accessing the fourth element (index 3) to demonstrate boundary check, if this line result in an panic error, the next lines will not be executed. Resulting in a good performance, because the program will terminate immediately without executing the next lines of code.

	// OBS: run this command [go run -gcflags="-d=ssa/check_bce" .] on the terminal to see the boundary check optimization in action. The compiler will optimize away the boundary check for the slice access, resulting in better performance.

	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
	fmt.Println(slice[3])
}

func changeFirst(slice [4]int) {
	slice[0] = 10
}

func changeFirstSlice(slice []int) {
	slice[0] = 10
}
