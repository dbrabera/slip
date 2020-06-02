package internal

type Enviroment struct {
	symbols map[string]Object
	parent  *Enviroment
}

func NewEnviroment() *Enviroment {
	return &Enviroment{symbols: make(map[string]Object)}
}

func NewChildEnviroment(parent *Enviroment) *Enviroment {
	return &Enviroment{symbols: make(map[string]Object), parent: parent}
}

func (self *Enviroment) Define(sym *Symbol, obj Object) {
	self.symbols[sym.Value] = obj
}

func (self *Enviroment) Resolve(sym *Symbol) Object {
	obj := self.symbols[sym.Value]
	if obj == nil && self.parent != nil {
		return self.parent.Resolve(sym)
	}
	return obj
}
