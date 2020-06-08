package internal

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	cases := []string{
		`1 true false "foo"`,
		"(foo 1 2 3)",
	}

	for i, c := range cases {
		values, err := Parse(c)
		if err != nil {
			t.Fatalf("%d: err: %v", i, err)
		}

		s := []string{}
		for _, v := range values {
			s = append(s, v.String())
		}

		found := strings.Join(s, " ")
		if c != found {
			t.Errorf("%d: expected = %v, found %v", i, c, found)
		}
	}
}
