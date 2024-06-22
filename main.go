package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Roccozzz/tree-util-cli/internal"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: tree-util-cli <absolute-path>")
		os.Exit(1)
	}

	var path string
	if os.Args[1] == "." {
		path, _ = os.Getwd()
	} else {
		path, _ = filepath.Abs(os.Args[1])
	}

	internal.PrintTree(path, "")
}
