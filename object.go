package main

import (
	"fmt"
	"strconv"
)

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

var (
	t Bool = Bool{Value: true}
	f Bool = Bool{Value: false}
)

func NewBool(value bool) Object {
	if value {
		return &t
	}
	return &f
}

func (self *Bool) Eval() Object {
	return self
}

func (self *Bool) String() string {
	return strconv.FormatBool(self.Value)
}

////////////////////////////////////////////////////////////////////////////////
// String

type String struct {
	Value string
}

func NewString(value string) Object {
	return &String{Value: value}
}

func (self *String) Eval() Object {
	return self
}

func (self *String) String() string {
	return fmt.Sprintf("\"%s\"", self.Value)
}
