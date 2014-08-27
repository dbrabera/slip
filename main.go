package main

import (
	"fmt"
	"io"
	"os"

	"github.com/peterh/liner"
)

type Enviroment struct {
	symbols map[string]Object
	parent  *Enviroment
}

func NewEnviroment() *Enviroment {
	return &Enviroment{symbols: make(map[string]Object)}
}

func NewChildEnviroment(parent *Enviroment) *Enviroment {
	return &Enviroment{symbols: make(map[string]Object), parent: parent}
}

func (self *Enviroment) Define(sym *Symbol, obj Object) {
	self.symbols[sym.Value] = obj
}

func (self *Enviroment) Resolve(sym *Symbol) Object {
	obj := self.symbols[sym.Value]
	if obj == nil && self.parent != nil {
		return self.parent.Resolve(sym)
	}
	return obj
}

func main() {
	os.Exit(repl())
}

func repl() int {
	liner := liner.NewLiner()
	defer liner.Close()

	reader := NewReader()
	env := NewEnviroment()

	for name, fn := range CoreFuncs {
		env.Define(NewSymbol(name), NewPrimFunc(fn))
	}

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
			fmt.Println(obj.Eval(env))
		}
	}

	return 0
}
