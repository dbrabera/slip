;; An annotated tour of all language features

;; Syntax
;;;;;;;;;;;;

;; Integers are signed 64-bit integer values.

0 ; => 0
12 ; => 12
-7 ; => -7

;; Strings are surronded by double quotes and can contain any Unicode character.

"Hello, 世界" ; => "Hello, 世界"

;; Booleans are predeclared constants.

true ; => true
false ; => false

;; Symbol are names associated with a value.

println ; => <function>
+ ; => <function>

;;; Special Forms
;;;;;;;;;;;;;;;;;;;;

;; and evaluates the expressions in order until any returns a false value, otherwise
;; it returns the result of the last one.

(and true false) ; => false

;; or evaluates the expressions in order until any returns a true value, otherwise it
;; returns the result last one.

(or true false) ; => true

;; if is a branching construct that evaluates one or another form
;; depending on the evaluation of a condition.

(if true "hello") ; => "hello"
(if false "hello") ; => nil
(if false "hello" "world") ; => "world"

;; fn creates a new function.

(fn (x y) (+ x y)) ; => <function>

;; def binds a value to a symbol.

(def x 1) ; => nil
x ; => 1

;; defn is a shorter version for creating a function and binding it to a symbol.

(defn sum (x y) (+ x y)) ; => nil
(sum 1 2) ; => 3

;; do creates a new lexical scope and evaluates a series of expressions in the new
;; scope, returning the result of the last one.

(do 1 2 3 4) ; => 4

;; let defines temporary bindings by creating a new lexical scope, binding the values 
;; to the symbols and evaluating the expressions in the new scope, returning the 
;; result of the last one.

(let ((x 1) (y 2))
  (+ x y)) ; => 3

;; quote returns the unevaluated expression.

(quote (+ 1 2)) ; => (+ 1 2)

;; Built-in Functions
;;;;;;;;;;;;;;;;;;;;;;;;

;; Use the usual arithmetic operators to perform basic math

(+ 1 1) ; => 2
(- 2 1) ; => 1
(* 1 2) ; => 2
(/ 4 2) ; => 2
(mod 4 2) ; => 0

;; all the operators are varidic and can have more than too values

(+ 1 2 3) ; => 6

;; Use the usual comparison operators to compare values

(> 1 2) ; => false
(>= 1 1) ; => true
(= 1 2) ; => false
(!= 1 2) ; => true
(<= 1 1) ; => true
(< 1 2) ; => true

;; all comparison operators are varidic too

(< 1 2 3 4)

;; Use the not operator to negate boolean values

(not true) ; => false

;; Check whether the values have a certain type

(bool? true) ; => true
(list? (quote (1, 2, 3))) ; => true
(nil? nil) ; => true
(int? 1) ; => true
(string? "foo") ; => true
(symbol? +) ; => true

;; or have a certain property

(neg? -1) ; => true
(pos? 1) ; => true
(zero? 0) ; => true

;; Use print or println to write to stdout

(print "Hello") ; => nil (prints "Hello")
(println ", world!") ; => nil (prints ", World\n")
