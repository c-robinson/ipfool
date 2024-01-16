package cmd

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/c-robinson/iplib/v2"
	"github.com/c-robinson/iplib/v2/iid"
)

var (
	v6IIDSecretFlag   string
	v6IIDCounterFlag  string
	v6IIDNeighborFlag string
)

var v6MakeIIDCmd = &cobra.Command{
	Use:   "make-iid <address> <mac>",
	Short: "Generate IPv6 EUI64 Interface Identifier",
	Long: `
'v6 make-iid' without flags generates an EUI64-style IPv6 IID. There are
significant security concerns with these addresses, and the flags below can be
used to instead generate an RFC 7217-compliant "semantically opaque" IID. 

For more information on IID's and the flags below see 'ipfool help about-iid'.

Flags:
  --secret <secret>  REQUIRED. A secret key, preferably from a secure source
  --neighbor <mac>   OPTIONAL. The MAC address of the next-hop router
  --count <number>   OPTIONAL. A monotonically incrementing number

Examples:
  % ipfool v6 make-iid 2001:db8:: 01:23:45:67:89:ab
  2001:db8::123:45ff:fe67:89ab

  % ipfool v6 make-iid --counter 1 --secret P00P570Rm 2001:db8:: 01:23:45:67:89:ab
  2001:db8::7da9:ddbb:107a:8c07
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ip := retrieveIPAddress(args[0], v46)
		if iplib.EffectiveVersion(ip) != 6 {
			fmt.Println("IID's can only be generated on IPv6 addresses")
			os.Exit(1)
		}

		mac, err := net.ParseMAC(args[1])
		if err != nil {
			fmt.Println("Could not parse supplied argument as a MAC address: ", args[1])
			os.Exit(1)
		}

		if len(v6IIDSecretFlag) > 1 {
			i, err := strconv.Atoi(v6IIDCounterFlag)
			if err != nil {
				fmt.Println("Counter value must be a number: ", v6IIDCounterFlag)
				os.Exit(1)
			}
			addr, err := iid.MakeOpaqueAddr(ip, mac, int64(i), []byte(v6IIDNeighborFlag), []byte(v6IIDSecretFlag))
			if err != nil {
				fmt.Println("Unable to generate IID: ", err.Error())
				os.Exit(1)
			}
			fmt.Println(addr)
			os.Exit(0)
		}
		addr := iid.MakeEUI64Addr(ip, mac, iid.ScopeLocal)
		fmt.Println(addr)
	},
}

func init() {
	v6RootCmd.AddCommand(v6MakeIIDCmd)
	v6MakeIIDCmd.Flags().StringVarP(&v6IIDCounterFlag, "counter", "c", "1", "monotonically incrementing number")
	v6MakeIIDCmd.Flags().StringVarP(&v6IIDNeighborFlag, "neighbor", "n", "", "MAC address of upstream router")
	v6MakeIIDCmd.Flags().StringVarP(&v6IIDSecretFlag, "secret", "s", "", "closely held secret key (or garbage)")
}
