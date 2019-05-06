package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ipfool",
	Short: "IP Fool is a utility for moneying with ip addresses",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Usage()
	},
}

// Execute runs rootCmd and needs normal printer in case onIntialize fails
// and no printer is available
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("while executing command: %s\n", err)
		os.Exit(1)
	}
}
