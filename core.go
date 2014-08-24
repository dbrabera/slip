package main

import "fmt"

var CoreFuncs = map[string]PrimFunc{
	// Arithmetic
	"+":   AddFunc,
	"-":   SubFunc,
	"*":   MulFunc,
	"/":   DivFunc,
	"rem": RemFunc,
	//"mod": ModFunc,
	"inc": IncFunc,
	"dec": DecFunc,

	// Relational
	">":  GtFunc,
	">=": GeFunc,
	"=":  EqFunc,
	"!=": NeFunc,
	"<=": LeFunc,
	"<":  LtFunc,

	// Test
	"nil?":    IsNilFunc,
	"zero?":   IsZeroFunc,
	"pos?":    IsPosFunc,
	"neg?":    IsNegFunc,
	"even?":   IsEvenFunc,
	"odd?":    IsOddFunc,
	"empty?":  IsEmptyFunc,
	"int?":    IsIntFunc,
	"double?": IsDoubleFunc,
	"bool?":   IsBoolFunc,
	"string?": IsStringFunc,
	"list?":   IsListFunc,
	"symbol?": IsSymbolFunc,

	// List
	"first": FirstFunc,
	"next":  NextFunc,
	"cons":  ConsFunc,

	// IO
	"print": PrintFunc,
	//"printf": PrintfFunc,
	"println": PrintlnFunc,
	"newline": NewlineFunc,
	//"readline": ReadlineFunc,
}

////////////////////////////////////////////////////////////////////////////////
// Arithmetic

func AddFunc(args List) Object {
	res := args.First().(Number)

	for !args.Next().Nil() {
		args = args.Next()
		res = res.Add(args.First().(Number))
	}

	return res
}

func SubFunc(args List) Object {
	res := args.First().(Number)

	for !args.Next().Nil() {
		args = args.Next()
		res = res.Sub(args.First().(Number))
	}

	return res
}

func MulFunc(args List) Object {
	res := args.First().(Number)

	for !args.Next().Nil() {
		args = args.Next()
		res = res.Mul(args.First().(Number))
	}

	return res
}

func DivFunc(args List) Object {
	res := args.First().(Number)

	for !args.Next().Nil() {
		args = args.Next()
		res = res.Div(args.First().(Number))
	}

	return res
}

func RemFunc(args List) Object {
	num := args.First().(*Int)
	div := args.Second().(*Int)
	return num.Rem(div)
}

func IncFunc(args List) Object {
	return args.First().(Number).Inc()
}

func DecFunc(args List) Object {
	return args.First().(Number).Dec()
}

////////////////////////////////////////////////////////////////////////////////
// Relational

func GtFunc(args List) Object {
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

func GeFunc(args List) Object {
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

func EqFunc(args List) Object {
	res := true
	x := args.First()

	for !args.Next().Nil() && res {
		args = args.Next()
		res = x.Equals(args.First())
	}

	return NewBool(res)
}

func NeFunc(args List) Object {
	return NewBool(!EqFunc(args).(*Bool).Value)
}

func LeFunc(args List) Object {
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

func LtFunc(args List) Object {
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

func NotFunc(args List) Object {
	return NewBool(args.First().Equals(&falsecons))
}

////////////////////////////////////////////////////////////////////////////////
// Tests

func IsNilFunc(args List) Object {
	return NewBool(args.First().Nil())
}

func IsZeroFunc(args List) Object {
	return NewBool(args.First().(Number).Equals(NewInt(0)))
}

func IsPosFunc(args List) Object {
	return NewBool(args.First().(Number).Gt(NewInt(0)))
}

func IsNegFunc(args List) Object {
	return NewBool(args.First().(Number).Lt(NewInt(0)))
}

func IsEvenFunc(args List) Object {
	x := args.First().(*Int)
	return NewBool(x.Value%2 == 0)
}

func IsOddFunc(args List) Object {
	x := args.First().(*Int)
	return NewBool(x.Value%2 != 0)
}

func IsEmptyFunc(args List) Object {
	l := args.First().(List)
	return NewBool(l.First().Nil())
}

func IsIntFunc(args List) Object {
	_, ok := args.First().(*Int)
	return NewBool(ok)
}

func IsDoubleFunc(args List) Object {
	_, ok := args.First().(*Double)
	return NewBool(ok)
}

func IsBoolFunc(args List) Object {
	_, ok := args.First().(*Bool)
	return NewBool(ok)
}

func IsStringFunc(args List) Object {
	_, ok := args.First().(*String)
	return NewBool(ok)
}

func IsListFunc(args List) Object {
	_, ok := args.First().(*Cell)
	return NewBool(ok)
}

func IsSymbolFunc(args List) Object {
	_, ok := args.First().(*Symbol)
	return NewBool(ok)
}

////////////////////////////////////////////////////////////////////////////////
// List

func FirstFunc(args List) Object {
	return args.First().(List).First()
}

func NextFunc(args List) Object {
	return args.First().(List).Next()
}

func ConsFunc(args List) Object {
	x := args.First()
	l := args.Second().(List)
	l.Cons(x)
	return l
}

////////////////////////////////////////////////////////////////////////////////
// IO

func PrintFunc(args List) Object {
	for !args.Nil() {
		fmt.Print(args.First())
		args = args.Next()
		if !args.Nil() {
			fmt.Print(" ")
		}
	}
	return nil
}

func PrintlnFunc(args List) Object {
	PrintFunc(args)
	fmt.Println()
	return nil
}

func NewlineFunc(args List) Object {
	fmt.Println()
	return nil
}
