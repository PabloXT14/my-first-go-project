package maps

import "fmt"

func Execute() {
	deleteKeyWhileIterating()
}

func simpleMap() {
	ages := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 35,
	}

	fmt.Println("Ages:", ages)
	fmt.Println("Alice's age:", ages["Alice"])
}

func startMapEmpty() {
	ages := make(map[string]int)

	ages["Alice"] = 25
	ages["Bob"] = 30

	fmt.Println("Ages:", ages)

	// Atenção: declarar apenas uma variável map não é suficiente para inserir elementos.
	/*
		ages := map[string]int{}

		fmt.Print(ages == nil) // true

		ages["Alice"] = 25 // panic: assignment to entry in nil map
	*/
}

func updateMapValue() {
	ages := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 35,
	}

	fmt.Println("Ages:", ages)

	ages["Alice"] = 26 // Atualizando o valor de Alice
	fmt.Println("Updated Ages:", ages)
}

func checkIfKeyExists() {
	ages := map[string]int{
		"Alice":   0,
		"Bob":     30,
		"Charlie": 35,
	}

	age, exists := ages["Alice"]
	fmt.Println(age)
	fmt.Println(exists)

	age, exists = ages["David"]
	fmt.Println(age)
	fmt.Println(exists)
}

func deleteKeyFromMap() {
	ages := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 35,
	}

	fmt.Println("Ages:", ages)

	delete(ages, "Bob") // Deletando o elemento com a chave "Bob" (obs: se a chave não existir, ou se o map estiver vazio, não ocorre erro)
	fmt.Println("Updated Ages:", ages)
}

func clearMap() {
	ages := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 35,
	}

	fmt.Println("Ages:", ages)

	clear(ages) // Limpando o map, mas mantém a mesma capacidade de armazenamento

	fmt.Println("Cleared Ages:", ages)
}

func iterateOverMap() {
	ages := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 35,
	}

	// obs: a ordem de iteração de um map em Go não é garantida. Não conte com uma ordem alfabética ou com a ordem em que os elementos foram adicionados.
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}
}

func deleteKeyWhileIterating() {
	ages := map[string]int{
		"Alice":   17,
		"Bob":     30,
		"Charlie": 35,
	}

	for name, age := range ages {
		if age < 18 {
			delete(ages, name)
		}
	}

	fmt.Println("Updated Ages:", ages)
}
