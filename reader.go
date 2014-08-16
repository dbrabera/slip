package main

import (
	"strconv"
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
		return self.readInteger()
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

func (self *Reader) ignoreWhitespace() {
	for {
		r := self.next()
		if !unicode.IsSpace(r) {
			self.undo()
			break
		}
	}
}

func (self *Reader) readInteger() Object {
	start := self.offset
	for {
		r := self.next()
		if !unicode.IsDigit(r) {
			self.undo()
			break
		}
	}
	value, _ := strconv.ParseInt(self.source[start:self.offset], 0, 0)

	return NewInt(int(value))
}
