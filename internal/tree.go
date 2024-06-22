package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

func PrintTree(path string, prefix string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, entry := range entries {
		isLast := i == len(entries)-1
		printEntry(entry, prefix, isLast)

		if entry.IsDir() {
			newPrefix := prefix
			if isLast {
				newPrefix += "    "
			} else {
				newPrefix += "│   "
			}
			PrintTree(filepath.Join(path, entry.Name()), newPrefix)
		}
	}
}

func printEntry(entry os.DirEntry, prefix string, isLast bool) {
	var connector string
	if isLast {
		connector = "└── "
	} else {
		connector = "├── "
	}
	fmt.Println(prefix + connector + entry.Name())
}
