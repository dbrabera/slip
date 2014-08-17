package main

import "strconv"

type Object interface {
	Eval() Object
	String() string
}

////////////////////////////////////////////////////////////////////////////////
// Integer

type Int struct {
	Value int
}

func NewInt(value int) *Int {
	return &Int{Value: value}
}

func (self *Int) Eval() Object {
	return self
}

func (self *Int) String() string {
	return strconv.Itoa(self.Value)
}

////////////////////////////////////////////////////////////////////////////////
// Boolean

type Bool struct {
	Value bool
}

func NewBool(value bool) Object {
	return &Bool{Value: value}
}

func (self *Bool) Eval() Object {
	return self
}

func (self *Bool) String() string {
	return strconv.FormatBool(self.Value)
}
