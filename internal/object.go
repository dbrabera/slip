package internal

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
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

type Number struct {
	Value float64
}

func NewNumber(value float64) *Number {
	return &Number{Value: value}
}

func (n *Number) Eval(env *Enviroment) Object {
	return n
}

func (n *Number) String() string {
	return strconv.FormatFloat(n.Value, 'f', -1, 64)
}

func (n *Number) Nil() bool {
	return n == nil
}

func (n *Number) Equals(obj Object) bool {
	if o, ok := obj.(*Number); ok {
		return n.Value == o.Value
	}

	return false
}

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

func (b *Bool) Eval(env *Enviroment) Object {
	return b
}

func (b *Bool) String() string {
	return strconv.FormatBool(b.Value)
}

func (b *Bool) Nil() bool {
	return b == nil
}

func (b *Bool) Equals(obj Object) bool {
	if o, ok := obj.(*Bool); ok {
		return b.Value == o.Value
	}

	return false
}

type String struct {
	Value string
}

func NewString(value string) *String {
	return &String{Value: value}
}

func (s *String) Eval(env *Enviroment) Object {
	return s
}

func (s *String) String() string {
	return fmt.Sprintf("\"%s\"", s.Value)
}

func (s *String) Nil() bool {
	return s == nil
}

func (s *String) Equals(obj Object) bool {
	if o, ok := obj.(*String); ok {
		return s.Value == o.Value
	}

	return false
}

type Symbol struct {
	Value string
}

func NewSymbol(value string) *Symbol {
	return &Symbol{Value: value}
}

func (s *Symbol) Eval(env *Enviroment) Object {
	return env.Resolve(s)
}

func (s *Symbol) String() string {
	return s.Value
}

func (s *Symbol) Nil() bool {
	return s == nil
}

func (s *Symbol) Equals(obj Object) bool {
	if o, ok := obj.(*Symbol); ok {
		return s.Value == o.Value
	}

	return false
}

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

func (c *Cell) Eval(env *Enviroment) Object {
	if c.IsEmpty() {
		return c
	}

	if sym, ok := c.Value.(*Symbol); ok {
		switch sym.Value {
		case "and":
			var last Object
			for exprs := c.Next(); !exprs.Nil(); exprs = exprs.Next() {
				last = exprs.First().Eval(env)
				if last == nil || last.Nil() || last == &falsecons {
					return last
				}
			}
			return last

		case "def":
			sym := c.Nth(1).(*Symbol)
			env.Define(sym, c.Nth(2).Eval(env))
			return nil

		case "defn":
			sym := c.Nth(1).(*Symbol)
			params := c.Nth(2).(List)
			exprs := c.Next().Next().Next()
			fn := NewCompFunc(params, exprs, env)
			env.Define(sym, fn)

			return nil

		case "do":
			var last Object
			for exprs := c.Next(); !exprs.Nil(); exprs = exprs.Next() {
				last = exprs.First().Eval(env)
			}
			return last

		case "fn":
			params := c.Nth(1).(List)
			exprs := c.Next().Next()
			return NewCompFunc(params, exprs, env)

		case "if":
			test := c.Nth(1).Eval(env)

			if b, ok := test.(*Bool); test != nil && (!ok || b.Value) {
				return c.Nth(2).Eval(env)
			} else if c.Nth(3) != nil {
				return c.Nth(3).Eval(env)
			}

			return nil

		case "let":
			exprs := c.Next().Next()
			params := NewCell(nil, nil)
			paramsCurr := params
			args := NewCell(nil, nil)
			argsCurr := args

			for bindings := c.Second().(List); !bindings.IsEmpty(); bindings = bindings.Next() {
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
			for exprs := c.Next(); !exprs.Nil(); exprs = exprs.Next() {
				last = exprs.First().Eval(env)
				if last != nil && !last.Nil() && last != &falsecons {
					return last
				}
			}
			return last

		case "quote":
			return c.Nth(1)
		}
	}

	fn := c.Value.Eval(env).(Func)
	var curr *Cell
	var front *Cell

	for args := c.Next(); !args.Nil(); args = args.Next() {
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

func (c *Cell) String() string {
	if c == nil {
		return fmt.Sprint(nil)
	}

	if c.IsEmpty() {
		return "()"
	}

	return fmt.Sprintf("(%s)", c.string())
}

func (c *Cell) string() string {
	if c.More == nil {
		return c.Value.String()
	}

	return fmt.Sprintf("%s %s", c.Value.String(), c.More.string())
}

func (c *Cell) Nil() bool {
	return c == nil
}

func (c *Cell) Equals(obj Object) bool {
	if o, ok := obj.(*Cell); ok {
		if c.Value.Equals(o.Value) {
			return (c.More == nil && o.More == nil) || c.More.Equals(o.More)
		}
	}

	return false
}

func (c *Cell) First() Object {
	if c == nil {
		return nil
	}

	return c.Value
}

func (c *Cell) Second() Object {
	return c.Next().First()
}

func (c *Cell) Next() List {
	if c == nil {
		return nil
	}

	return c.More
}

func (c *Cell) Nth(n int) Object {
	if c == nil {
		return nil
	}

	for i, cell := 0, c; i <= n && cell != nil; i, cell = i+1, cell.More {
		if i == n {
			return cell.Value
		}
	}

	return nil
}

func (c *Cell) Cons(obj Object) List {
	return NewCell(obj, c)
}

func (c *Cell) IsEmpty() bool {
	return c == nil || c.Value == nil
}

func (c *Cell) Vector() []Object {
	curr := c
	length := c.Len()
	vec := make([]Object, length)

	for i := 0; i < length; i++ {
		vec[i] = curr.Value
		curr = curr.More
	}

	return vec
}

func (c *Cell) Len() int {
	if c == nil {
		return 0
	}

	i := 1
	for curr := c; curr.More != nil; curr = curr.More {
		i += 1
	}
	return i
}

type PrimFunc struct {
	fn interface{}
}

func NewPrimFunc(fn interface{}) *PrimFunc {
	return &PrimFunc{fn: fn}
}

func (pf *PrimFunc) Eval(env *Enviroment) Object {
	return pf
}

func (pf *PrimFunc) Apply(args List) Object {
	vargs := args.Vector()

	switch fn := pf.fn.(type) {
	case func() Object:
		if len(vargs) != 0 {
			panic(errors.New("wrong number of arguments"))
		}
		return fn()
	case func(Object) Object:
		if len(vargs) != 1 {
			panic(errors.New("wrong number of arguments"))
		}
		return fn(vargs[0])
	case func(Object, Object) Object:
		if len(vargs) != 2 {
			panic(errors.New("wrong number of arguments"))
		}
		return fn(vargs[0], vargs[1])
	case func(...Object) Object:
		return fn(vargs...)
	}

	panic(errors.New("invalid function"))
}

func (pf *PrimFunc) String() string {
	return "<function>"
}

func (pf *PrimFunc) Nil() bool {
	return pf == nil
}

func (pf *PrimFunc) Equals(obj Object) bool {
	return false
}

type CompFunc struct {
	params List
	exprs  List
	env    *Enviroment
}

func NewCompFunc(params List, exprs List, env *Enviroment) *CompFunc {
	return &CompFunc{params: params, exprs: exprs.Cons(NewSymbol("do")), env: env}
}

func (cf *CompFunc) Eval(env *Enviroment) Object {
	return cf
}

func (cf *CompFunc) Apply(args List) Object {
	// set args in the enviroment
	env := NewChildEnviroment(cf.env)
	params := cf.params

	for !params.IsEmpty() {
		psym := params.First().(*Symbol)
		env.Define(psym, args.First())
		params = params.Next()
		args = args.Next()
	}

	return cf.exprs.Eval(env)
}

func (cf *CompFunc) String() string {
	return "<function>"
}

func (cf *CompFunc) Nil() bool {
	return cf == nil
}

func (cf *CompFunc) Equals(obj Object) bool {
	return false
}
