package main

import (
	"fmt"
	"os"

	"github.com/GZGavinZhao/libstone-go/cli/cmd"
)

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
