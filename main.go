package main

import (
	"os"
)

func main() {
	r := NewRootCmd()
	if err := r.Execute(); err != nil {
		os.Exit(1)
	}
}
