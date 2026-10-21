package arrays_and_slices

import "fmt"

func Arrays() {
	var numbers [3]int = [3]int{1, 2, 3}

	numbers[1] = 20

	fmt.Println("Array:", numbers)
	fmt.Println("Element at index 1:", numbers[1])

	a := [3]int{10, 20, 30}
	b := a

	b[0] = 99

	fmt.Println("Array a:", a)
	fmt.Println("Array b:", b)
}

func Slices() {
	slice := []int{1, 2, 3}

	slice[1] = 20

	fmt.Println("Slice:", slice)
	fmt.Println("Length of slice:", len(slice))
	fmt.Println("Capacity of slice:", cap(slice))

	fmt.Println("------------------------------------------------------------")

	numbers2 := []int{10, 20, 30, 40, 50}

	slices := numbers2[1:4] // seleciona os elementos do índice 1 até antes do índice 4
	// Regras importantes
	/*
		- numbers[:3] → elementos dos índices `0` a `2`.
		- numbers[2:] → elementos do índice `2` até o final.
		- numbers[:] → todos os elementos.
	*/

	fmt.Println("Slice2:", slices)

	fmt.Println("------------------------------------------------------------")

	numbers3 := [3]int{10, 20, 30}
	slice3 := numbers3[0:3]

	slice3[0] = 99

	fmt.Println("Slice3:", slice3)
	fmt.Println("Numbers3:", numbers3)

	fmt.Println("------------------------------------------------------------")

	original := []int{10, 20, 30}
	clone := make([]int, len(original))

	copy(clone, original)

	clone[0] = 99

	fmt.Println("Original:", original)
	fmt.Println("Clone:", clone)

	fmt.Println("------------------------------------------------------------")

	numbers4 := [5]int{10, 20, 30, 40, 50}
	slice4 := numbers4[1:3]

	fmt.Println("Length of slice4:", len(slice4))
	fmt.Println("Capacity of slice4:", cap(slice4))

	fmt.Println("------------------------------------------------------------")

	var slice5 []int

	fmt.Println(slice5 == nil) // true
	fmt.Println("Length of slice5:", len(slice5))
	fmt.Println("Capacity of slice5:", cap(slice5))

	fmt.Println("------------------------------------------------------------")

	var x []int         // nil
	y := []int{}        // vazio, mas não nil
	z := make([]int, 0) // vazio, mas não nil

	fmt.Println(x == nil) // true
	fmt.Println(y == nil) // false
	fmt.Println(z == nil) // false
}
