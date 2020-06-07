package internal

import (
	"fmt"
	"strings"
)

type Value interface {
	Eval(env *Enviroment) Value
	String() string
	Equals(Value) bool
}

type Int int64

func NewInt(i int64) Int {
	return Int(i)
}

func (i Int) Eval(env *Enviroment) Value {
	return i
}

func (i Int) String() string {
	return fmt.Sprint(int(i))
}

func (i Int) IsNil() bool {
	return false
}

func (i Int) Equals(val Value) bool {
	if v, ok := val.(Int); ok {
		return i == v
	}
	return false
}

type Bool bool

var (
	True  = Bool(true)
	False = Bool(false)
)

func NewBool(b bool) Bool {
	if b {
		return True
	}
	return False
}

func (b Bool) Eval(env *Enviroment) Value {
	return b
}

func (b Bool) String() string {
	return fmt.Sprint(bool(b))
}

func (b Bool) Equals(val Value) bool {
	if v, ok := val.(Bool); ok {
		return b == v
	}
	return false
}

type String string

func NewString(s string) String {
	return String(s)
}

func (s String) Eval(env *Enviroment) Value {
	return s
}

func (s String) String() string {
	return fmt.Sprintf("%q", string(s))
}

func (s String) Equals(val Value) bool {
	if v, ok := val.(String); ok {
		return s == v
	}
	return false
}

type Symbol string

func NewSymbol(s string) Symbol {
	return Symbol(s)
}

func (s Symbol) Eval(env *Enviroment) Value {
	return env.Resolve(s)
}

func (s Symbol) String() string {
	return string(s)
}

func (s Symbol) Equals(val Value) bool {
	if v, ok := val.(Symbol); ok {
		return s == v
	}
	return false
}

type List []Value

func NewList(vals ...Value) List {
	return List(vals)
}

func (l List) Eval(env *Enviroment) Value {
	if l.IsEmpty() {
		return nil
	}

	if sym, ok := l[0].(Symbol); ok {
		switch sym {
		case "and":
			var last Value
			for _, expr := range l[1:] {
				last = expr.Eval(env)
				if last == nil || last.Equals(False) {
					return last
				}
			}
			return last

		case "def":
			env.Define(l[1].(Symbol), l[2].Eval(env))
			return nil

		case "defn":
			sym := l[1].(Symbol)
			params := l[2].(List)
			exprs := l[3:]
			fn := NewFunc(params, exprs, env)
			env.Define(sym, fn)
			return nil

		case "do":
			var last Value
			for _, expr := range l[1:] {
				last = expr.Eval(env)
			}
			return last

		case "fn":
			params := l[1].(List)
			exprs := l[2:]
			return NewFunc(params, exprs, env)

		case "if":
			test := l[1].Eval(env)

			if b, ok := test.(Bool); ok && bool(b) {
				return l[2].Eval(env)
			} else if len(l) >= 4 {
				return l[3].Eval(env)
			}

			return nil

		case "let":
			bindings := l[1].(List)
			exprs := NewList(l[2].(List))

			params := NewList()
			args := NewList()

			for _, binding := range bindings {
				b := binding.(List)
				params = append(params, b[0])
				args = append(args, b[1])
			}

			return NewFunc(params, exprs, env).Apply(args)

		case "or":
			var last Value
			for _, expr := range l[1:] {
				last = expr.Eval(env)
				if last != nil && !last.Equals(False) {
					return last
				}
			}
			return last

		case "quote":
			return l[1]
		}
	}

	args := NewList()
	for _, expr := range l[1:] {
		args = append(args, expr.Eval(env))
	}

	switch fn := l[0].Eval(env).(type) {
	case *Func:
		return fn.Apply(args)
	case NativeFunc:
		return fn.Apply(args)
	default:
		if fn == nil {
			return nil
		}
		panic(fmt.Errorf("unexpected type '%T'", fn))
	}
}

func (l List) String() string {
	if l.IsEmpty() {
		return "()"
	}

	elems := make([]string, len(l))
	for i, v := range l {
		elems[i] = v.String()
	}

	return "(" + strings.Join(elems, " ") + ")"
}

func (l List) Equals(val Value) bool {
	if v, ok := val.(List); ok {
		if len(l) != len(v) {
			return false
		}

		for i, e := range l {
			if !e.Equals(v[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func (l List) IsEmpty() bool {
	return len(l) == 0
}

func (l List) Len() int {
	return len(l)
}

type NativeFunc func(...Value) Value

func NewNativeFunc(fn func(...Value) Value) NativeFunc {
	return NativeFunc(fn)
}

func (nf NativeFunc) Eval(env *Enviroment) Value {
	return nf
}

func (nf NativeFunc) Apply(args List) Value {
	return nf(args...)
}

func (nf NativeFunc) String() string {
	return "<function>"
}

func (nf NativeFunc) Equals(val Value) bool {
	return false
}

type Func struct {
	params List
	exprs  List
	env    *Enviroment
}

func NewFunc(params List, exprs List, env *Enviroment) *Func {
	list := NewList(NewSymbol("do"))
	list = append(list, exprs...)
	return &Func{params: params, exprs: list, env: env}
}

func (f *Func) Eval(env *Enviroment) Value {
	return f
}

func (f *Func) Apply(args List) Value {
	env := NewChildEnviroment(f.env)
	for i, param := range f.params {
		env.Define(param.(Symbol), args[i])
	}
	return f.exprs.Eval(env)
}

func (f *Func) String() string {
	return "<function>"
}

func (f *Func) Equals(val Value) bool {
	if v, ok := val.(*Func); ok {
		return f == v
	}
	return false
}
