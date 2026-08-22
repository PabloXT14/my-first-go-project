package switches

func Execute() {
	// BASIC SWITCH (DON'T NEED BREAKS)
	/*
		day := 8

		switch day {
		case 1:
			println("Monday")
		case 2:
			println("Tuesday")
		case 3:
			println("Wednesday")
		case 4:
			println("Thursday")
		case 5:
			println("Friday")
		case 6:
			println("Saturday")
		case 7:
			println("Sunday")
		default:
			println("Invalid day")
		}
	*/

	// DON'T NEED TO BE INTEGERS
	/*
		command := "start"
		switch command {
		case "start":
			println("Starting the process...")
		case "stop":
			println("Stopping the process...")
		case "pause":
			println("Pausing the process...")
		default:
			println("Unknown command")
		}
	*/

	// ACCEPT MULTIPLE VALUES IN THE SAME CASE
	/*
		fruit := "banana"

		switch fruit {
		case "apple", "banana", "cherry":
			println("This is a common fruit.")
		case "kiwi", "mango":
			println("This is an exotic fruit.")
		default:
			println("Unknown fruit.")
		}
	*/

	// SWITCH WIT INITIALIZATION
	/*
		switch length := len("Hello, World!"); length {
		case 0:
			println("Empty string")
		case 1, 2, 3, 4, 5:
			println("Short string")
		default:
			println("Long string")
		}
	*/

	// SWITCH WITOUT EXPRESSION (WORKS LIKE IF-ELSE)
	/*
		number := 15

		switch {
		case number < 0:
			println("Negative number")
		case number == 0:
			println("Zero")
		case number > 0 && number <= 10:
			println("Positive number between 1 and 10")
		default:
			println("Positive number greater than 10")
		}
	*/

	// FALLTHROUGH (EXECUTES THE NEXT CASE)
	number := 2

	switch number {
	case 1:
		println("One")
	case 2:
		println("Two")
		fallthrough
	case 3:
		println("Three")
	default:
		println("Other number")
	}

}
