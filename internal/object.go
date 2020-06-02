package internal

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

type List interface {
	Object
	First() Object
	Second() Object
	Next() List
	Nth(int) Object
	Cons(Object) List
	IsEmpty() bool
	Vector() []Object
	Len() int
}

type Func interface {
	Object
	Apply(List) Object
}

////////////////////////////////////////////////////////////////////////////////
// Number

type Number struct {
	Value float64
}

func NewNumber(value float64) *Number {
	return &Number{Value: value}
}

func (self *Number) Eval(env *Enviroment) Object {
	return self
}

func (self *Number) String() string {
	return strconv.FormatFloat(self.Value, 'f', -1, 64)
}

func (self *Number) Nil() bool {
	return self == nil
}

func (self *Number) Equals(obj Object) bool {
	if n, ok := obj.(*Number); ok {
		return self.Value == n.Value
	}

	return false
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
	if len(objs) == 0 {
		return NewCell(nil, nil)
	}

	head := NewCell(objs[0], nil)
	curr := head

	for _, obj := range objs[1:] {
		curr.More = NewCell(obj, nil)
		curr = curr.More
	}

	return head
}

func (self *Cell) Eval(env *Enviroment) Object {
	if self.IsEmpty() {
		return self
	}

	if sym, ok := self.Value.(*Symbol); ok {
		switch sym.Value {
		case "and":
			var last Object
			for exprs := self.Next(); !exprs.Nil(); exprs = exprs.Next() {
				last = exprs.First().Eval(env)
				if last == nil || last.Nil() || last == &falsecons {
					return last
				}
			}
			return last

		case "def":
			sym := self.Nth(1).(*Symbol)
			env.Define(sym, self.Nth(2).Eval(env))
			return nil

		case "defn":
			sym := self.Nth(1).(*Symbol)
			params := self.Nth(2).(List)
			exprs := self.Next().Next().Next()
			fn := NewCompFunc(params, exprs, env)
			env.Define(sym, fn)

			return nil

		case "do":
			var last Object
			for exprs := self.Next(); !exprs.Nil(); exprs = exprs.Next() {
				last = exprs.First().Eval(env)
			}
			return last

		case "fn":
			params := self.Nth(1).(List)
			exprs := self.Next().Next()
			return NewCompFunc(params, exprs, env)

		case "if":
			test := self.Nth(1).Eval(env)

			if b, ok := test.(*Bool); test != nil && (!ok || b.Value) {
				return self.Nth(2).Eval(env)
			} else if self.Nth(3) != nil {
				return self.Nth(3).Eval(env)
			}

			return nil

		case "let":
			exprs := self.Next().Next()
			params := NewCell(nil, nil)
			paramsCurr := params
			args := NewCell(nil, nil)
			argsCurr := args

			for bindings := self.Second().(List); !bindings.IsEmpty(); bindings = bindings.Next() {
				binding := bindings.First().(List)

				paramsCurr.Value = binding.First()
				paramsCurr.More = NewCell(nil, nil)
				paramsCurr = paramsCurr.More

				argsCurr.Value = binding.Second().Eval(env)
				argsCurr.More = NewCell(nil, nil)
				argsCurr = argsCurr.More
			}

			return NewCompFunc(params, exprs, env).Apply(args)

		case "or":
			var last Object
			for exprs := self.Next(); !exprs.Nil(); exprs = exprs.Next() {
				last = exprs.First().Eval(env)
				if last != nil && !last.Nil() && last != &falsecons {
					return last
				}
			}
			return last

		case "quote":
			return self.Nth(1)
		}
	}

	fn := self.Value.Eval(env).(Func)
	var curr *Cell
	var front *Cell

	for args := self.Next(); !args.Nil(); args = args.Next() {
		if curr == nil {
			curr = NewCell(args.First().Eval(env), nil)
			front = curr
		} else {
			curr.More = NewCell(args.First().Eval(env), nil)
			curr = curr.More
		}
	}

	return fn.Apply(front)
}

func (self *Cell) String() string {
	if self == nil {
		return fmt.Sprint(nil)
	}

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

func (self *Cell) Cons(obj Object) List {
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
	exprs  List
	env    *Enviroment
}

func NewCompFunc(params List, exprs List, env *Enviroment) *CompFunc {
	return &CompFunc{params: params, exprs: exprs.Cons(NewSymbol("do")), env: env}
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

	return self.exprs.Eval(env)
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
