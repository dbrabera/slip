package main

import "strconv"

type Object interface {
	Eval() Object
	String() string
}

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
