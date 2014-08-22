package main

import (
	"fmt"
	"strconv"
)

type Object interface {
	Eval(env *Enviroment) Object
	String() string
}

////////////////////////////////////////////////////////////////////////////////
// Integer

type Int struct {
	Value int
}

func NewInt(value int) *Int {
	return &Int{Value: value}
}

func (self *Int) Eval(env *Enviroment) Object {
	return self
}

func (self *Int) String() string {
	return strconv.Itoa(self.Value)
}

////////////////////////////////////////////////////////////////////////////////
// Double

type Double struct {
	Value float64
}

func NewDouble(value float64) *Double {
	return &Double{Value: value}
}

func (self *Double) Eval(env *Enviroment) Object {
	return self
}

func (self *Double) String() string {
	return strconv.FormatFloat(self.Value, 'f', -1, 64)
}

////////////////////////////////////////////////////////////////////////////////
// Boolean

type Bool struct {
	Value bool
}

var (
	t Bool = Bool{Value: true}
	f Bool = Bool{Value: false}
)

func NewBool(value bool) *Bool {
	if value {
		return &t
	}
	return &f
}

func (self *Bool) Eval(env *Enviroment) Object {
	return self
}

func (self *Bool) String() string {
	return strconv.FormatBool(self.Value)
}

////////////////////////////////////////////////////////////////////////////////
// String

type String struct {
	Value string
}

func NewString(value string) *String {
	return &String{Value: value}
}

func (self *String) Eval(env *Enviroment) Object {
	return self
}

func (self *String) String() string {
	return fmt.Sprintf("\"%s\"", self.Value)
}

////////////////////////////////////////////////////////////////////////////////
// Symbol

type Symbol struct {
	Value string
}

func NewSymbol(value string) *Symbol {
	return &Symbol{Value: value}
}

func (self *Symbol) Eval(env *Enviroment) Object {
	return env.Resolve(self)
}

func (self *Symbol) String() string {
	return self.Value
}

////////////////////////////////////////////////////////////////////////////////
// Cell

type Cell struct {
	Value Object
	Next  *Cell
}

func NewCell(first Object, more *Cell) *Cell {
	return &Cell{Value: first, Next: more}
}

func NewList(objs ...Object) *Cell {
	var first *Cell = nil
	var curr *Cell = nil

	for _, obj := range objs {
		if first == nil {
			first = NewCell(obj, nil)
			curr = first
		} else {
			curr.Next = NewCell(obj, nil)
			curr = curr.Next
		}
	}

	return first
}

func (self *Cell) Eval(env *Enviroment) Object {
	sym, ok := self.Value.(*Symbol)
	if !ok {
		panic("Expecting symbol")
	}

	switch sym.Value {
	case "quote":
		return self.Nth(1)

	case "def":
		if sym, ok := self.Nth(1).(*Symbol); ok {
			env.Define(sym, self.Nth(2))
			return nil
		}
		panic("Expecting symbol")

	case "if":
		test := self.Nth(1).Eval(env)
		if b, ok := test.(*Bool); test != nil && (!ok || b.Value) {
			return self.Nth(2).Eval(env)
		} else if self.Nth(3) != nil {
			return self.Nth(3).Eval(env)
		}
		return nil

	default:
		panic("Unknown symbol")
	}
}

func (self *Cell) String() string {
	return fmt.Sprintf("(%s)", self.string())
}

func (self *Cell) string() string {
	if self.Next == nil {
		return self.Value.String()
	}

	return fmt.Sprintf("%s %s", self.Value.String(), self.Next.string())
}

func (self *Cell) Nth(n int) Object {
	for i, cell := 0, self; i <= n && cell != nil; i, cell = i+1, cell.Next {
		if i == n {
			return cell.Value
		}
	}

	return nil
}
