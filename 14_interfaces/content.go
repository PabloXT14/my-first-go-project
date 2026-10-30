package interfaces

import "fmt"

func Execute() {
	interfaceUtilityExample()
}

type Speaker interface {
	Speak()
}

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Woof!")
}

type Cat struct{}

func (c Cat) Speak() {
	fmt.Println("Meow!")
}

func MakeSound(s Speaker) {
	s.Speak()
}

func simpleExample() {
	dog := Dog{}
	cat := Cat{}

	MakeSound(dog)
	MakeSound(cat)
}

// Empty interface (here we can store any type of value)
func emptyInterfaceExample() {
	var value interface{}

	value = 42
	fmt.Println(value)

	value = "hello"
	fmt.Println(value)

	value = true
	fmt.Println(value)
}

func simulatePanicWithNilInterface() {
	var speaker Speaker

	fmt.Println(speaker == nil)

	speaker.Speak() // This will cause a panic because speaker is nil
}

// Interface with nil pointer example

type DogWithNilPointer struct{}

func (d *DogWithNilPointer) Speak() {
	if d == nil {
		fmt.Println("Dog is nil")
		return
	}

	fmt.Println("Woof!")
}

func interfaceWithNilPointerExample() {
	var d *DogWithNilPointer

	var speaker Speaker = d

	fmt.Println(speaker == nil) // This will print false because speaker holds a non-nil interface value

	speaker.Speak() // This will not panic, it will print "Dog is nil"
}

// Why interfaces are useful: They allow us to define a set of methods that a type must implement, enabling polymorphism and code reusability. By using interfaces, we can write functions that can operate on different types as long as they satisfy the interface, making our code more flexible and extensible.

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}

func (e EmailNotifier) Send(message string) {
	fmt.Println("Sending email with message:", message)
}

type SMSNotifier struct{}

func (s SMSNotifier) Send(message string) {
	fmt.Println("Sending SMS with message:", message)
}

func Notify(n Notifier) {
	n.Send("Hello! You received a notification.")
}

func interfaceUtilityExample() {
	email := EmailNotifier{}
	sms := SMSNotifier{}

	Notify(email)
	Notify(sms)
}
