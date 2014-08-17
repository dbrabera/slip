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
// Double

type Double struct {
	Value float64
}

func NewDouble(value float64) *Double {
	return &Double{Value: value}
}

func (self *Double) Eval() Object {
	return self
}

func (self *Double) String() string {
	return strconv.FormatFloat(self.Value, 'f', -1, 64)
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
