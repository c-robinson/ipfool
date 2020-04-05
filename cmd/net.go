package cmd

import (
	"github.com/spf13/cobra"
)

var netRootCmd = &cobra.Command{
	Use:   "net",
	Short: "commands for viewing and manipulating IP netblocks",
	Long: `
The 'net subtree holds commands for viewing IP netblocks as well as
experimenting with subnets and supernets.`,
}

func init() {
	rootCmd.AddCommand(netRootCmd)
}
