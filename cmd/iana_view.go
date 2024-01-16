package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/c-robinson/iplib/v2/iana"
	"github.com/spf13/cobra"
)

var ianaViewV4Flag, ianaViewV6Flag bool

var ianaViewCmd = &cobra.Command{
	Use:   "view",
	Short: "view special registry RFCs",
	Long: `
With no arguments, the 'iana rfc' subcommand returns the list of all RFCs in
the IPv4 and IPv6 special purpose registries.

Flags: 
  --v4   only show IPv4 RFCs
  --v6   only show IPv6 RFCs
`,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		showV4 := true
		showV6 := true
		if ianaViewV4Flag && !ianaViewV6Flag {
			showV6 = false
		}
		if !ianaViewV4Flag && ianaViewV6Flag {
			showV4 = false
		}
		ianaViewHeader()
		for _, res := range iana.Registry {
			if !showV4 && res.Network.Version() == 4 {
				continue
			}
			if !showV6 && res.Network.Version() == 6 {
				continue
			}
			ianaViewLine(res)
		}
	},
}

func ianaViewHeader() {
	fmt.Printf("%-18s %-42s %-16s\n", "Network", "Description", "RFC(s)")
}

func ianaViewLine(res *iana.Reservation) {
	for i, ref := range res.RFC {
		res.RFC[i], _ = strings.CutPrefix(ref, "RFC")
	}
	sort.Strings(res.RFC)
	fmt.Printf("%-18s %-42s %-16s\n", res.Network.String(), res.Title, strings.Join(res.RFC, ", "))
}

func init() {
	ianaRootCmd.AddCommand(ianaViewCmd)
	ianaViewCmd.Flags().BoolVar(&ianaViewV4Flag, "v4", false, "view only IPv4 registry")
	ianaViewCmd.Flags().BoolVar(&ianaViewV6Flag, "v6", false, "view only IPv6 registry")
}
