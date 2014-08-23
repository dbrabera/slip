package main

type ProcFn func(List) Object

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

func FirstProc(args List) Object {
	return args.First().(List).First()
}

func NextProc(args List) Object {
	return args.First().(List).Next()
}
