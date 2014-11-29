package slip

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
	"bool?":   IsBool,
	"list?":   IsList,
	"neg?":    IsNeg,
	"nil?":    IsNil,
	"number?": IsNumber,
	"pos?":    IsPos,
	"string?": IsString,
	"symbol?": IsSymbol,
	"zero?":   IsZero,

	// List
	"cons":   Cons,
	"empty?": IsEmpty,
	"first":  First,
	"next":   Next,

	// IO
	"print": Print,
	//"printf": Printf,
	"println": Println,
	"newline": Newline,
	//"readline": Readline,
}

////////////////////////////////////////////////////////////////////////////////
// Arithmetic

func Add(nums ...Object) Object {
	res := nums[0].(*Number).Value

	for _, arg := range nums[1:] {
		res += arg.(*Number).Value
	}

	return NewNumber(res)
}

func Sub(nums ...Object) Object {
	if len(nums) == 1 {
		return NewNumber(-nums[0].(*Number).Value)
	}

	res := nums[0].(*Number).Value

	for _, arg := range nums[1:] {
		res -= arg.(*Number).Value
	}

	return NewNumber(res)
}

func Mul(nums ...Object) Object {
	res := nums[0].(*Number).Value

	for _, arg := range nums[1:] {
		res *= arg.(*Number).Value
	}

	return NewNumber(res)
}

func Div(nums ...Object) Object {
	if len(nums) == 1 {
		return NewNumber(1.0 / nums[0].(*Number).Value)
	}

	res := nums[0].(*Number).Value

	for _, arg := range nums[1:] {
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

func Inc(num Object) Object {
	return NewNumber(num.(*Number).Value + 1.0)
}

func Dec(num Object) Object {
	return NewNumber(num.(*Number).Value - 1.0)
}

////////////////////////////////////////////////////////////////////////////////
// Relational

func Gt(nums ...Object) Object {
	x := nums[0].(*Number).Value

	for _, arg := range nums[1:] {
		y := arg.(*Number).Value
		if x <= y {
			return NewBool(false)
		}
		x = y
	}

	return NewBool(true)
}

func Ge(nums ...Object) Object {
	x := nums[0].(*Number).Value

	for _, arg := range nums[1:] {
		y := arg.(*Number).Value
		if x < y {
			return NewBool(false)
		}
		x = y
	}

	return NewBool(true)
}

func Eq(objs ...Object) Object {
	x := objs[0]

	for _, arg := range objs[1:] {
		if !x.Equals(arg) {
			return NewBool(false)
		}
	}

	return NewBool(true)
}

func Ne(objs ...Object) Object {
	return Not(Eq(objs...))
}

func Le(nums ...Object) Object {
	x := nums[0].(*Number).Value

	for _, arg := range nums[1:] {
		y := arg.(*Number).Value
		if x > y {
			return NewBool(false)
		}
		x = y
	}

	return NewBool(true)
}

func Lt(nums ...Object) Object {
	x := nums[0].(*Number).Value

	for _, arg := range nums[1:] {
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
	if x == nil {
		return NewBool(true)
	}
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

func Cons(x Object, ls Object) Object {
	return ls.(List).Cons(x)
}

func IsEmpty(ls Object) Object {
	if ls == nil {
		return nil
	}
	return NewBool(ls.(List).IsEmpty())
}

func First(ls Object) Object {
	if ls == nil {
		return nil
	}
	return ls.(List).First()
}

func Next(ls Object) Object {
	if ls == nil {
		return nil
	}
	return ls.(List).Next()
}

////////////////////////////////////////////////////////////////////////////////
// IO

func Print(objs ...Object) Object {
	l := len(objs)
	for i, arg := range objs {
		fmt.Print(arg)
		if i < l-1 {
			fmt.Print(" ")
		}
	}
	return nil
}

func Println(objs ...Object) Object {
	Print(objs...)
	fmt.Println()
	return nil
}

func Newline() Object {
	fmt.Println()
	return nil
}
