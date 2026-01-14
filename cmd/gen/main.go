package main

import (
	"fmt"
	"os"

	"study-stack/internal/codegen/module"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("  go generate mod <name>")
		os.Exit(1)
	}

	cmd := os.Args[1]
	arg := os.Args[2]

	switch cmd {
	case "mod":
		gen := module.New(arg)
		gen.Generate()
	default:
		fmt.Println("Unknown generator:", cmd)
	}
}
