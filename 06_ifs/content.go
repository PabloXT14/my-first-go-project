package ifs

import (
	"errors"
	"fmt"
)

func Execute() {
	// BASIC IF
	/*
		age := 13

		if age >= 18 {
			fmt.Println("You are an adult.")
		} else if age >= 13 {
			fmt.Println("You are a teenager.")
		} else {
			fmt.Println("You are a minor.")
		}
	*/

	// IF WITH SHORT STATEMENT
	/*
		if age := 20; age >= 18 {
			fmt.Println("You are an adult.")
		}
	*/

	// SHORT STATEMENT WITH ERROR
	result, err := 10/2, errors.New("Division by zero") // Simulating a division operation that could produce an error

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}
