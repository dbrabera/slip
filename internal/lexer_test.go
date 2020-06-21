package internal

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	cases := []struct {
		s        string
		expected []TokenKind
	}{
		{"", []TokenKind{}},
		{" \n\t", []TokenKind{}},
		{"( )", []TokenKind{TLeftParen, TRightParen}},
		{`"" "foo"`, []TokenKind{TString, TString}},
		{"1 12 123", []TokenKind{TInt, TInt, TInt}},
		{"- -1 -a", []TokenKind{TSymbol, TInt, TSymbol}},
		{"true false foo", []TokenKind{TBool, TBool, TSymbol}},
		{"!@$%^&*-_+=|~:<>.?\\/,", []TokenKind{TSymbol}},
		{"; foo", []TokenKind{}},
		{"1 ; foo\n 2", []TokenKind{TInt, TInt}},
	}

	for i, c := range cases {
		tokens, err := Tokenize(c.s)
		if err != nil {
			t.Fatalf("%d: err: %v", i, err)
		}

		found := make([]TokenKind, len(tokens))
		for i, t := range tokens {
			found[i] = t.Kind
		}

		if !reflect.DeepEqual(c.expected, found) {
			t.Errorf("%d: expected = %v, found %v", i, c.expected, found)
		}
	}
}
