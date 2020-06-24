package internal

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// REPL executes a Read-eval-print loop until the program is closed.
func REPL() error {
	fmt.Printf("Slip %s\n", Version)

	env := NewEnviroment()
	reader := bufio.NewReader(os.Stdin)

	for lineNo := 0; ; lineNo++ {
		fmt.Printf("slip:%d:> ", lineNo)

		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println()
				return nil
			}
			return fmt.Errorf("failed to read line: %v", err)
		}

		values, err := Parse(line)
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, value := range values {
			res := value.Eval(env)
			if res == nil {
				fmt.Println("nil")
			} else {
				fmt.Println(res)
			}
		}
	}
}
