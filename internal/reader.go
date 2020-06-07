package internal

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/pkg/errors"
)

const eof = -1

type Reader struct {
	source string
	line   int
	offset int
}

func NewReader() *Reader {
	return &Reader{}
}

func (r *Reader) Init(source string) {
	r.source = source
	r.line = 1
	r.offset = 0
}

func (r *Reader) Read() Value {
	r.ignoreWhitespace()

	switch ch := r.next(); {
	case ch == eof:
		return nil
	case unicode.IsDigit(ch):
		r.undo()
		return r.readNumber()
	case ch == '+' || ch == '-':
		if unicode.IsDigit(r.peek()) {
			r.undo()
			return r.readNumber()
		}
		fallthrough
	case isIdentHead(ch):
		r.undo()
		return r.readIdent()
	case ch == '(':
		r.undo()
		return r.readList()
	case ch == '\'':
		return NewList(NewSymbol("quote"), r.Read())
	case ch == '"':
		r.undo()
		return r.readString()
	default:
		panic(errors.Errorf("unexpected character '%c'", ch))
	}
}

func (r *Reader) next() rune {
	var ch rune
	if r.offset < len(r.source) {
		ch = rune(r.source[r.offset])
	} else {
		ch = eof
	}
	r.offset++

	return ch
}

func (r *Reader) undo() {
	r.offset--
}

func (r *Reader) peek() rune {
	ch := r.next()
	r.undo()
	return ch
}

func (r *Reader) ignoreWhitespace() {
	for {
		ch := r.next()
		if !unicode.IsSpace(ch) {
			r.undo()
			break
		}
	}
}

func (r *Reader) readNumber() Value {
	start := r.offset
	point := false

	if ch := r.next(); ch != '+' && ch != '-' && !unicode.IsDigit(ch) {
		panic(errors.Errorf("unexpected character '%c'", ch))
	}

	for {
		ch := r.next()
		if !point && ch == '.' {
			point = true
			continue
		}

		if !unicode.IsDigit(ch) {
			r.undo()
			break
		}
	}

	value, err := strconv.ParseInt(r.source[start:r.offset], 10, 64)
	if err != nil {
		panic(errors.Wrap(err, "failed to parse float"))
	}
	return NewInt(value)
}

func (r *Reader) readIdent() Value {
	start := r.offset
	if ch := r.next(); !isIdentHead(ch) {
		panic(errors.Errorf("unexpected character '%c'", ch))
	}

	for {
		if !isIdentBody(r.next()) {
			r.undo()
			break
		}
	}

	switch lexeme := r.source[start:r.offset]; lexeme {
	case "true":
		return True
	case "false":
		return False
	default:
		return NewSymbol(lexeme)
	}
}

func isIdentHead(r rune) bool {
	return unicode.IsLetter(r) || strings.IndexRune("+-*/_!?><=$", r) >= 0
}

func isIdentBody(r rune) bool {
	return isIdentHead(r) || unicode.IsDigit(r)
}

func (r *Reader) readString() Value {
	if ch := r.next(); ch != '"' {
		panic(errors.Errorf("unexpected character '%c'", ch))
	}

	start := r.offset
	for {
		switch ch := r.next(); ch {
		case '"':
			return NewString(r.source[start : r.offset-1])
		case '\n', eof:
			panic(errors.Errorf("unexpected character '%c'", ch))
		default:
			continue
		}
	}
}

func (r *Reader) readList() Value {
	if ch := r.next(); ch != '(' {
		panic(errors.Errorf("unexpected character '%c'", ch))
	}

	l := NewList()

	for {
		r.ignoreWhitespace()

		ch := r.next()

		if ch == eof {
			panic(errors.Errorf("unexpected character '%c'", ch))
		} else if ch == ')' {
			break
		}

		r.undo()

		l = append(l, r.Read())
	}

	return l
}
