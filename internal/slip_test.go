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
		{"(+ 1 2)", "3"},
		{"(+ 1 2 3)", "6"},

		{"(- 1)", "-1"},
		{"(- 1 1)", "0"},
		{"(- 1 1 2)", "-2"},

		{"(* 4)", "4"},
		{"(* 4 2)", "8"},
		{"(* 4 2 4)", "32"},

		{"(/ 4)", "0"},
		{"(/ 4 2)", "2"},
		{"(/ 4 2 4)", "0"},

		{"(mod 5 2)", "1"},

		{"(inc 1)", "2"},
		{"(dec 1)", "0"},

		// Relational
		{"(> 3)", "true"},
		{"(> 3 2)", "true"},
		{"(> 3 2 1)", "true"},
		{"(> 3 1 2)", "false"},

		{"(>= 3)", "true"},
		{"(>= 3 3)", "true"},
		{"(>= 3 3 1)", "true"},
		{"(>= 3 1 2)", "false"},

		{"(<= 1)", "true"},
		{"(<= 1 1)", "true"},
		{"(<= 1 1 3)", "true"},
		{"(<= 1 3 2)", "false"},

		{"(< 1)", "true"},
		{"(< 1 2)", "true"},
		{"(< 1 2 3)", "true"},
		{"(< 1 3 2)", "false"},

		{"(= 1)", "true"},
		{"(= 1 1)", "true"},
		{"(= true true)", "true"},
		{"(= true false)", "false"},
		{"(= \"abc\" \"abc\")", "true"},
		{"(= \"abc\" \"xyz\")", "false"},
		{"(= '(1 1 true \"abc\") '(1 1 true \"abc\"))", "true"},
		{"(= '(1 1 true \"abc\") '(1 1 false \"abc\"))", "false"},
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

		{"(int? 1)", "true"},
		{"(int? \"str\")", "false"},

		{"(bool? true)", "true"},
		{"(bool? 1)", "false"},

		{"(string? \"str\")", "true"},
		{"(string? 1)", "false"},

		{"(symbol? 'a)", "true"},
		{"(symbol? 1)", "false"},

		{"(list? '(1 2 3))", "true"},
		{"(list? 1)", "false"},
	}

	for _, c := range cases {
		slip := NewSlip()
		res := fmt.Sprint(slip.Exec(c.Src))
		if c.Res != res {
			t.Errorf("\nInput: %s\nActual: %s\nExpected: %s", c.Src, res, c.Res)
		}
	}
}
