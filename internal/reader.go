package internal

import (
	"strconv"
	"strings"
	"unicode"
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

func (self *Reader) Init(source string) {
	self.source = source
	self.line = 1
	self.offset = 0
}

func (self *Reader) Read() Object {
	self.ignoreWhitespace()

	switch r := self.next(); {
	case r == eof:
		return nil
	case unicode.IsDigit(r):
		self.undo()
		return self.readNumber()
	case r == '+' || r == '-':
		if unicode.IsDigit(self.peek()) {
			self.undo()
			return self.readNumber()
		}
		fallthrough
	case isIdentHead(r):
		self.undo()
		return self.readIdent()
	case r == '(':
		self.undo()
		return self.readList()
	case r == '\'':
		return NewList(NewSymbol("quote"), self.Read())
	case r == '"':
		self.undo()
		return self.readString()
	default:
		panic("Unexpected character")
	}
}

func (self *Reader) next() rune {
	var r rune
	if self.offset < len(self.source) {
		r = rune(self.source[self.offset])
	} else {
		r = eof
	}
	self.offset += 1

	return r
}

func (self *Reader) undo() {
	self.offset -= 1
}

func (self *Reader) peek() rune {
	r := self.next()
	self.undo()
	return r
}

func (self *Reader) ignoreWhitespace() {
	for {
		r := self.next()
		if !unicode.IsSpace(r) {
			self.undo()
			break
		}
	}
}

func (self *Reader) readNumber() Object {
	start := self.offset
	point := false

	if r := self.next(); r != '+' && r != '-' && !unicode.IsDigit(r) {
		panic("Unexpected character")
	}

	for {
		r := self.next()
		if !point && r == '.' {
			point = true
			continue
		}

		if !unicode.IsDigit(r) {
			self.undo()
			break
		}
	}

	value, _ := strconv.ParseFloat(self.source[start:self.offset], 64)
	return NewNumber(value)
}

func (self *Reader) readIdent() Object {
	start := self.offset
	if !isIdentHead(self.next()) {
		panic("Unexpected character")
	}

	for {
		if !isIdentBody(self.next()) {
			self.undo()
			break
		}
	}

	switch lexeme := self.source[start:self.offset]; lexeme {
	case "true":
		return NewBool(true)
	case "false":
		return NewBool(false)
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

func (self *Reader) readString() Object {
	if self.next() != '"' {
		panic("Unexpected character")
	}

	start := self.offset
	for {
		switch self.next() {
		case '"':
			return NewString(self.source[start : self.offset-1])
		case '\n', eof:
			panic("Unexpected character")
		default:
			continue
		}
	}
}

func (self *Reader) readList() Object {
	if r := self.next(); r != '(' {
		panic("Unexpected character")
	}

	var front *Cell = &Cell{}
	var curr *Cell = nil

	for {
		self.ignoreWhitespace()

		r := self.next()

		if r == eof {
			panic("EOF while reading")
		} else if r == ')' {
			break
		}

		self.undo()

		obj := self.Read()
		if curr == nil {
			front.Value = obj
			curr = front
		} else {
			curr.More = NewCell(obj, nil)
			curr = curr.More
		}
	}

	return front
}
