package main

import (
	"fmt"
	"os"

	"github.com/der-eismann/libstone/cli/cmd"
)

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
