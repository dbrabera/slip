package main

import (
	"fmt"
	"io"
	"os"

	"github.com/peterh/liner"
)

func main() {
	os.Exit(repl())
}

func repl() int {
	liner := liner.NewLiner()
	defer liner.Close()

	reader := NewReader()

	for {
		line, err := liner.Prompt("slip> ")

		if err == io.EOF {
			fmt.Println()
			return 0
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading line: %s\n", err)
			return 1
		}

		reader.Init(line)
		for obj := reader.Read(); obj != nil; obj = reader.Read() {
			fmt.Println(obj.Eval())
		}
	}

	return 0
}
