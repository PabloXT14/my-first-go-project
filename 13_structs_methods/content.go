package structs_and_methods

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City  string
	State string
}

type User struct {
	Name  string `json:"name"` // `` this specifies structs tags, which are metadata that can be used by other packages (like encoding/json) to customize the behavior of the struct fields during serialization and deserialization.
	Age   int    `json:"age"`
	Email string `json:"email_address"`
	Address
}

func Execute() {
	structTagsExample()
}

func simpleExample() {
	user := User{
		Name:  "John Doe",
		Age:   30,
		Email: "johndoe@email.com",
	}

	fmt.Println(user)
	fmt.Println(user.Name)
	fmt.Println(user.Age)
	fmt.Println(user.Email)
}

func copyExample() {
	user1 := User{
		Name:  "John Doe",
		Age:   30,
		Email: "johndoe@email.com",
	}

	user2 := user1

	user2.Name = "Anne Smith"

	fmt.Println(user1)
	fmt.Println(user2)
}

// Receiver by value (receiver is a copy of the original struct, so changes made to the receiver do not affect the original struct)
func (user User) Introduce() {
	fmt.Println("Hi, my name is", user.Name)
}

func methodsExample() {
	user := User{
		Name:  "John Doe",
		Age:   30,
		Email: "johndoe@email.com",
	}

	user.Introduce()
}

// Receiver by pointer (receiver is a pointer to the original struct, so changes made to the receiver affect the original struct)
func (user *User) ChangeName(newName string) { // test remove pointer receiver and see the difference
	user.Name = newName
}

func pointerReceiverExample() {
	user := User{
		Name:  "John Doe",
		Age:   30,
		Email: "johndoe@email.com",
	}

	user.ChangeName("Bob Doe")
	fmt.Println(user)
}

func embeddedStructExample() {
	user := User{
		Name:  "John Doe",
		Age:   30,
		Email: "johndoe@email.com",
		Address: Address{
			City:  "New York",
			State: "NY",
		},
	}

	fmt.Println(user)
	fmt.Println(user.Name)
	fmt.Println(user.Address.City)
	fmt.Println(user.Address.State)
}

func structTagsExample() {
	user := User{
		Name:  "John Doe",
		Age:   30,
		Email: "johndoe@email.com",
	}

	data, _ := json.Marshal(user)

	fmt.Println(string(data))
}
