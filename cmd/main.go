package main

import (
	"fmt"
	"os"

	"github.com/diegolikescode/go-interpreter/internal/repl"
)

func main() {
	fmt.Println("Hey hey programmer, this is your REPL")

	repl.Start(os.Stdin, os.Stdout)
}
