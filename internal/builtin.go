package internal

import (
	"fmt"
	"strings"
)

// BuiltInFuncs contains all the native functions predefined in the global environment.
//
// The documentation for the functions can be found in `docs/builtin.md`
var BuiltInFuncs = map[string]NativeFunc{
	// Arithmetic
	"+":   add,
	"-":   sub,
	"*":   mul,
	"/":   div,
	"mod": mod,
	"inc": inc,
	"dec": dec,

	// Relational
	">":  gt,
	">=": ge,
	"=":  eq,
	"!=": ne,
	"<=": le,
	"<":  lt,

	// Logic
	"not": not,

	// Test
	"bool?":   isBool,
	"list?":   isList,
	"neg?":    isNeg,
	"nil?":    isNil,
	"int?":    isInt,
	"pos?":    isPos,
	"string?": isString,
	"symbol?": isSymbol,
	"zero?":   isZero,

	// IO
	"print":   print,
	"println": println,
}

func add(args ...Value) Value {
	res := Int(0)

	for _, arg := range args {
		res += arg.(Int)
	}

	return res
}

func sub(args ...Value) Value {
	if len(args) == 0 {
		return Int(0)
	} else if len(args) == 1 {
		return Int(-args[0].(Int))
	}

	res := args[0].(Int)

	for _, arg := range args[1:] {
		res -= arg.(Int)
	}

	return res
}

func mul(args ...Value) Value {
	res := Int(1)

	for _, arg := range args {
		res *= arg.(Int)
	}

	return res
}

func div(args ...Value) Value {
	if len(args) == 0 {
		return Int(1)
	} else if len(args) == 1 {
		return Int(1 / args[0].(Int))
	}

	res := args[0].(Int)

	for _, arg := range args[1:] {
		res /= arg.(Int)
	}

	return res
}

func mod(args ...Value) Value {
	return Int(args[0].(Int) % args[1].(Int))
}

func inc(args ...Value) Value {
	return args[0].(Int) + 1
}

func dec(args ...Value) Value {
	return args[0].(Int) - 1
}

func gt(args ...Value) Value {
	if len(args) == 0 {
		return True
	}

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

func ge(args ...Value) Value {
	if len(args) == 0 {
		return True
	}

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

func eq(args ...Value) Value {
	if len(args) == 0 {
		return True
	}

	x := args[0]

	for _, arg := range args[1:] {
		if !x.Equals(arg) {
			return False
		}
	}

	return True
}

func ne(args ...Value) Value {
	return not(eq(args...))
}

func le(args ...Value) Value {
	if len(args) == 0 {
		return True
	}

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

func lt(args ...Value) Value {
	if len(args) == 0 {
		return True
	}

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

func not(args ...Value) Value {
	if args[0] == nil {
		return True
	}
	return NewBool(args[0].Equals(False))
}

func isNil(args ...Value) Value {
	return NewBool(args[0] == nil)
}

func isZero(args ...Value) Value {
	return NewBool(args[0].(Int) == 0)
}

func isPos(args ...Value) Value {
	return NewBool(args[0].(Int) > 0)
}

func isNeg(args ...Value) Value {
	return NewBool(args[0].(Int) < 0)
}

func isInt(args ...Value) Value {
	_, ok := args[0].(Int)
	return NewBool(ok)
}

func isBool(args ...Value) Value {
	_, ok := args[0].(Bool)
	return NewBool(ok)
}

func isString(args ...Value) Value {
	_, ok := args[0].(String)
	return NewBool(ok)
}

func isList(args ...Value) Value {
	_, ok := args[0].(List)
	return NewBool(ok)
}

func isSymbol(args ...Value) Value {
	_, ok := args[0].(Symbol)
	return NewBool(ok)
}

func isEmpty(args ...Value) Value {
	return NewBool(args[0].(List).IsEmpty())
}

func print(args ...Value) Value {
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

func println(args ...Value) Value {
	print(args...)
	fmt.Println()
	return nil
}
