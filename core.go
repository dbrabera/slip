package main

import "fmt"

var CoreFuncs = map[string]interface{}{
	// Arithmetic
	"+":   Add,
	"-":   Sub,
	"*":   Mul,
	"/":   Div,
	"rem": Rem,
	//"mod": Mod,
	"inc": Inc,
	"dec": Dec,

	// Relational
	">":  Gt,
	">=": Ge,
	"=":  Eq,
	"!=": Ne,
	"<=": Le,
	"<":  Lt,

	// Logic
	"not": Not,

	// Test
	"nil?":    IsNil,
	"zero?":   IsZero,
	"pos?":    IsPos,
	"neg?":    IsNeg,
	"even?":   IsEven,
	"odd?":    IsOdd,
	"empty?":  IsEmpty,
	"int?":    IsInt,
	"double?": IsDouble,
	"bool?":   IsBool,
	"string?": IsString,
	"list?":   IsList,
	"symbol?": IsSymbol,

	// List
	"first": First,
	"next":  Next,
	"cons":  Cons,

	// IO
	"print": Print,
	//"printf": Printf,
	"println": Println,
	"newline": Newline,
	//"readline": Readline,
}

////////////////////////////////////////////////////////////////////////////////
// Arithmetic

func Add(args ...Object) Object {
	res := args[0].(Number)

	for _, arg := range args[1:] {
		res = res.Add(arg.(Number))
	}

	return res
}

func Sub(args ...Object) Object {
	if len(args) == 1 {
		return NewInt(0).Sub(args[0].(Number))
	}

	res := args[0].(Number)

	for _, arg := range args[1:] {
		res = res.Sub(arg.(Number))
	}

	return res
}

func Mul(args ...Object) Object {
	res := args[0].(Number)

	for _, arg := range args[1:] {
		res = res.Mul(arg.(Number))
	}

	return res
}

func Div(args ...Object) Object {
	if len(args) == 1 {
		return NewDouble(1.0).Div(args[0].(Number))
	}

	res := args[0].(Number)

	for _, arg := range args[1:] {
		res = res.Div(arg.(Number))
	}

	return res
}

func Rem(num Object, div Object) Object {
	return num.(*Int).Rem(div.(*Int))
}

func Inc(x Object) Object {
	return x.(Number).Inc()
}

func Dec(x Object) Object {
	return x.(Number).Dec()
}

////////////////////////////////////////////////////////////////////////////////
// Relational

func Gt(args ...Object) Object {
	x := args[0].(Number)

	for _, arg := range args[1:] {
		y := arg.(Number)
		if !x.Gt(y) {
			return NewBool(false)
		}
		x = y
	}

	return NewBool(true)
}

func Ge(args ...Object) Object {
	x := args[0].(Number)

	for _, arg := range args[1:] {
		y := arg.(Number)
		if !x.Ge(y) {
			return NewBool(false)
		}
		x = y
	}

	return NewBool(true)
}

func Eq(args ...Object) Object {
	x := args[0]

	for _, arg := range args[1:] {
		y := arg
		if !x.Equals(y) {
			return NewBool(false)
		}
		x = y
	}

	return NewBool(true)
}

func Ne(args ...Object) Object {
	return NewBool(!Eq(args...).(*Bool).Value)
}

func Le(args ...Object) Object {
	x := args[0].(Number)

	for _, arg := range args[1:] {
		y := arg.(Number)
		if !x.Le(y) {
			return NewBool(false)
		}
		x = y
	}

	return NewBool(true)
}

func Lt(args ...Object) Object {
	x := args[0].(Number)

	for _, arg := range args[1:] {
		y := arg.(Number)
		if !x.Lt(y) {
			return NewBool(false)
		}
		x = y
	}

	return NewBool(true)
}

////////////////////////////////////////////////////////////////////////////////
// Logic

func Not(x Object) Object {
	return NewBool(x.Equals(&falsecons))
}

////////////////////////////////////////////////////////////////////////////////
// Tests

func IsNil(x Object) Object {
	return NewBool(x.Nil())
}

func IsZero(x Object) Object {
	return NewBool(x.(Number).Equals(NewInt(0)))
}

func IsPos(x Object) Object {
	return NewBool(x.(Number).Gt(NewInt(0)))
}

func IsNeg(x Object) Object {
	return NewBool(x.(Number).Lt(NewInt(0)))
}

func IsEven(x Object) Object {
	return NewBool(x.(*Int).Value%2 == 0)
}

func IsOdd(x Object) Object {
	return NewBool(x.(*Int).Value%2 != 0)
}

func IsEmpty(ls Object) Object {
	return NewBool(ls.(List).First().Nil())
}

func IsInt(x Object) Object {
	_, ok := x.(*Int)
	return NewBool(ok)
}

func IsDouble(x Object) Object {
	_, ok := x.(*Double)
	return NewBool(ok)
}

func IsBool(x Object) Object {
	_, ok := x.(*Bool)
	return NewBool(ok)
}

func IsString(x Object) Object {
	_, ok := x.(*String)
	return NewBool(ok)
}

func IsList(x Object) Object {
	_, ok := x.(*Cell)
	return NewBool(ok)
}

func IsSymbol(x Object) Object {
	_, ok := x.(*Symbol)
	return NewBool(ok)
}

////////////////////////////////////////////////////////////////////////////////
// List

func First(ls Object) Object {
	return ls.(List).First()
}

func Next(ls Object) Object {
	return ls.(List).Next()
}

func Cons(x Object, ls Object) Object {
	return ls.(List).Cons(x)
}

////////////////////////////////////////////////////////////////////////////////
// IO

func Print(args ...Object) Object {
	l := len(args)
	for i, arg := range args {
		fmt.Print(arg)
		if i < l-1 {
			fmt.Print(" ")
		}
	}
	return nil
}

func Println(args ...Object) Object {
	Print(args...)
	fmt.Println()
	return nil
}

func Newline() Object {
	fmt.Println()
	return nil
}
