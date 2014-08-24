package main

import "fmt"

type ProcFn func(List) Object

var CoreProcs = map[string]ProcFn{
	// Arithmetic
	"+":   AddProc,
	"-":   SubProc,
	"*":   MulProc,
	"/":   DivProc,
	"rem": RemProc,
	//"mod": ModProc,
	"inc": IncProc,
	"dec": DecProc,

	// Relational
	">":  GtProc,
	">=": GeProc,
	"=":  EqProc,
	"!=": NeProc,
	"<=": LeProc,
	"<":  LtProc,

	// Test
	"nil?":    IsNilProc,
	"zero?":   IsZeroProc,
	"pos?":    IsPosProc,
	"neg?":    IsNegProc,
	"even?":   IsEvenProc,
	"odd?":    IsOddProc,
	"empty?":  IsEmptyProc,
	"int?":    IsIntProc,
	"double?": IsDoubleProc,
	"bool?":   IsBoolProc,
	"string?": IsStringProc,
	"list?":   IsListProc,
	"symbol?": IsSymbolProc,

	// List
	"first": FirstProc,
	"next":  NextProc,
	"cons":  ConsProc,

	// IO
	"print": PrintProc,
	//"printf": PrintfProc,
	"println": PrintlnProc,
	"newline": NewlineProc,
	//"readline": ReadlineProc,
}

////////////////////////////////////////////////////////////////////////////////
// Arithmetic

func AddProc(args List) Object {
	res := args.First().(Number)

	for !args.Next().Nil() {
		args = args.Next()
		res = res.Add(args.First().(Number))
	}

	return res
}

func SubProc(args List) Object {
	res := args.First().(Number)

	for !args.Next().Nil() {
		args = args.Next()
		res = res.Sub(args.First().(Number))
	}

	return res
}

func MulProc(args List) Object {
	res := args.First().(Number)

	for !args.Next().Nil() {
		args = args.Next()
		res = res.Mul(args.First().(Number))
	}

	return res
}

func DivProc(args List) Object {
	res := args.First().(Number)

	for !args.Next().Nil() {
		args = args.Next()
		res = res.Div(args.First().(Number))
	}

	return res
}

func RemProc(args List) Object {
	num := args.First().(*Int)
	div := args.Second().(*Int)
	return num.Rem(div)
}

func IncProc(args List) Object {
	return args.First().(Number).Inc()
}

func DecProc(args List) Object {
	return args.First().(Number).Dec()
}

////////////////////////////////////////////////////////////////////////////////
// Relational

func GtProc(args List) Object {
	res := true
	x := args.First().(Number)

	for !args.Next().Nil() && res {
		args = args.Next()
		y := args.First().(Number)
		res = x.Gt(y)
		x = y
	}

	return NewBool(res)
}

func GeProc(args List) Object {
	res := true
	x := args.First().(Number)

	for !args.Next().Nil() && res {
		args = args.Next()
		y := args.First().(Number)
		res = x.Ge(y)
		x = y
	}

	return NewBool(res)
}

func EqProc(args List) Object {
	res := true
	x := args.First()

	for !args.Next().Nil() && res {
		args = args.Next()
		res = x.Equals(args.First())
	}

	return NewBool(res)
}

func NeProc(args List) Object {
	return NewBool(!EqProc(args).(*Bool).Value)
}

func LeProc(args List) Object {
	res := true
	x := args.First().(Number)

	for !args.Next().Nil() && res {
		args = args.Next()
		y := args.First().(Number)
		res = x.Le(y)
		x = y
	}

	return NewBool(res)
}

func LtProc(args List) Object {
	res := true
	x := args.First().(Number)

	for !args.Next().Nil() && res {
		args = args.Next()
		y := args.First().(Number)
		res = x.Lt(y)
		x = y
	}

	return NewBool(res)
}

////////////////////////////////////////////////////////////////////////////////
// Logic

func NotProc(args List) Object {
	return NewBool(args.First().Equals(&falsecons))
}

////////////////////////////////////////////////////////////////////////////////
// Tests

func IsNilProc(args List) Object {
	return NewBool(args.First().Nil())
}

func IsZeroProc(args List) Object {
	return NewBool(args.First().(Number).Equals(NewInt(0)))
}

func IsPosProc(args List) Object {
	return NewBool(args.First().(Number).Gt(NewInt(0)))
}

func IsNegProc(args List) Object {
	return NewBool(args.First().(Number).Lt(NewInt(0)))
}

func IsEvenProc(args List) Object {
	x := args.First().(*Int)
	return NewBool(x.Value%2 == 0)
}

func IsOddProc(args List) Object {
	x := args.First().(*Int)
	return NewBool(x.Value%2 != 0)
}

func IsEmptyProc(args List) Object {
	l := args.First().(List)
	return NewBool(l.First().Nil())
}

func IsIntProc(args List) Object {
	_, ok := args.First().(*Int)
	return NewBool(ok)
}

func IsDoubleProc(args List) Object {
	_, ok := args.First().(*Double)
	return NewBool(ok)
}

func IsBoolProc(args List) Object {
	_, ok := args.First().(*Bool)
	return NewBool(ok)
}

func IsStringProc(args List) Object {
	_, ok := args.First().(*String)
	return NewBool(ok)
}

func IsListProc(args List) Object {
	_, ok := args.First().(*Cell)
	return NewBool(ok)
}

func IsSymbolProc(args List) Object {
	_, ok := args.First().(*Symbol)
	return NewBool(ok)
}

////////////////////////////////////////////////////////////////////////////////
// List

func FirstProc(args List) Object {
	return args.First().(List).First()
}

func NextProc(args List) Object {
	return args.First().(List).Next()
}

func ConsProc(args List) Object {
	x := args.First()
	l := args.Second().(List)
	l.Cons(x)
	return l
}

////////////////////////////////////////////////////////////////////////////////
// IO

func PrintProc(args List) Object {
	for !args.Nil() {
		fmt.Print(args.First())
		args = args.Next()
		if !args.Nil() {
			fmt.Print(" ")
		}
	}
	return nil
}

func PrintlnProc(args List) Object {
	PrintProc(args)
	fmt.Println()
	return nil
}

func NewlineProc(args List) Object {
	fmt.Println()
	return nil
}
