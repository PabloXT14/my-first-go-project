package system_types_and_constants

import (
	"fmt"
	"strconv"
)

// BASIC TYPES
var active bool = true

var age int = 20
var height float64 = 1.75
var number complex128 = 2 + 3i

// INTEGER TYPES
var int8Var int8 = 127                   // int8 range: -128 to 127
var int16Var int16 = 32767               // int16 range: -32768 to 32767
var int32Var int32 = 2147483647          // int32 range: -2147483648 to 2147483647
var int64Var int64 = 9223372036854775807 // int64 range: -9223372036854775808 to 9223372036854775807
var idade uint = 20                      // uint represents integers that aren't negative

// BYTE AND RUNE TYPES
var letter byte = 'A' // byte is an alias for uint8, used to represent ASCII characters
var symbol rune = '🚀' // rune is an alias for int32, used to represent Unicode code points

// CONVERSION BETWEEN TYPES
var a int = 10
var b float64 = 2.5

// CONSTANTS
const pi = 3.14159

// pi = 3.14 // This will cause a compilation error because constants cannot be reassigned

const (
	Name string = "John Doe"
	Age  int    = 30
)

// LITERAL CONSTANTS
const x = 100     // untyped constant
const y int = 200 // typed constant
const z = "Hello, World!"

func Execute() {
	var x int = 10084

	s := string(x) // Will convert the integer to a string representation of its Unicode code point (a rune), not the string "10084"

	s2 := strconv.FormatInt(int64(x), 10) // Correct way to convert int to string

	fmt.Println("Converted int to string:", s)
	fmt.Println("Converted int to string using strconv:", s2)
}
