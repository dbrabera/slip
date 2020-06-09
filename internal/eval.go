package internal

// Eval evaluates the Slip program represented as a string
// on the given environment, returning the value of the last expression.
func Eval(s string, env *Enviroment) (Value, error) {
	values, err := Parse(string(s))
	if err != nil {
		return nil, err
	}

	var out Value
	for _, value := range values {
		out = value.Eval(env)
	}

	return out, nil
}

type Enviroment struct {
	symbols map[string]Value
	parent  *Enviroment
}

func NewEnviroment() *Enviroment {
	env := &Enviroment{symbols: make(map[string]Value)}

	for name, fn := range CoreFuncs {
		env.Define(NewSymbol(name), fn)
	}

	return env
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
