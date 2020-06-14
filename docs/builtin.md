# Built-In Functions

## +

Usage: (+ & nums)

Returns the addition of nums. If nums is empty returns `0`.

```clojure
(+) ; => 0
(+ 1 2) ; => 3
```

## -

Usage: (- & nums)

Returns the substraction of nums. If num has a single value it returns
its negation, if its empty it returns `0`.

```clojure
(-) ; => 0
(- 1) ; => -1
(- 1 2) ; => -1
```

## \*

Usage: (\* & nums)

Returns the product of all nums. If nums is empty returns `1`.

```clojure
(*) ; => 1
(* 4 2) ; => 8
```

## /

Usage: (/ & nums)

Returns the division of nums. If num has a single value it returns 1 / num, if
its empty it returns `1`.

```clojure
(/) ; => 1
(/ 4) ; => 0
(/ 4 2) ; => 2
```

## mod

Usage: (mod num div)

Returns the modulus of num and div.

```clojure
(mod 5 2) ; => 1
```

## inc

Usage: (inc num)

Returns num + 1.

```clojure
(inc 1) ; => 2
```

## dec

Usage: (dec num)

Returns num - 1.

```clojure
(dec 1) ; => 0
```

## <

Usage: (< & nums)

Returns whether the nums are in ascending order.

```clojure
(< 1 2 3) ; => true
```

## <=

Usage: (<= & nums)

Returns whether the nums are not in descending order.

```clojure
(<= 1 2 2) ; => true
```

## =

Usage: (= & vals)

Returns whether the values are equal.

```clojure
(= 1 1) ; => true
```

## !=

Usage: (!= & vals)

Returns whether the values are not equal.

```clojure
(!= 1 2) ; => true
```

## >=

Usage: (>= & vals)

Returns whether the nums are not in ascending order.

```clojure
(> 3 2 2) ; => true
```

## >

Usage: (> & vals)

Returns whether the nums are in descending order.

```clojure
(> 3 2 1) ; => true
```

## not

Usage: (not x)

Returns the negation of x.

```clojure
(not false) ; => true
```

## bool?

Usage: (bool? x)

Returns whether x is a bool.

```clojure
(bool? true) ; => true
```

## list?

Usage: (list? x)

Returns whether x is a list.

```clojure
(list? '(1 2 3)) ; => true
```

## neg?

Usage: (neg? x)

Returns whether x is less than zero.

```clojure
(neg? -1) ; => true
```

## nil?

Usage: (nil? x)

Returns whether x is nil.

```clojure
(nil? nil) ; => true
```

## int?

Usage: (int? x)

Returns whether x is an int.

```clojure
(int? 1) ; => true
```

## pos?

Usage: (pos? x)

Returns whether x is bigger than zero.

```clojure
(pos? 1) ; => true
```

## string?

Usage: (string? x)

Returns whether x is a string.

```clojure
(string? "str") ; => true
```

## symbol?

Usage: (symbol? x)

Returns whether x is a symbol.

```clojure
(symbol? 'a) ; => true
```

## zero?

Usage: (zero? x)

Returns whether x is zero.

```clojure
(zero? 0) ; => true
```

## print

Usage: (print & vals)

Print the values to stdout.

```clojure
(print 1 true "str" '(1 2 3))
```

## println

Usage: (println & vals)

Print the values to stdout followed by a newline

```clojure
(println 1 true "str" '(1 2 3))
```
