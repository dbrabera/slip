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

func (e *Enviroment) Define(sym *Symbol, obj Object) {
	e.symbols[sym.Value] = obj
}

func (e *Enviroment) Resolve(sym *Symbol) Object {
	obj := e.symbols[sym.Value]
	if obj == nil && e.parent != nil {
		return e.parent.Resolve(sym)
	}
	return obj
}
