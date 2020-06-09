package internal

import (
	"fmt"
	"io"

	"github.com/peterh/liner"
	"github.com/pkg/errors"
)

// REPL executes a Read-eval-print loop until the program is closed.
func REPL() error {
	env := NewEnviroment()

	liner := liner.NewLiner()
	defer liner.Close()

	fmt.Printf("Slip %s\n", Version)

	for ln := 1; true; ln++ {
		line, err := liner.Prompt(fmt.Sprintf("slip:%d:> ", ln))
		if err != nil {
			if err == io.EOF {
				fmt.Println()
				return nil
			}
			return errors.Wrap(err, "failed to read line")
		}

		values, err := Parse(line)
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, value := range values {
			// use core Println function to correctly print all values
			Println(value.Eval(env))
		}
	}

	return nil
}
