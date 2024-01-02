package cmd

import (
	"github.com/spf13/cobra"
)

var v6FromCmd = &cobra.Command{
	Use:                   "from",
	Short:                 "convert to IPv6 from ip.arpa or integer",
	DisableFlagsInUseLine: true,
}

func init() {
	v6RootCmd.AddCommand(v6FromCmd)
}
