package cmd

import (
	"github.com/spf13/cobra"
)

var differenceRootCmd = &cobra.Command{
	Use:   "difference",
	Short: "commands for calculating the difference between IP addresses",
	Long: `
The 'difference' subtree holds commands that calculate the difference between
IP addresses: delta, decrement and increment.
	`,
}

func init() {
	rootCmd.AddCommand(differenceRootCmd)
}
