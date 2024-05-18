// Package main is the entry point of the honeycomb application.
package main

import (
	"fmt"
	"os"

	hc "github.com/nao1215/honeycomb/cmd/honeycomb"
)

func main() {
	if err := hc.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "honeycomb: %s\n", err.Error())
	}
}
