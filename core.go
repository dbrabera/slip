package main

import (
	"fmt"
	"math"
)

var CoreFuncs = map[string]interface{}{
	// Arithmetic
	"+":   Add,
	"-":   Sub,
	"*":   Mul,
	"/":   Div,
	"rem": Rem,
	"mod": Mod,
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
	"empty?":  IsEmpty,
	"number?": IsNumber,
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
	res := args[0].(*Number).Value

	for _, arg := range args[1:] {
		res += arg.(*Number).Value
	}

	return NewNumber(res)
}

func Sub(args ...Object) Object {
	if len(args) == 1 {
		return NewNumber(-args[0].(*Number).Value)
	}

	res := args[0].(*Number).Value

	for _, arg := range args[1:] {
		res -= arg.(*Number).Value
	}

	return NewNumber(res)
}

func Mul(args ...Object) Object {
	res := args[0].(*Number).Value

	for _, arg := range args[1:] {
		res *= arg.(*Number).Value
	}

	return NewNumber(res)
}

func Div(args ...Object) Object {
	if len(args) == 1 {
		return NewNumber(1.0 / args[0].(*Number).Value)
	}

	res := args[0].(*Number).Value

	for _, arg := range args[1:] {
		res /= arg.(*Number).Value
	}

	return NewNumber(res)
}

func Rem(num Object, div Object) Object {
	return NewNumber(math.Remainder(num.(*Number).Value, div.(*Number).Value))
}

func Mod(num Object, div Object) Object {
	return NewNumber(math.Mod(num.(*Number).Value, div.(*Number).Value))
}

func Inc(x Object) Object {
	return NewNumber(x.(*Number).Value + 1.0)
}

func Dec(x Object) Object {
	return NewNumber(x.(*Number).Value - 1.0)
}

////////////////////////////////////////////////////////////////////////////////
// Relational

func Gt(args ...Object) Object {
	x := args[0].(*Number).Value

	for _, arg := range args[1:] {
		y := arg.(*Number).Value
		if x <= y {
			return NewBool(false)
		}
		x = y
	}

	return NewBool(true)
}

func Ge(args ...Object) Object {
	x := args[0].(*Number).Value

	for _, arg := range args[1:] {
		y := arg.(*Number).Value
		if x < y {
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
	x := args[0].(*Number).Value

	for _, arg := range args[1:] {
		y := arg.(*Number).Value
		if x > y {
			return NewBool(false)
		}
		x = y
	}

	return NewBool(true)
}

func Lt(args ...Object) Object {
	x := args[0].(*Number).Value

	for _, arg := range args[1:] {
		y := arg.(*Number).Value
		if x >= y {
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
	return NewBool(x.(*Number).Value == 0.0)
}

func IsPos(x Object) Object {
	return NewBool(x.(*Number).Value > 0)
}

func IsNeg(x Object) Object {
	return NewBool(x.(*Number).Value < 0)
}

func IsEmpty(ls Object) Object {
	return NewBool(ls.(List).First().Nil())
}

func IsNumber(x Object) Object {
	_, ok := x.(*Number)
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
