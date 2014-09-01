#Slip

A toy Lisp dialect.

## Usage

```
Usage: slip [options] [script [args]]

Options:

-h, --help       Show this help
-v, --version    Show version information
```

## Special forms

### And
Usage: (and exprs\*)

Evaluates the exprs in order. If an expr returns a logical false, stops evaluating and
returns that value, else returns the value of the last expr.

```clojure
(and true) ; => true
(and true "hello") ; => "hello"
(and false "hello") ; => false
(and true "hello" nil) ; => nil
```

### Def
Usage: (def sym init)

Defines a variable in the current enviroment, evaluates init and binds its value
to the defined variable.

```clojure
(def a "hello")
a ; => "hello"
```

### Defn
Usage: (defn name (params\*) exprs\*)

Same as (def name (fn (params\*) exprs\*))

```clojure
(defn sum (x y) (+ x y))
(sum 1 2) ; => 3
```

### Do
Usage: (do exprs\*)

Evaluates the exprs in order and returns the value of the last expr.

```clojure
(do "hello" "world") ; => "world"
```

### Fn
Usage: (fn (params\*) exprs\*)

Returns a function.

```clojure
((fn (x y) (+ x y)) 1 2) ; => 3
```

### If
Usage: (if test then else?)

Evaluates test. If test is a logical true, evaluates and returns then, otherwise,
evaluates and returns else. If else is not provided returns nil.

```clojure
(if true "hello") ; => "hello"
(if false "hello") ; => nil
(if false "hello" "world") ; => "hello"
```

### Let
Usage: (let ((sym init)\*) exprs\*)

Evaluates the exprs in order in a new enviroment in wich the syms are binded to the
init values. Returns the value of the last expr.

```clojure
(let ((x 1) (y 2)) (+ x y)) ; => 3
```

### Or
Usage: (or exprs\*)

Evaluates the exprs in order. If a expr returns a logical true, stops evaluating and
returns that value, else returns the value of the last expr.

```clojure
(or true) ; => true
(or true "hello") ; => true
(or false "hello") ; => "hello"
(or false false nil) ; => nil
```

### Quote
Usage: (quote expr)

Returns the unevaluated expr.

```clojure
(quote (+ 1 2)) ; => (+ 1 2)
'(+ 1 2) ; => (+ 1 2)
```


## Core library

### +
Usage: (+ nums+)

Returns the sum of the nums.

```clojure
(+ 1) ; => 1
(+ 1 2.2) ; => 3.2
(+ 1 2.2 3) ; => 6.2
```

### -
Usage: (- nums+)

If only a number is supplied, returns its negation, else returns the
substraction of nums.

```clojure
(- 1) ; => -1
(- 1 2.2) ; => -1.2
(- 1 2.2 3) ; => -4.3
```

### *
Usage: (* nums+)

If only a number is supplied, returns 1*num, else returns the product of nums.

```clojure
(* 4) ; => 4
(* 4 2.0) ; => 8
(* 4 2.0 4) ; => 32
```

### /
Usage: (/ nums+)

If only a number is supplied, returns 1/num, else returns the division of nums.

```clojure
(/ 4) ; => 0.25
(/ 4 2.0) ; => 2
(/ 4 2.0 4) ; => 0.5
```

### rem
Usage: (rem num div)

Returns the remainder of dividing num and div.

```clojure
(rem 5 1) ; => 1
```

### mod
Usage: (mod num div)

Returns the modulus of num and div.

```clojure
(mod 5 1) ; => 1
```

### inc
Usage: (inc num)

Returns a number one greater than num.

```clojure
(inc 1) ; => 2
```

### dec
Usage: (dec num)

Returns a number one less than num.

```clojure
(dec 1) ; => 0
```

### <
Usage: (< nums+)

Returns true if the nums are in monotonically increasing order, false otherwise.

```clojure
(< 1) ; => true
(< 1 2.2) ; => true
(< 1 2.2 3) ; => true
```

### <=
Usage: (<= nums+)

Returns true if the nums are in monotonically non-decreasing order, false otherwise.

```clojure
(<= 1) ; => true
(<= 1 2.2) ; => true
(<= 1 2.2 2.2) ; => true
```

### =
Usage: (= objs+)

Returns true if the objs are equals, false otherwise

```clojure
(= 1) ; => true
(= 1 1.0) ; => true
(= 1 1.0 1) ; => true
```

### !=
Usage: (!= objs+)

Returns true if the objs are not equals, false otherwise

```clojure
(!= 1) ; => false
(!= 1 2) ; => true
(!= 1 2 1) ; => true
```

### >=
Usage: (>= nums+)

Returns true if the nums are in monotonically non-increasing order, false otherwise.

```clojure
(> 3) ; => true
(> 3 2.2) ; => true
(> 3 2.2 2.2) ; => true
```

### >
Usage: (> nums+)

Returns true if the nums are in monotonically decreasing order, false otherwise.

```clojure
(> 3) ; => true
(> 3 2.2) ; => true
(> 3 2.2 1) ; => true
```

### not
Usage: (not x)

Returns true if x is a logical false, false otherwise.

```clojure
(not false) ; => true
(not nil) ; => true
(not "str") ; => false
```

### bool?
Usage: (bool? x)

Returns true if x is bool, false otherwise.

```clojure
(bool? true) ; => true
```

### list?
Usage: (list? x)

Returns true if x is a list, false otherwise.

```clojure
(list? '(1 2 3)) ; => true
```

### neg?
Usage: (neg? x)

Returns true if x is less than zero, false otherwise.

```clojure
(neg? -1) ; => true
```

### nil?
Usage: (nil? x)

Returns true if x is a nil, false otherwise.

```clojure
(nil? nil) ; => true
```

### number?
Usage: (number? x)

Returns true if x is a number, false otherwise.

```clojure
(number? 1) ; => true
```

### pos?
Usage: (pos? x)

Returns true if x is bigger than zero, false otherwise.

```clojure
(pos? 1) ; => true
```

### string?
Usage: (string? x)

Returns true if x is a string, false otherwise.

```clojure
(string? "str") ; => true
```

### symbol?
Usage: (symbol? x)

Returns true if x is a symbol, false otherwise.

```clojure
(symbol? 'a) ; => true
```

### zero?
Usage: (zero? x)

Returns true if x is zero, false otherwise.

```clojure
(zero? 0) ; => true
```

### cons
Usage: (cons x ls)

Returns a new list where x is the first element and ls is the rest.

```clojure
(cons 1 '(2 3)) ; => (1 2 3)
```

### empty?
Usage: (empty? ls)
Returns true if ls is empty, false otherwise.

```clojure
(empty? '()) ; => true
```

### first
Usage: (first ls)

Returns the first element in ls. If ls is empty or nil, returns nil.

```clojure
(first '(1 2 3)) ; => 1
(first '()) ; => nil
(first nil) ; => nil
```

### next
Usage: (next ls)

Returns a list with the elements after the first in ls. If there are no more
elements or ls is empty or nil, returns nil.

```clojure
(next '(1 2 3)) ; => (2 3)
(next '(1) ; => nil
(next '()) ; => nil
(next nil) ; => nil
```

### print
Usage: (print objs*)

Print the objs to stdout.

```clojure
(print 1 true "str" '(1 2 3))
```

### println
Usage: (println objs*)

Print the objs to stdout followed by a newline

```clojure
(println 1 true "str" '(1 2 3))
```

### newline
Usage: (newline)

Print a newline to stdout

```clojure
(newline)
```

##License

MIT