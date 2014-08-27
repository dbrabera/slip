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
	Inc() Number
	Dec() Number
}

type List interface {
	Object
	First() Object
	Second() Object
	Next() List
	Nth(int) Object
	Cons(Object) Object
	IsEmpty() bool
	Vector() []Object
	Len() int
}

type Func interface {
	Object
	Apply(List) Object
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

func (self *Int) Rem(div Number) Number {
	return NewInt(self.Value % div.(*Int).Value)
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

func (self *Int) Inc() Number {
	return NewInt(self.Value + 1)
}

func (self *Int) Dec() Number {
	return NewInt(self.Value - 1)
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

func (self *Double) Inc() Number {
	return NewDouble(self.Value + 1.0)
}

func (self *Double) Dec() Number {
	return NewDouble(self.Value - 1.0)
}

////////////////////////////////////////////////////////////////////////////////
// Boolean

type Bool struct {
	Value bool
}

var (
	truecons  Bool = Bool{Value: true}
	falsecons Bool = Bool{Value: false}
)

func NewBool(value bool) *Bool {
	if value {
		return &truecons
	}
	return &falsecons
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
	if self.IsEmpty() {
		return self
	}

	sym := self.Value.(*Symbol)

	switch sym.Value {
	case "quote":
		return self.Nth(1)

	case "def":
		sym := self.Nth(1).(*Symbol)
		env.Define(sym, self.Nth(2).Eval(env))
		return nil

	case "if":
		test := self.Nth(1).Eval(env)

		if b, ok := test.(*Bool); test != nil && (!ok || b.Value) {
			return self.Nth(2).Eval(env)
		} else if self.Nth(3) != nil {
			return self.Nth(3).Eval(env)
		}

		return nil

	case "fn":
		return NewCompFunc(self.Nth(1).(List), self.Nth(2), env)

	case "defn":
		sym := self.Nth(1).(*Symbol)
		fn := NewCompFunc(self.Nth(2).(List), self.Nth(3), env)
		env.Define(sym, fn)

		return nil

	default:
		proc := env.Resolve(sym).(Func)
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
	if self.IsEmpty() {
		return "()"
	}

	return fmt.Sprintf("(%s)", self.string())
}

func (self *Cell) string() string {
	if self.More == nil {
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
			return (self.More == nil && c.More == nil) || self.More.Equals(c.More)
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

func (self *Cell) Second() Object {
	return self.Next().First()
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

func (self *Cell) Cons(obj Object) Object {
	return NewCell(obj, self)
}

func (self *Cell) IsEmpty() bool {
	return self == nil || self.Value == nil
}

func (self *Cell) Vector() []Object {
	curr := self
	length := self.Len()
	vec := make([]Object, length)

	for i := 0; i < length; i++ {
		vec[i] = curr.Value
		curr = curr.More
	}

	return vec
}

func (self *Cell) Len() int {
	if self == nil {
		return 0
	}

	i := 1
	for curr := self; curr.More != nil; curr = curr.More {
		i += 1
	}
	return i
}

////////////////////////////////////////////////////////////////////////////////
// Primitive Function

type PrimFunc struct {
	fn interface{}
}

func NewPrimFunc(fn interface{}) *PrimFunc {
	return &PrimFunc{fn: fn}
}

func (self *PrimFunc) Eval(env *Enviroment) Object {
	return self
}

func (self *PrimFunc) Apply(args List) Object {
	vargs := args.Vector()

	switch fn := self.fn.(type) {
	case func() Object:
		if len(vargs) != 0 {
			panic("Wrong number of arguments")
		}
		return fn()
	case func(Object) Object:
		if len(vargs) != 1 {
			panic("Wrong number of arguments")
		}
		return fn(vargs[0])
	case func(Object, Object) Object:
		if len(vargs) != 2 {
			panic("Wrong number of arguments")
		}
		return fn(vargs[0], vargs[1])
	case func(...Object) Object:
		return fn(vargs...)
	}

	panic("Invalid primary function")
}

func (self *PrimFunc) String() string {
	return "<function>"
}

func (self *PrimFunc) Nil() bool {
	return self == nil
}

func (self *PrimFunc) Equals(obj Object) bool {
	return false
}

////////////////////////////////////////////////////////////////////////////////
// Compound Function

type CompFunc struct {
	params List
	body   Object
	env    *Enviroment
}

func NewCompFunc(params List, body Object, env *Enviroment) *CompFunc {
	return &CompFunc{params: params, body: body, env: env}
}

func (self *CompFunc) Eval(env *Enviroment) Object {
	return self
}

func (self *CompFunc) Apply(args List) Object {
	// set args in the enviroment
	env := NewChildEnviroment(self.env)
	params := self.params

	for !params.IsEmpty() {
		psym := params.First().(*Symbol)
		env.Define(psym, args.First())
		params = params.Next()
		args = args.Next()
	}

	return self.body.Eval(env)
}

func (self *CompFunc) String() string {
	return "<function>"
}

func (self *CompFunc) Nil() bool {
	return self == nil
}

func (self *CompFunc) Equals(obj Object) bool {
	return false
}
