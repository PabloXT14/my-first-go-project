package deferContent

func Execute() {
	// BASIC DEFER (POSTPONED THE EXECUTION OF A FUNCTION UNTIL THE SURROUNDING FUNCTION RETURNS)
	/*
		defer println("Fim")

		println("Inicio")
	*/

	// ARGUMENTS ARE EVALUATED IMMEDIATELY, BUT THE FUNCTION CALL IS POSTPONED
	/*
		number := 10

		defer println(number)

		number = 20

		println(number)
	*/

	// DEFER STACK (LIFO)
	/*
		defer println("First")

		defer println("Second")

		defer println("Third")
	*/

	// RELEASE RESOURCES (CLOSE FILES, DATABASE CONNECTIONS, ETC.)
	/*
		file, err := os.Open("data.txt")

		if err != nil {
			println("Error opening file:", err.Error())
			return
		}

		defer file.Close()

		// Read file...
	*/

	// DEFER SCOPES (DEFER EXECUTES IN THE SCOPE OF THE FUNCTION WHERE IT WAS DECLARED)
	defer println("End of Execute()")

	println("Start of Execute()")

}
