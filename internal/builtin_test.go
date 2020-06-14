package internal

import "testing"

type funcTestCase struct {
	args     []Value
	expected Value
}

func testFunc(t *testing.T, fn NativeFunc, cases []funcTestCase) {
	for i, c := range cases {
		found := fn(c.args...)
		if found != c.expected {
			t.Errorf("%d: expected = %v, found %v", i, c.expected, found)
		}
	}
}

func TestAdd(t *testing.T) {
	testFunc(t, add, []funcTestCase{
		{[]Value{}, Int(0)},
		{[]Value{Int(1)}, Int(1)},
		{[]Value{Int(1), Int(2)}, Int(3)},
	})
}

func TestSub(t *testing.T) {
	testFunc(t, sub, []funcTestCase{
		{[]Value{}, Int(0)},
		{[]Value{Int(1)}, Int(-1)},
		{[]Value{Int(1), Int(2)}, Int(-1)},
	})
}

func TestDiv(t *testing.T) {
	testFunc(t, div, []funcTestCase{
		{[]Value{}, Int(1)},
		{[]Value{Int(4)}, Int(0)},
		{[]Value{Int(4), Int(2)}, Int(2)},
	})
}

func TestMod(t *testing.T) {
	testFunc(t, mod, []funcTestCase{
		{[]Value{Int(5), Int(2)}, Int(1)},
	})
}

func TestInc(t *testing.T) {
	testFunc(t, inc, []funcTestCase{
		{[]Value{Int(1)}, Int(2)},
	})
}

func TestDec(t *testing.T) {
	testFunc(t, dec, []funcTestCase{
		{[]Value{Int(1)}, Int(0)},
	})
}

func TestGt(t *testing.T) {
	testFunc(t, gt, []funcTestCase{
		{[]Value{}, True},
		{[]Value{Int(3)}, True},
		{[]Value{Int(3), Int(2)}, True},
		{[]Value{Int(3), Int(2), Int(2)}, False},
		{[]Value{Int(1), Int(2)}, False},
	})
}

func TestGe(t *testing.T) {
	testFunc(t, ge, []funcTestCase{
		{[]Value{}, True},
		{[]Value{Int(3)}, True},
		{[]Value{Int(3), Int(2)}, True},
		{[]Value{Int(3), Int(2), Int(2)}, True},
		{[]Value{Int(1), Int(2)}, False},
	})
}

func TestEq(t *testing.T) {
	testFunc(t, eq, []funcTestCase{
		{[]Value{}, True},
		{[]Value{Int(1)}, True},
		{[]Value{Int(1), Int(1)}, True},
		{[]Value{Int(1), Int(1)}, True},
		{[]Value{Int(1), Int(2)}, False},
	})
}

func TestNe(t *testing.T) {
	testFunc(t, ne, []funcTestCase{
		{[]Value{}, False},
		{[]Value{Int(1)}, False},
		{[]Value{Int(1), Int(1)}, False},
		{[]Value{Int(1), Int(1)}, False},
		{[]Value{Int(1), Int(2)}, True},
	})
}

func TestLe(t *testing.T) {
	testFunc(t, le, []funcTestCase{
		{[]Value{}, True},
		{[]Value{Int(1)}, True},
		{[]Value{Int(1), Int(2)}, True},
		{[]Value{Int(1), Int(2), Int(2)}, True},
		{[]Value{Int(2), Int(1)}, False},
	})
}

func TestLt(t *testing.T) {
	testFunc(t, lt, []funcTestCase{
		{[]Value{}, True},
		{[]Value{Int(1)}, True},
		{[]Value{Int(1), Int(2)}, True},
		{[]Value{Int(1), Int(2), Int(2)}, False},
		{[]Value{Int(2), Int(1)}, False},
	})
}

func TestNot(t *testing.T) {
	testFunc(t, not, []funcTestCase{
		{[]Value{True}, False},
		{[]Value{False}, True},
	})
}

func TestIsNil(t *testing.T) {
	testFunc(t, isNil, []funcTestCase{
		{[]Value{nil}, True},
		{[]Value{False}, False},
	})
}

func TestIsZero(t *testing.T) {
	testFunc(t, isZero, []funcTestCase{
		{[]Value{Int(0)}, True},
		{[]Value{Int(1)}, False},
	})
}
func TestIsPos(t *testing.T) {
	testFunc(t, isPos, []funcTestCase{
		{[]Value{Int(1)}, True},
		{[]Value{Int(-1)}, False},
	})
}

func TestIsNeg(t *testing.T) {
	testFunc(t, isNeg, []funcTestCase{
		{[]Value{Int(-1)}, True},
		{[]Value{Int(1)}, False},
	})
}

func TestIsInt(t *testing.T) {
	testFunc(t, isInt, []funcTestCase{
		{[]Value{Int(1)}, True},
		{[]Value{String("s")}, False},
	})
}

func TestIsBool(t *testing.T) {
	testFunc(t, isBool, []funcTestCase{
		{[]Value{True}, True},
		{[]Value{String("s")}, False},
	})
}

func TestIsString(t *testing.T) {
	testFunc(t, isString, []funcTestCase{
		{[]Value{String("s")}, True},
		{[]Value{Int(1)}, False},
	})
}

func TestIsList(t *testing.T) {
	testFunc(t, isList, []funcTestCase{
		{[]Value{NewList()}, True},
		{[]Value{Int(1)}, False},
	})
}

func TestIsSymbol(t *testing.T) {
	testFunc(t, isSymbol, []funcTestCase{
		{[]Value{Symbol("foo")}, True},
		{[]Value{Int(1)}, False},
	})
}

func TestIsEmpty(t *testing.T) {
	testFunc(t, isEmpty, []funcTestCase{
		{[]Value{NewList()}, True},
		{[]Value{NewList(Int(1))}, False},
	})
}
