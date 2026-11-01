package errors

import (
	"errors"
	"fmt"
)

func Execute() {
	// handlingError()
	err := saveUser("Pablo")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Usuário salvo com sucesso!")
}

// Default error interface in go
// type error interface {
// 	Error() string
// }

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		// panic("Não é possível dividir por zero") // This will cause the program to crash and print a stack trace, just use in exception scenarios

		return 0, errors.New("não é possível dividir por zero")
	}

	return a / b, nil
}

func handlingError() {
	// result, err := Divide(10, 2)
	result, err := Divide(10, 0)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}

// ASSIGNING FUNCTIONS AND ERROR RETURNING

func saveUser(name string) error {
	if name == "" {
		return errors.New("o nome não poder estar vazio")
	}

	fmt.Println("Salvando usuário:", name)

	return nil
}
