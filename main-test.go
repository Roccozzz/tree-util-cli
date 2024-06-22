package main

import (
	"flag"
	"os"

	"github.com/Roccozzz/tree-util-cli/internal"
)

func main() {
	var flagRecursive *bool = flag.Bool("r", false, "recursive tree display")
	flag.Parse()

	internal.PrintTree(os.Args[0], *flagRecursive)
}
