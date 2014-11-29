# Special forms


## And
Usage: (and exprs\*)

Evaluates the exprs in order. If an expr returns a logical false, stops evaluating and
returns that value, else returns the value of the last expr.

```clojure
(and true) ; => true
(and true "hello") ; => "hello"
(and false "hello") ; => false
(and true "hello" nil) ; => nil
```

## Def
Usage: (def sym init)

Defines a variable in the current enviroment, evaluates init and binds its value
to the defined variable.

```clojure
(def a "hello")
a ; => "hello"
```

## Defn
Usage: (defn name (params\*) exprs\*)

Same as (def name (fn (params\*) exprs\*))

```clojure
(defn sum (x y) (+ x y))
(sum 1 2) ; => 3
```

## Do
Usage: (do exprs\*)

Evaluates the exprs in order and returns the value of the last expr.

```clojure
(do "hello" "world") ; => "world"
```

## Fn
Usage: (fn (params\*) exprs\*)

Returns a function.

```clojure
((fn (x y) (+ x y)) 1 2) ; => 3
```

## If
Usage: (if test then else?)

Evaluates test. If test is a logical true, evaluates and returns then, otherwise,
evaluates and returns else. If else is not provided returns nil.

```clojure
(if true "hello") ; => "hello"
(if false "hello") ; => nil
(if false "hello" "world") ; => "hello"
```

## Let
Usage: (let ((sym init)\*) exprs\*)

Evaluates the exprs in order in a new enviroment in wich the syms are binded to the
init values. Returns the value of the last expr.

```clojure
(let ((x 1) (y 2)) (+ x y)) ; => 3
```

## Or
Usage: (or exprs\*)

Evaluates the exprs in order. If a expr returns a logical true, stops evaluating and
returns that value, else returns the value of the last expr.

```clojure
(or true) ; => true
(or true "hello") ; => true
(or false "hello") ; => "hello"
(or false false nil) ; => nil
```

## Quote
Usage: (quote expr)

Returns the unevaluated expr.

```clojure
(quote (+ 1 2)) ; => (+ 1 2)
'(+ 1 2) ; => (+ 1 2)
```
