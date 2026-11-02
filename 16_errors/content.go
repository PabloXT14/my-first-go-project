package errors

import (
	"errors"
	"fmt"
)

func Execute() {
	// result := withdraw(100, 200)
	// if result != nil {
	// 	fmt.Println("Error:", result)
	// }

	// result = withdraw(200, 100)
	// if result != nil {
	// 	fmt.Println("Error:", result)
	// }

	// errorIsExample()

	// errorAsExample()

	// result := getUser()
	// if result != nil && errors.Is(result, ErrDatabaseConnection) {
	// 	fmt.Println(result)
	// }

	var err = validateUser("", "")

	fmt.Println("Errors:", err)

	if errors.Is(err, ErrInvalidName) {
		fmt.Println("Ocorreu um erro de nome inválido")
	}

	if errors.Is(err, ErrInvalidEmail) {
		fmt.Println("Ocorreu um erro de email inválido")
	}
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

// ASSIGNING FUNCTIONS AND ERROR RETURNING Ação

func saveUser(name string) error {
	if name == "" {
		return errors.New("o nome não poder estar vazio")
	}

	fmt.Println("Salvando usuário:", name)

	return nil
}

// CREATING CUSTOM ERROR

type InsufficientBalanceError struct {
	Balance float64
}

func (e *InsufficientBalanceError) Error() string {
	return "saldo insuficiente"
}

func withdraw(balance, amount float64) error {
	if amount > balance {
		return &InsufficientBalanceError{
			Balance: balance,
		}
	}

	return nil
}

// ERRORS.IS

var ErrNotFound = errors.New("not found")

func findUser() error {
	return fmt.Errorf("falha ao busca usuário: %w", ErrNotFound)
}

func errorIsExample() {
	err := findUser()

	if errors.Is(err, ErrNotFound) {
		fmt.Println("Usuário não encontrado")
	}

}

// ERRORS.AS

func errorAsExample() {
	err := withdraw(100, 150)

	var balanceErr *InsufficientBalanceError

	if errors.As(err, &balanceErr) {
		fmt.Println("Error: ", balanceErr.Error())
		fmt.Println("Saldo atual: ", balanceErr.Balance)
	}
}

// ERROR WRAPPING

func getUser() error {
	err := Database()

	if err != nil {
		return fmt.Errorf("erro ao buscar usuário: %w", err)
	}

	return nil
}

// Obs: O erro percorre várias camadas da aplicação, mas sua origem continua identificável.

var ErrDatabaseConnection = errors.New("conexão recusada")

func Database() error {
	return fmt.Errorf("erro no banco de dados: %w", ErrDatabaseConnection)
}

func Repository() error {
	err := Database()

	return fmt.Errorf("erro no repository: %w", err)
}

func Service() error {
	err := Repository()

	return fmt.Errorf("erro no service: %w", err)
}

// MULTIPLE ERRORS

var ErrInvalidName = errors.New("nome inválido")
var ErrInvalidEmail = errors.New("email inválido")

func validateUser(name, email string) error {
	var errorsResult error

	if name == "" {
		errorsResult = errors.Join(errorsResult, ErrInvalidName)
	}

	if email == "" {
		errorsResult = errors.Join(errorsResult, ErrInvalidEmail)
	}

	return errorsResult
}
