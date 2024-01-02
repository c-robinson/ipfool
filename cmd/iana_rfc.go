package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-robinson/iplib/v2/iana"
	"github.com/spf13/cobra"
)

var ianaRFCViewFlag bool

var ianaRFCCmd = &cobra.Command{
	Use:   "rfc <rfc number>",
	Short: "list networks affected by a specific special registry RFC",
	Long: `
The 'iana rfc' subcommand returns the list of all networks affected by the
given RFC.

Examples:
  % ipfool iana rfc 1918
  10.0.0.0/8
  172.16.0.0/12
  192.168.0.0/16

  % ipfool iana rfc RFC3849
  2001:db8::/32
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		req := strings.TrimPrefix(strings.ToLower(args[0]), "rfc")

		reservations := []*iana.Reservation{}
		for _, rfc := range iana.Registry {
			for _, ref := range rfc.RFC {
				ref = strings.TrimPrefix(strings.ToLower(ref), "rfc")
				if ref == req {
					reservations = append(reservations, rfc)
				}
			}
		}
		if !ianaRFCViewFlag {
			for _, res := range reservations {
				fmt.Println(res.Network.String())
			}
			os.Exit(0)
		}
		ianaViewHeader()
		for _, res := range reservations {
			ianaViewLine(res)
		}
	},
}

func init() {
	ianaRootCmd.AddCommand(ianaRFCCmd)
	ianaRFCCmd.Flags().BoolVar(&ianaRFCViewFlag, "view", false, "get expanded view of subnets")
}
