package internal

import (
	"fmt"
	"testing"
)

func TestSlip_Exec(t *testing.T) {
	cases := []struct {
		Src, Res string
	}{
		// Special forms
		{"(and true)", "true"},
		{"(and true \"hello\")", "\"hello\""},
		{"(and false \"hello\")", "false"},
		{"(and true \"hello\" nil)", "<nil>"},

		{"(do (def a \"hello\") a)", "\"hello\""},

		{"(do (defn sum (x y) (+ x y)) (sum 1 2))", "3"},

		{"(do \"hello\" \"world\")", "\"world\""},

		{"((fn (x y) (+ x y)) 1 2)", "3"},

		{"(if true \"hello\")", "\"hello\""},
		{"(if false \"hello\")", "<nil>"},
		{"(if false \"hello\" \"world\")", "\"world\""},

		{"(let ((x 1) (y 2)) (+ x y))", "3"},

		{"(or true)", "true"},
		{"(or true \"hello\")", "true"},
		{"(or false \"hello\")", "\"hello\""},
		{"(or false false nil)", "<nil>"},

		{"(quote (+ 1 2))", "(+ 1 2)"},
		{"'(+ 1 2)", "(+ 1 2)"},

		// Core functions

		// Arithmetic
		{"(+ 1)", "1"},
		{"(+ 1 2.2)", "3.2"},
		{"(+ 1 2.2 3)", "6.2"},

		{"(- 1)", "-1"},
		{"(- 1 1.123)", "-0.123"},
		{"(- 1 1.123 3)", "-3.123"},

		{"(* 4)", "4"},
		{"(* 4 2.0)", "8"},
		{"(* 4 2.0 4)", "32"},

		{"(/ 4)", "0.25"},
		{"(/ 4 2.0)", "2"},
		{"(/ 4 2.0 4)", "0.5"},

		{"(mod 5 2)", "1"},
		{"(rem 5 2)", "1"},

		{"(inc 1)", "2"},
		{"(dec 1)", "0"},

		// Relational
		{"(> 3)", "true"},
		{"(> 3 2.4)", "true"},
		{"(> 3 2.4 1)", "true"},
		{"(> 3 1 2.4)", "false"},

		{"(>= 3)", "true"},
		{"(>= 3 3.0)", "true"},
		{"(>= 3 3.0 1)", "true"},
		{"(>= 3 1 2.0)", "false"},

		{"(<= 1)", "true"},
		{"(<= 1 1.0)", "true"},
		{"(<= 1 1.0 3)", "true"},
		{"(<= 1 3 2.0)", "false"},

		{"(< 1)", "true"},
		{"(< 1 2.0)", "true"},
		{"(< 1 2.0 3)", "true"},
		{"(< 1 3 2.0)", "false"},

		{"(= 1)", "true"},
		{"(= 1 1.0)", "true"},
		{"(= true true)", "true"},
		{"(= true false)", "false"},
		{"(= \"abc\" \"abc\")", "true"},
		{"(= \"abc\" \"xyz\")", "false"},
		{"(= '(1 1.0 true \"abc\") '(1 1.0 true \"abc\"))", "true"},
		{"(= '(1 1.0 true \"abc\") '(1 1.0 false \"abc\"))", "false"},
		{"(= 1 1 1 1)", "true"},

		{"(!= 1 2)", "true"},
		{"(!= 1 1)", "false"},

		// Logic
		{"(not false)", "true"},
		{"(not true)", "false"},
		{"(not nil)", "true"},
		{"(not \"str\")", "false"},

		// Test
		{"(zero? 0)", "true"},
		{"(zero? 1)", "false"},

		{"(pos? 1)", "true"},
		{"(pos? -1)", "false"},

		{"(neg? -1)", "true"},
		{"(neg? 1)", "false"},

		{"(number? 1)", "true"},
		{"(number? \"str\")", "false"},

		{"(bool? true)", "true"},
		{"(bool? 1)", "false"},

		{"(string? \"str\")", "true"},
		{"(string? 1)", "false"},

		{"(symbol? 'a)", "true"},
		{"(symbol? 1)", "false"},

		{"(list? '(1 2 3))", "true"},
		{"(list? 1)", "false"},

		// List
		{"(cons 1 '(2 3))", "(1 2 3)"},

		{"(empty? '(1 2 3))", "false"},
		{"(empty? '())", "true"},

		{"(next '(1 2 3))", "(2 3)"},
		{"(next '(1))", "<nil>"},
		{"(next '())", "<nil>"},
		{"(next nil)", "<nil>"},

		{"(first '(1 2 3))", "1"},
		{"(first '())", "<nil>"},
		{"(first nil)", "<nil>"},
	}

	for _, c := range cases {
		slip := NewSlip()
		res := fmt.Sprint(slip.Exec(c.Src))
		if c.Res != res {
			t.Errorf("\nInput: %s\nActual: %s\nExpected: %s", c.Src, res, c.Res)
		}
	}
}
