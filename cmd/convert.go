package cmd

import (
	"github.com/spf13/cobra"
)

var convertRootCmd = &cobra.Command{
	Use:   "convert",
	Short: "commands converting IP addresses between formats",
	Long: `
The 'convert' subtree holds commands that convert between representations of
IP addresses: dotted decimal, hex and integer.
	`,
}

func init() {
	rootCmd.AddCommand(convertRootCmd)
}
