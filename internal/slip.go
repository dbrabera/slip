package internal

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/peterh/liner"
	"github.com/pkg/errors"
)

type Slip struct {
	env *Enviroment
}

func NewSlip() *Slip {
	env := NewEnviroment()

	for name, fn := range CoreFuncs {
		env.Define(NewSymbol(name), fn)
	}

	return &Slip{env: env}
}

func (s *Slip) Repl() error {
	liner := liner.NewLiner()
	defer liner.Close()

	for {
		line, err := liner.Prompt("slip> ")
		if err != nil {
			if err == io.EOF {
				fmt.Println()
				return nil
			}
			return errors.Wrap(err, "failed to read line")
		}

		values, err := Parse(line)
		if err != nil {
			return err
		}

		for _, value := range values {
			fmt.Println(value.Eval(s.env))
		}
	}

	return nil
}

func (s *Slip) Run(filename string) error {
	src, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.Wrap(err, "failed to read file")
	}

	values, err := Parse(string(src))
	if err != nil {
		return err
	}

	for _, value := range values {
		value.Eval(s.env)
	}
	return nil
}

func (s *Slip) Exec(src string) (Value, error) {
	values, err := Parse(string(src))
	if err != nil {
		return nil, err
	}

	var out Value
	for _, value := range values {
		out = value.Eval(s.env)
	}
	return out, nil
}
