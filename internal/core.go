package internal

import (
	"fmt"
	"strings"
)

var CoreFuncs = map[string]NativeFunc{
	// Arithmetic
	"+":   Add,
	"-":   Sub,
	"*":   Mul,
	"/":   Div,
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
	"bool?":   IsBool,
	"list?":   IsList,
	"neg?":    IsNeg,
	"nil?":    IsNil,
	"int?":    IsInt,
	"pos?":    IsPos,
	"string?": IsString,
	"symbol?": IsSymbol,
	"zero?":   IsZero,

	// IO
	"print":   Print,
	"println": Println,
}

func Add(args ...Value) Value {
	res := args[0].(Int)

	for _, arg := range args[1:] {
		res += arg.(Int)
	}

	return res
}

func Sub(args ...Value) Value {
	if len(args) == 1 {
		return Int(-args[0].(Int))
	}

	res := args[0].(Int)

	for _, arg := range args[1:] {
		res -= arg.(Int)
	}

	return res
}

func Mul(args ...Value) Value {
	res := args[0].(Int)

	for _, arg := range args[1:] {
		res *= arg.(Int)
	}

	return res
}

func Div(args ...Value) Value {
	if len(args) == 1 {
		return Int(1.0 / args[0].(Int))
	}

	res := args[0].(Int)

	for _, arg := range args[1:] {
		res /= arg.(Int)
	}

	return res
}

func Mod(args ...Value) Value {
	return Int(args[0].(Int) % args[1].(Int))
}

func Inc(args ...Value) Value {
	return args[0].(Int) + 1
}

func Dec(args ...Value) Value {
	return args[0].(Int) - 1
}

func Gt(args ...Value) Value {
	x := args[0].(Int)

	for _, arg := range args[1:] {
		y := arg.(Int)
		if x <= y {
			return False
		}
		x = y
	}

	return True
}

func Ge(args ...Value) Value {
	x := args[0].(Int)

	for _, arg := range args[1:] {
		y := arg.(Int)
		if x < y {
			return False
		}
		x = y
	}

	return True
}

func Eq(args ...Value) Value {
	x := args[0]

	for _, arg := range args[1:] {
		if !x.Equals(arg) {
			return False
		}
	}

	return True
}

func Ne(args ...Value) Value {
	return Not(Eq(args...))
}

func Le(args ...Value) Value {
	x := args[0].(Int)

	for _, arg := range args[1:] {
		y := arg.(Int)
		if x > y {
			return False
		}
		x = y
	}

	return True
}

func Lt(args ...Value) Value {
	x := args[0].(Int)

	for _, arg := range args[1:] {
		y := arg.(Int)
		if x >= y {
			return False
		}
		x = y
	}

	return True
}

func Not(args ...Value) Value {
	if args[0] == nil {
		return True
	}
	return NewBool(args[0].Equals(False))
}

func IsNil(args ...Value) Value {
	return NewBool(args[0] == nil)
}

func IsZero(args ...Value) Value {
	return NewBool(args[0].(Int) == 0)
}

func IsPos(args ...Value) Value {
	return NewBool(args[0].(Int) > 0)
}

func IsNeg(args ...Value) Value {
	return NewBool(args[0].(Int) < 0)
}

func IsInt(args ...Value) Value {
	_, ok := args[0].(Int)
	return NewBool(ok)
}

func IsBool(args ...Value) Value {
	_, ok := args[0].(Bool)
	return NewBool(ok)
}

func IsString(args ...Value) Value {
	_, ok := args[0].(String)
	return NewBool(ok)
}

func IsList(args ...Value) Value {
	_, ok := args[0].(List)
	return NewBool(ok)
}

func IsSymbol(args ...Value) Value {
	_, ok := args[0].(Symbol)
	return NewBool(ok)
}

func IsEmpty(args ...Value) Value {
	return NewBool(args[0].(List).IsEmpty())
}

func Print(args ...Value) Value {
	elems := make([]string, len(args))
	for i, arg := range args {
		if arg == nil {
			elems[i] = "nil"
		} else {
			elems[i] = fmt.Sprint(arg)
		}
	}
	fmt.Print(strings.Join(elems, " "))
	return nil
}

func Println(args ...Value) Value {
	Print(args...)
	fmt.Println()
	return nil
}
