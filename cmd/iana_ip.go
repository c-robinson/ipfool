package cmd

import (
	"fmt"
	"net"
	"os"
	"sort"

	"github.com/c-robinson/iplib/v2"
	"github.com/c-robinson/iplib/v2/iana"
	"github.com/spf13/cobra"
)

var ianaIPViewFlag bool

var ianaIPCmd = &cobra.Command{
	Use:   "ip <address|network>",
	Short: "list special registry RFCs that apply to an IP address or network",
	Long: `
'iana ip'  returns the list of all RFCs that apply to the given IP address or
network.

Examples:
  % ipfool iana ip 192.168.1.0/24
  RFC1918

  % ipfool iana ip 2001::/32    
  RFC2928
  RFC4380
  RFC8190
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		reservations := []*iana.Reservation{}

		_, cnet, err := iplib.ParseCIDR(args[0])
		if err == nil {
			for _, res := range iana.Registry {
				if res.Network.ContainsNet(cnet) {
					reservations = append(reservations, res)
				}
			}
		} else if caddr := net.ParseIP(args[0]); caddr != nil {
			version := iplib.EffectiveVersion(caddr)
			for _, res := range iana.Registry {

				// this is necessary to prevent all IPv4 addresses from
				// matching the IPv6 4in6 reservation from RFC4291
				if res.Network.Version() != version {
					continue
				}
				if res.Network.Contains(caddr) {
					reservations = append(reservations, res)
				}
			}
		} else {
			os.Exit(1)
		}

		if !ianaIPViewFlag {
			for _, res := range reservations {
				sort.Strings(res.RFC)
				for _, rfc := range res.RFC {
					fmt.Println(rfc)
				}
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
	ianaRootCmd.AddCommand(ianaIPCmd)
	ianaIPCmd.Flags().BoolVar(&ianaIPViewFlag, "view", false, "get expanded view of subnets")

}
