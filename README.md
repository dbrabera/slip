# Slip

[![Build Status](https://travis-ci.org/dbrabera/slip.svg?branch=master)](https://travis-ci.org/dbrabera/slip)

**Slip** is a Lisp dialect built to learn more about the implementation of programming languages and for the fun of it.

## Try it yourself

At the moment there are no binary releases for the language. To try it yourself you would need to build it from source by either cloning this repository or by executing `go get`:

```
$ go get github.com/dbrabera/slip
```

Once you have slip command avaliable, you can start the REPL and run expressions by executing it without arguments:

```
$ slip
Slip f29f33b
slip:1:> (println "It's ALIVE!")
"It's ALIVE!"
nil
slip:2:>
```

Alternatively you can execute pass the path for a slip `.sp` file to execute a script:

```
$ cat hello.sp
(println "It's ALIVE!")
```

```
$ slip hellp.sp
It's ALIVE!
```

## Development

### Prerequisites

- Go (1.14 or later)
- Golangci-lint (1.27 or later)

### Build

To build the source code do:

```
$ make build
```

### Lint

To lint the source code do:

```
$ make lint
```

### Test

To run the tests do:

```
$ make test
```

## Documentation

- [Special Forms](./docs/forms.md)

- [Built-In Functions](./docs/builtin.md)

## License

MIT
