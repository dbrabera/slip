package main

import (
	"fmt"
	"io"
	"os"

	"github.com/peterh/liner"
)

type Enviroment struct {
	symbols map[string]Object
}

func NewEnviroment() *Enviroment {
	return &Enviroment{symbols: make(map[string]Object)}
}

func (self *Enviroment) Define(sym *Symbol, obj Object) {
	self.symbols[sym.Value] = obj
}

func (self *Enviroment) Resolve(sym *Symbol) Object {
	return self.symbols[sym.Value]
}

func main() {
	os.Exit(repl())
}

func repl() int {
	liner := liner.NewLiner()
	defer liner.Close()

	reader := NewReader()
	env := NewEnviroment()

	env.Define(NewSymbol("+"), NewProcedure(AddProc))
	env.Define(NewSymbol("-"), NewProcedure(SubProc))
	env.Define(NewSymbol("*"), NewProcedure(MulProc))
	env.Define(NewSymbol("/"), NewProcedure(DivProc))
	env.Define(NewSymbol(">"), NewProcedure(GtProc))
	env.Define(NewSymbol(">="), NewProcedure(GeProc))
	env.Define(NewSymbol("="), NewProcedure(EqProc))
	env.Define(NewSymbol("<="), NewProcedure(LeProc))
	env.Define(NewSymbol("<"), NewProcedure(LtProc))
	env.Define(NewSymbol("first"), NewProcedure(FirstProc))
	env.Define(NewSymbol("next"), NewProcedure(NextProc))

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
