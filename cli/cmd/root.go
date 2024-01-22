package cmd

import (
	"github.com/spf13/cobra"
)

var Version string

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "libstone",
		Short:   "A golang implementation for stone binary packages",
		Version: Version,
	}
	cmd.AddCommand(
		&cobra.Command{
			Use:     "inspect",
			Short:   "Inspect stone package contents",
			Version: Version,
			RunE:    Inspect,
		},
	)
	return cmd
}
