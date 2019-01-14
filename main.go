package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/flamacue/go-brainfuck-but-not-yourself/compiler"
	"github.com/flamacue/go-brainfuck-but-not-yourself/machine"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file argument")
		os.Exit(1)
	}
	file := os.Args[1]
	code, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	compiler := compiler.New(string(code))
	instructions := compiler.Compile()
	m := machine.New(instructions, os.Stdin, os.Stdout)
	m.Execute()
}
