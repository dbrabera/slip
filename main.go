package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dbrabera/slip/internal"
)

type Options struct {
	Help    bool
	Version bool
	Script  string
	Args    []string
}

func parse(args []string) (*Options, error) {
	var options Options

	for i, arg := range args {
		if arg == "-h" || arg == "--help" {
			options.Help = true
			continue
		}

		if arg == "-v" || arg == "--version" {
			options.Version = true
			continue
		}

		if arg[0] == '-' {
			return nil, errors.New("Unexpected flag")
		}

		options.Script = arg
		options.Args = args[i+1:]
		break
	}

	return &options, nil
}

func main() {
	options, err := parse(os.Args[1:])
	if err != nil {
		exit(usage())
	}

	if options.Help {
		exit(usage())
	}

	if options.Version {
		exit(version())
	}

	if options.Script != "" {
		exit(exec(options.Script))
	}

	exit(repl())
}

func exec(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	_, err = internal.Eval(string(data), internal.NewEnviroment())
	return err
}

func repl() error {
	return internal.REPL()
}

func version() error {
	fmt.Printf("Slip %s\n", internal.Version)
	return nil
}

func usage() error {
	fmt.Fprintln(os.Stderr, "Usage: slip [options] [script [args]]")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Options:")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "-h, --help       Show this help")
	fmt.Fprintln(os.Stderr, "-v, --version    Show version information")
	fmt.Fprintln(os.Stderr, "")
	return nil
}

func exit(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
