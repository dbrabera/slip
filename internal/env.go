package internal

type Enviroment struct {
	symbols map[string]Value
	parent  *Enviroment
}

func NewEnviroment() *Enviroment {
	return &Enviroment{symbols: make(map[string]Value)}
}

func NewChildEnviroment(parent *Enviroment) *Enviroment {
	return &Enviroment{symbols: make(map[string]Value), parent: parent}
}

func (e *Enviroment) Define(sym Symbol, val Value) {
	e.symbols[string(sym)] = val
}

func (e *Enviroment) Resolve(sym Symbol) Value {
	val := e.symbols[string(sym)]
	if val == nil && e.parent != nil {
		return e.parent.Resolve(sym)
	}
	return val
}
