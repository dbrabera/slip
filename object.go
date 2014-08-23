package main

import (
	"fmt"
	"strconv"
)

type Object interface {
	Eval(env *Enviroment) Object
	String() string
	Equals(Object) bool
	Nil() bool
}

type Number interface {
	Object
	Add(Number) Number
	Sub(Number) Number
	Mul(Number) Number
	Div(Number) Number
	Gt(Number) bool
	Ge(Number) bool
	Le(Number) bool
	Lt(Number) bool
}

type List interface {
	Object
	First() Object
	Next() List
	Nth(int) Object
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

func (self *Int) Nil() bool {
	return self == nil
}

func (self *Int) Equals(obj Object) bool {
	if n, ok := obj.(*Double); ok {
		return float64(self.Value) == n.Value
	}

	if n, ok := obj.(*Int); ok {
		return self.Value == n.Value
	}

	return false
}

func (self *Int) Add(n Number) Number {
	if n, ok := n.(*Double); ok {
		return NewDouble(float64(self.Value) + n.Value)
	}
	return NewInt(self.Value + n.(*Int).Value)
}

func (self *Int) Sub(n Number) Number {
	if n, ok := n.(*Double); ok {
		return NewDouble(float64(self.Value) - n.Value)
	}
	return NewInt(self.Value - n.(*Int).Value)
}

func (self *Int) Mul(n Number) Number {
	if n, ok := n.(*Double); ok {
		return NewDouble(float64(self.Value) * n.Value)
	}
	return NewInt(self.Value * n.(*Int).Value)
}

func (self *Int) Div(n Number) Number {
	if n, ok := n.(*Double); ok {
		return NewDouble(float64(self.Value) / n.Value)
	}
	return NewInt(self.Value / n.(*Int).Value)
}

func (self *Int) Gt(n Number) bool {
	if n, ok := n.(*Double); ok {
		return float64(self.Value) > n.Value
	}
	return self.Value > n.(*Int).Value
}

func (self *Int) Ge(n Number) bool {
	if n, ok := n.(*Double); ok {
		return float64(self.Value) >= n.Value
	}
	return self.Value >= n.(*Int).Value
}

func (self *Int) Le(n Number) bool {
	if n, ok := n.(*Double); ok {
		return float64(self.Value) <= n.Value
	}
	return self.Value <= n.(*Int).Value
}

func (self *Int) Lt(n Number) bool {
	if n, ok := n.(*Double); ok {
		return float64(self.Value) < n.Value
	}
	return self.Value < n.(*Int).Value
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

func (self *Double) Nil() bool {
	return self == nil
}

func (self *Double) Equals(obj Object) bool {
	if n, ok := obj.(*Double); ok {
		return self.Value == n.Value
	}

	if n, ok := obj.(*Int); ok {
		return self.Value == float64(n.Value)
	}

	return false
}

func (self *Double) Add(n Number) Number {
	if n, ok := n.(*Int); ok {
		return NewDouble(self.Value + float64(n.Value))
	}
	return NewDouble(self.Value + n.(*Double).Value)
}

func (self *Double) Sub(n Number) Number {
	if n, ok := n.(*Int); ok {
		return NewDouble(self.Value - float64(n.Value))
	}
	return NewDouble(self.Value - n.(*Double).Value)
}

func (self *Double) Mul(n Number) Number {
	if n, ok := n.(*Int); ok {
		return NewDouble(self.Value * float64(n.Value))
	}
	return NewDouble(self.Value * n.(*Double).Value)
}

func (self *Double) Div(n Number) Number {
	if n, ok := n.(*Int); ok {
		return NewDouble(self.Value / float64(n.Value))
	}
	return NewDouble(self.Value / n.(*Double).Value)
}

func (self *Double) Gt(n Number) bool {
	if n, ok := n.(*Int); ok {
		return self.Value > float64(n.Value)
	}
	return self.Value > n.(*Double).Value
}

func (self *Double) Ge(n Number) bool {
	if n, ok := n.(*Int); ok {
		return self.Value >= float64(n.Value)
	}
	return self.Value >= n.(*Double).Value
}

func (self *Double) Le(n Number) bool {
	if n, ok := n.(*Int); ok {
		return self.Value <= float64(n.Value)
	}
	return self.Value <= n.(*Double).Value
}

func (self *Double) Lt(n Number) bool {
	if n, ok := n.(*Int); ok {
		return self.Value < float64(n.Value)
	}
	return self.Value < n.(*Double).Value
}

////////////////////////////////////////////////////////////////////////////////
// Boolean

type Bool struct {
	Value bool
}

var (
	truecons Bool = Bool{Value: true}
	falscons Bool = Bool{Value: false}
)

func NewBool(value bool) *Bool {
	if value {
		return &truecons
	}
	return &falscons
}

func (self *Bool) Eval(env *Enviroment) Object {
	return self
}

func (self *Bool) String() string {
	return strconv.FormatBool(self.Value)
}

func (self *Bool) Nil() bool {
	return self == nil
}

func (self *Bool) Equals(obj Object) bool {
	if b, ok := obj.(*Bool); ok {
		return self.Value == b.Value
	}

	return false
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

func (self *String) Nil() bool {
	return self == nil
}

func (self *String) Equals(obj Object) bool {
	if s, ok := obj.(*String); ok {
		return self.Value == s.Value
	}

	return false
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

func (self *Symbol) Nil() bool {
	return self == nil
}

func (self *Symbol) Equals(obj Object) bool {
	if s, ok := obj.(*Symbol); ok {
		return self.Value == s.Value
	}

	return false
}

////////////////////////////////////////////////////////////////////////////////
// Cell

type Cell struct {
	Value Object
	More  *Cell
}

func NewCell(first Object, more *Cell) *Cell {
	return &Cell{Value: first, More: more}
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
	sym := self.Value.(*Symbol)

	switch sym.Value {
	case "quote":
		return self.Nth(1)

	case "def":
		sym := self.Nth(1).(*Symbol)
		env.Define(sym, self.Nth(2))
		return nil

	case "if":
		test := self.Nth(1).Eval(env)

		if b, ok := test.(*Bool); test != nil && (!ok || b.Value) {
			return self.Nth(2).Eval(env)
		} else if self.Nth(3) != nil {
			return self.Nth(3).Eval(env)
		}

		return nil

	default:
		proc := env.Resolve(sym).(*Procedure)
		args := self.Next()
		var curr *Cell
		var front *Cell

		for !args.Nil() {
			if curr == nil {
				curr = NewCell(args.First().Eval(env), nil)
				front = curr
			} else {
				curr.More = NewCell(args.First().Eval(env), nil)
				curr = curr.More
			}
			args = args.Next()
		}

		return proc.Apply(front)
	}
}

func (self *Cell) String() string {
	return fmt.Sprintf("(%s)", self.string())
}

func (self *Cell) string() string {
	if self.Next == nil {
		return self.Value.String()
	}

	return fmt.Sprintf("%s %s", self.Value.String(), self.More.string())
}

func (self *Cell) Nil() bool {
	return self == nil
}

func (self *Cell) Equals(obj Object) bool {
	if c, ok := obj.(*Cell); ok {
		if self.Value.Equals(c.Value) {
			return self.More.Equals(c.More)
		}
	}

	return false
}

func (self *Cell) First() Object {
	if self == nil {
		return nil
	}

	return self.Value
}

func (self *Cell) Next() List {
	if self == nil {
		return nil
	}

	return self.More
}

func (self *Cell) Nth(n int) Object {
	if self == nil {
		return nil
	}

	for i, cell := 0, self; i <= n && cell != nil; i, cell = i+1, cell.More {
		if i == n {
			return cell.Value
		}
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Procedure

type Procedure struct {
	Value ProcFn
}

func NewProcedure(value ProcFn) *Procedure {
	return &Procedure{Value: value}
}

func (self *Procedure) Eval(env *Enviroment) Object {
	return self
}

func (self *Procedure) Apply(args List) Object {
	return self.Value(args)
}

func (self *Procedure) String() string {
	return "<procedure>"
}

func (self *Procedure) Nil() bool {
	return self == nil
}

func (self *Procedure) Equals(obj Object) bool {
	return false
}
