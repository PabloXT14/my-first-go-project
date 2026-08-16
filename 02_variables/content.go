package variables

import "fmt"

// SIMPLE VARIABLE
var name string = "John Doe"
var age int

// start values
/*
string -> ""
int -> 0
float -> 0.0
bool -> false
*/

// TYPE INFERENCING
var name2 = "John Doe"
var age2 = 20

// SHORT DECLARATION (:= just work inside functions)
func greet() {
	name3 := "John Doe"
	age3 := 20

	fmt.Println(name3)
	fmt.Println(age3)
}

// MULTIPLE VARIABLES
var city, country string = "New York", "USA"

var state, zipCode = "California", 90001
