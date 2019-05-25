package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"github.com/c-robinson/iplib"
)

var deltaCmd = &cobra.Command{
	Use:   "delta",
	Short: "find the distance between two IP addresses",
	Long:  "",
	DisableFlagsInUseLine: true,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ipa := retrieveIPAddress(args[0], v46)
		ipb := retrieveIPAddress(args[0], v46)

		if iplib.EffectiveVersion(ipa) != iplib.EffectiveVersion(ipb) {
			fmt.Println("supplied IP's have mismatched versions")
			os.Exit(1)
		}

		if iplib.EffectiveVersion(ipa) == 4 {
			fmt.Println(iplib.DeltaIP4(ipa, ipb))
		} else {
			fmt.Println(iplib.DeltaIP6(ipa, ipb))
		}
	},
}

func init() {
	differenceRootCmd.AddCommand(deltaCmd)
}
