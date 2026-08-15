package pacote

import (
	"fmt"
	"my-first-go-project/pacote/internal/foo"
)

var Message2 string = "Hello from pacote/b.go"

func PrintMessage() {
	fmt.Println(foo.LocalMessage)
}
