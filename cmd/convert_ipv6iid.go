package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net"
	"os"
	"strconv"

	"github.com/c-robinson/iplib"
	"github.com/c-robinson/iplib/iid"
)

var (
	secret   string
	counter  string
	neighbor string
)

var ipv6iidCmd = &cobra.Command{
	Use:   "iid",
	Short: "Generate IPv6 EUI64 Interface Identifier",
	Long: `
The 'iid' subcommand expects two arguments: an IPv6 address and a hardware
address (typically the target interface's MAC address). Given no other inputs
it will generate an EUI64-style IPv6 IID scoped for local-subnet use ONLY,
this is because and IID generated in this way leaks personal information about
the addresses owner (specifically it leaks the MAC address of the host, which
can be problematic for laptops or mobile devices).

There are three additional flags to 'iid' that can be used to generate a
"semantically opaque" IID as described in RFC 7217. The only required one is
--secret which takes any  text as an argument and uses it as an encryption
key. You SHOULD also supply --neighbor and give it, as an argument, the MAC
address of the next-hop router, as well as --count, which should be a
monotonically incrementing number. These two additional pieces of data allow
'iid' to more effectively randomize the address. Especially for mobile
devices.`,
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

		if len(secret) > 1 {
			i, err := strconv.Atoi(counter)
			if err != nil {
				fmt.Println("Counter value must be a number: ", counter)
				os.Exit(1)
			}
			addr, err := iid.MakeOpaqueAddr(ip, mac, int64(i), []byte(neighbor), []byte(secret))
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
	convertRootCmd.AddCommand(ipv6iidCmd)
	ipv6iidCmd.Flags().StringVarP(&counter, "counter", "c", "1", "monotonically incrementing number")
	ipv6iidCmd.Flags().StringVarP(&neighbor, "neighbor", "n", "", "MAC address of neighbor")
	ipv6iidCmd.Flags().StringVarP(&secret, "secret", "s", "", "closely held secret key (or garbage)")
}
