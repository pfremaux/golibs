package main

import (
	"flag"
	"fmt"

	"github.com/pfremaux/golibs/files/pkg/files"
)

func main() {
	// This is the main entry point of the application.
	flag.Parse()
	directory := flag.Arg(0)
	f, err := files.ListFiles(directory)
	if err != nil {
		fmt.Printf("Error listing files: %v\n", err)
		return
	}
	for _, file := range f {
		fmt.Println(file)
	}
}
