package cmd

import (
	"github.com/spf13/cobra"
)

var v6FromCmd = &cobra.Command{
	Use:                   "from",
	Short:                 "convert from ip6.arpa or integer to IPv6",
	DisableFlagsInUseLine: true,
}

func init() {
	v6RootCmd.AddCommand(v6FromCmd)
}
