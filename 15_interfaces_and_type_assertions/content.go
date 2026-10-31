package interfaces_and_type_assertions

import (
	"fmt"
)

func Execute() {
	// emptyInterface()
	// typeAssertion()
	// typeSwitch()
	interfaceStringer()
}

// IMPLICIT INTERFACE IMPLEMENTATION

type Speaker interface {
	Speak()
}

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Woof!")
}

func implicitInterfaceImplementation() {
	var d Speaker = Dog{}

	d.Speak()
}

// EMPTY INTERFACE (interface{} or any)

func ShowValue(value any) {
	fmt.Println(value)
}

func emptyInterface() {
	ShowValue(42)
	ShowValue("hello")
	ShowValue(true)
}

// TYPE ASSERTION

func typeAssertion() {
	var value any = "Hello, World!"
	// var value any = 42

	text, ok := value.(string)

	if ok {
		fmt.Println("É uma string: ", text)
	} else {
		fmt.Println("Não é uma string.")
	}
}

// TYPE SWITCH

func checkType(value any) {
	switch v := value.(type) {
	case int:
		fmt.Println("É um int:", v)
	case string:
		fmt.Println("É uma string:", v)
	case bool:
		fmt.Println("É um booleano:", v)
	default:
		fmt.Println("Tipo desconhecido")
	}
}

func typeSwitch() {
	checkType(42)
	checkType("hello")
	checkType(true)
	checkType(3.14)
}

// INTERFACE STRINGER (the fmt package has a Stringer interface that defines the String() method, which is used to format values as strings)

type Person struct {
	Name string
	Age  int
}

// Try commenting this method and see what happens when you print a Person struct
func (p Person) String() string {
	return "Name: " + p.Name + ", Age: " + fmt.Sprint(p.Age)
}

func interfaceStringer() {
	p := Person{Name: "John Doe", Age: 30}
	fmt.Println(p)
}
