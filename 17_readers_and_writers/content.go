package readers_and_writers

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func Execute() {
	readDataFromBuffer()
	// exampleWriteOnConsole()
	// exampleReadAndWriteTogether()
	// exampleReadAll()
	// exampleMultiWriter()
}

func readDataFromBuffer() {
	reader := strings.NewReader("Hello, Go!")

	buffer := make([]byte, 5)

	for {
		n, err := reader.Read(buffer)

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			fmt.Println("Error:", err)
			panic(err)
		}

		fmt.Printf("Lido: %s\n", buffer[:n])
	}
}

func exampleWriteOnConsole() {
	message := []byte("Hello, Go!")

	os.Stdout.Write(message)
}

func exampleReadAndWriteTogether() {
	reader := strings.NewReader("Learning Readers and Writes in Go!")

	io.Copy(os.Stdout, reader)
}

func exampleReadAll() {
	reader := strings.NewReader("Hello, Go!")

	data, err := io.ReadAll(reader)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Lido:", string(data))
}

func exampleMultiWriter() {
	var buffer bytes.Buffer

	writer := io.MultiWriter(os.Stdout, &buffer)

	writer.Write([]byte("Hello, Go!"))

	fmt.Println("\nBuffer content:", buffer.String())
}
