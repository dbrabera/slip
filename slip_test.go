package main

import "testing"

func TestSlip(t *testing.T) {
	cases := []struct {
		Src, Res string
	}{
		// Special forms
		{"(quote (1 2 3))", "(1 2 3)"},
		{"'(1 2 3)", "(1 2 3)"},

		// Core functions

		// Arithmetic
		{"(+ 1)", "1"},
		{"(+ 1 2)", "3"},
		{"(+ 1.1 2.0)", "3.1"},
		{"(+ 1.1 1 2)", "4.1"},

		{"(- 1)", "-1"},
		{"(- 1 1)", "0"},
		{"(- 1.1 1.0)", "0.1"},
		{"(- 1.1 1 1)", "-0.9"},

		{"(* 2)", "2"},
		{"(* 2 2)", "4"},
		{"(* 2.5 2.5)", "6.25"},
		{"(* 2 2.5 3)", "15"},

		{"(/ 2)", "0.5"},
		{"(/ 4 2)", "2"},
		{"(/ 4.0 2.0)", "2"},
		{"(/ 4.0 2.0 4)", "0.5"},

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

		{"(= 1 1)", "true"},
		{"(= 1 2)", "false"},
		{"(= 1.0 1.0)", "true"},
		{"(= 1.0 2.0)", "false"},
		{"(= 1 1.0)", "true"},
		{"(= 1 2.0)", "false"},
		{"(= true true)", "true"},
		{"(= true false)", "false"},
		{"(= \"abc\" \"abc\")", "true"},
		{"(= \"abc\" \"xyz\")", "false"},
		{"(= '(1 1.0 true \"abc\") '(1 1.0 true \"abc\"))", "true"},
		{"(= '(1 1.0 true \"abc\") '(1 1.0 false \"abc\"))", "false"},
		{"(= 1 1 1 1)", "true"},

		{"(!= true false)", "true"},
		{"(!= true true)", "false"},

		// Logic
		{"(not false)", "true"},
		{"(not true)", "false"},
		{"(not 1)", "false"},

		// Test
		{"(zero? 0)", "true"},
		{"(zero? 0.0)", "true"},
		{"(zero? 1)", "false"},
		{"(zero? 1.0)", "false"},

		{"(pos? 1)", "true"},
		{"(pos? 1.0)", "true"},
		{"(pos? -1)", "false"},
		{"(pos? -1.0)", "false"},

		{"(neg? -1)", "true"},
		{"(neg? -1.0)", "true"},
		{"(neg? 1)", "false"},
		{"(neg? 1.0)", "false"},

		{"(even? 2)", "true"},
		{"(even? 1)", "false"},

		{"(odd? 1)", "true"},
		{"(odd? 2)", "false"},

		{"(int? 1)", "true"},
		{"(int? 1.0)", "false"},

		{"(double? 1.0)", "true"},
		{"(double? 1)", "false"},

		{"(bool? true)", "true"},
		{"(bool? 1)", "false"},

		{"(string? \"str\")", "true"},
		{"(string? 1)", "false"},

		{"(symbol? 'a)", "true"},
		{"(symbol? 1)", "false"},

		{"(list? '(1 2 3))", "true"},
		{"(list? 1)", "false"},

		// List
		{"(first '(1 2 3))", "1"},
		{"(next '(1 2 3))", "(2 3)"},
		{"(cons 1 '(2 3))", "(1 2 3)"},
	}

	for _, c := range cases {
		slip := NewSlip()
		res := slip.Exec(c.Src).String()
		if c.Res != res {
			t.Errorf("\nInput: %s\nActual: %s\nExpected: %s", c.Src, res, c.Res)
		}
	}
}
