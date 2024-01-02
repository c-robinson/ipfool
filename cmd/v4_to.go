package cmd

import (
	"github.com/spf13/cobra"
)

var v4ToCmd = &cobra.Command{
	Use:                   "to",
	Short:                 "convert IPv4 to in-addr.arpa, binary, hexadecimal or integer",
	DisableFlagsInUseLine: true,
}

func init() {
	v4RootCmd.AddCommand(v4ToCmd)
}
