package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dbrabera/slip/internal"
)

// options contain the parsed command-line flags and arguments of the program.
type options struct {
	ShowHelp    bool
	ShowVersion bool
	Args        []string
}

func main() {
	opts := parse()

	switch {
	case opts.ShowHelp:
		exit(usage())
	case opts.ShowVersion:
		exit(version())
	case len(opts.Args) == 0:
		exit(repl())
	default:
		exit(exec(opts.Args[0]))
	}
}

// parse parses the command-line flags and arguments.
func parse() options {
	opts := options{}

	flag.BoolVar(&opts.ShowHelp, "h", false, "Show this help")
	flag.BoolVar(&opts.ShowVersion, "v", false, "Show this version")

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: slip [options] [script]")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "An experimental lisp dialect.")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	opts.Args = flag.Args()
	return opts
}

// exec read, parses and evaluates the given file.
func exec(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	_, err = internal.Eval(string(data), internal.NewEnviroment())
	return err
}

// repl executes a Read-eval-print loop until the program is terminated.
func repl() error {
	return internal.REPL()
}

// version prints the version number.
func version() error {
	fmt.Printf("Slip %s\n", internal.Version)
	return nil
}

// usage prints the usage
func usage() error {
	flag.Usage()
	return nil
}

// exit terminates the program with the conventional exit codes depending
// on the given error.
func exit(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
