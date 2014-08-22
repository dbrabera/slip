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
	return env.Symbols[self.Value]
}

func (self *Symbol) String() string {
	return self.Value
}

////////////////////////////////////////////////////////////////////////////////
// Cell

type Cell struct {
	First Object
	More  *Cell
}

func NewCell(first Object, more *Cell) *Cell {
	return &Cell{First: first, More: more}
}

func NewList(objs ...Object) *Cell {
	var first *Cell = nil
	var curr *Cell = nil

	for _, obj := range objs {
		if first == nil {
			first = NewCell(obj, nil)
			curr = first
		} else {
			curr.More = NewCell(obj, nil)
			curr = curr.More
		}
	}

	return first
}

func (self *Cell) Eval(env *Enviroment) Object {
	sym, ok := self.First.(*Symbol)
	if !ok {
		panic("Expecting symbol")
	}

	switch sym.Value {
	case "quote":
		return self.More.First
	default:
		panic("Unknown symbol")
	}
}

func (self *Cell) String() string {
	return fmt.Sprintf("(%s)", self.string())
}

func (self *Cell) string() string {
	if self.More == nil {
		return self.First.String()
	}

	return fmt.Sprintf("%s %s", self.First.String(), self.More.string())
}
