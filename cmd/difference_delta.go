package cmd

import (
	"errors"
	"fmt"
	"net"

	"github.com/spf13/cobra"

	"github.com/c-robinson/iplib"
)

var deltaCmd = &cobra.Command{
	Use:   "delta",
	Short: "find the distance between two IP addresses",
	Long:  "",
	Args: func(cmd *cobra.Command, args []string) error {
		var ipa, ipb net.IP
		if len(args) != 2 {
			return errors.New("requires 2 IP addresses as arguments")
		}
		if ipa = net.ParseIP(args[0]); ipa == nil {
			return errors.New("first argument is not a valid IP address")
		}
		if ipb = net.ParseIP(args[1]); ipb == nil {
            return errors.New("second argument is not a valid IP address")
        }
		if iplib.EffectiveVersion(ipa) != iplib.EffectiveVersion(ipb) {
			return errors.New("mismatched IP versions")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		ipa := net.ParseIP(args[0])
		ipb := net.ParseIP(args[1])
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
