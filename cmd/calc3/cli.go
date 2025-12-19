package main

import (
	"calcPlus/internal/calc3"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <calc3-file>\n", os.Args[0])
		os.Exit(1)
	}

	codeFile := os.Args[1]
	code, err := os.ReadFile(codeFile)
	if err != nil {
		return
	}

	calc3.RunInterpreter(string(code), os.Stdin, os.Stdout)
}
