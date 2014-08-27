package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

const Version = "0.0.1"

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
		os.Exit(usage())
	}

	if options.Help {
		os.Exit(usage())
	}

	if options.Version {
		os.Exit(version())
	}

	slip := NewSlip()

	if options.Script != "" {
		source, err := ioutil.ReadFile(options.Script)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading script: %s\n", err)
			os.Exit(1)
		}
		os.Exit(slip.Run(string(source)))
	} else {
		os.Exit(slip.Repl())
	}
}

func version() int {
	fmt.Printf("Slip %s\n", Version)
	return 0
}

func usage() int {
	fmt.Fprintln(os.Stderr, "Usage: slip [options] [script [args]]")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Options:")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "-h, --help       Show this help")
	fmt.Fprintln(os.Stderr, "-v, --version    Show version information")
	fmt.Fprintln(os.Stderr, "")
	return 1
}
