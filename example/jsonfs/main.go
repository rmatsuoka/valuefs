package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rmatsuoka/futil"
	"github.com/rmatsuoka/valuefs"
)

func usage() {
	fmt.Fprintf(os.Stderr, "jsonfs jsonfile")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	b, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "jsonfs: %v\n", err)
		usage()
	}

	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		fmt.Fprintf(os.Stderr, "jsonfs: %v\n", err)
		usage()
	}
	futil.Shell(valuefs.New(v), os.Stdin, os.Stdout, os.Stderr, "jsonfs % ")
}
