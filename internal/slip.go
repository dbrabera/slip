package internal

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/peterh/liner"
	"github.com/pkg/errors"
)

type Slip struct {
	env    *Enviroment
	reader *Reader
}

func NewSlip() *Slip {
	env := NewEnviroment()

	for name, fn := range CoreFuncs {
		env.Define(NewSymbol(name), NewPrimFunc(fn))
	}

	return &Slip{env: env, reader: NewReader()}
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

		s.reader.Init(line)
		for obj := s.reader.Read(); obj != nil; obj = s.reader.Read() {
			fmt.Println(obj.Eval(s.env))
		}
	}

	return nil
}

func (s *Slip) Run(filename string) error {
	src, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.Wrap(err, "failed to read file")
	}

	s.reader.Init(string(src))
	for obj := s.reader.Read(); obj != nil; obj = s.reader.Read() {
		obj.Eval(s.env)
	}
	return nil
}

func (s *Slip) Exec(src string) Object {
	s.reader.Init(src)
	obj := s.reader.Read()
	return obj.Eval(s.env)
}
