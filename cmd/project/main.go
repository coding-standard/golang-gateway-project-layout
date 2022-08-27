package main

import (
	"fmt"
	"github.com/coding-standard/golang-project-layout/internal/cmd"
	"os"
)

func main() {
	if err := cmd.GetRootCommand().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
