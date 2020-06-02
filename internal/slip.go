package internal

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/peterh/liner"
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

func (self *Slip) Repl() int {
	liner := liner.NewLiner()
	defer liner.Close()

	for {
		line, err := liner.Prompt("slip> ")

		if err == io.EOF {
			fmt.Println()
			return 1
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading line: %s\n", err)
			return 1
		}

		self.reader.Init(line)
		for obj := self.reader.Read(); obj != nil; obj = self.reader.Read() {
			fmt.Println(obj.Eval(self.env))
		}
	}

	return 0
}

func (self *Slip) Run(filename string) int {
	src, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading script: %s\n", err)
		return 1
	}

	self.reader.Init(string(src))
	for obj := self.reader.Read(); obj != nil; obj = self.reader.Read() {
		obj.Eval(self.env)
	}
	return 0
}

func (self *Slip) Exec(src string) Object {
	self.reader.Init(src)
	obj := self.reader.Read()
	return obj.Eval(self.env)
}