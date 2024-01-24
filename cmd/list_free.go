package cmd

import (
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var listFreeSummaryFlag string

var listFreeCmd = &cobra.Command{
	Use:   "free <network> <list...>",
	Short: "find free addresses in network, based on list",
	Long: `
'list free' takes a network and a list of IP addresses via file or stdin. It
returns a list of all addresses in the network that are not in the list, as
well as summarizing usage.

ipfool.go list free 2001:db8::/125 -
2001:db8::2
2001:db8::5
2001:db8::6
^D
---
2001:db8::
2001:db8::1
2001:db8::3
2001:db8::4
2001:db8::6
2001:db8::7
Network: 2001:db8::/125
  Size: 8,  Used: 3 (37.50%)  Free: 6
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.MaximumNArgs(2),
	ValidArgs:             []string{"file", "summary"},
	Run: func(cmd *cobra.Command, args []string) {
		var (
			showSummary = false
			showFree    = true
		)
		switch listFreeSummaryFlag {
		case "include":
			showSummary = true
		case "exclude":
			showSummary = false
		case "only":
			showSummary = true
			showFree = false
		default:
			fmt.Println("Invalid argument: --summary must be one of 'include', 'exclude', or 'only'")
			os.Exit(1)
		}

		iplist, isStdin, err := retrieveIPList(args, cmd.InOrStdin())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if len(iplist) == 0 {
			fmt.Println(err)
			os.Exit(1)
		}

		if len(args) == 0 {
			fmt.Println("Missing required argument: <network>")
			os.Exit(1)
		}
		ipnet := retrieveIPNetwork(args[0], 46)

		sort.Sort(iplib.ByIP(iplist))
		iplist = ipListUniquer(iplist)
		iplist = ipListPruner(ipnet, iplist)

		if isStdin {
			// this is to provide some visual separation between the input and
			// output when using stdin
			fmt.Println("\n---")
		}

		if len(iplist) == 0 {
			fmt.Println("All space free (did you mean to use 'net enumerate'?")
			os.Exit(1)
		}

		xip := ipnet.FirstAddress()
		pos := 0
		free := 0
		var done error
		for {
			if iplist[pos].Equal(xip) && pos < len(iplist)-1 {
				pos++
			} else {
				free++
				if showFree {
					fmt.Println(xip)
				}
			}
			xip, done = genericNetIncrement(xip, ipnet)
			if done != nil {
				break
			}
		}

		if showSummary {
			count, pct := genericNetCountAndUsedPct(ipnet, len(iplist))
			fmt.Printf("Network: %s\n", ipnet)
			fmt.Printf("  Size: %s,  Used: %d (%0.2f%%)  Free: %d\n", count, len(iplist), pct, free)
		}
	},
}

// genericNetIncrement is a helper function that takes an iplib.Net and an
// int corresponding to the number of addresses in that that are in use. It
// returns the count as a string and the percentage of addresses in use as a
// float64
func genericNetCountAndUsedPct(n iplib.Net, used int) (string, float64) {
	switch n.Version() {
	case 4:
		pct := (float64(used) / float64(n.(iplib.Net4).Count())) * 100
		return strconv.Itoa(int(n.(iplib.Net4).Count())), pct

	case 6:
		// uuuugh. There's no easy way to coerce a uint128 to a float64 so
		// we need to convert to our old friend *big
		cntBigFloat := new(big.Float).SetInt(n.(iplib.Net6).Count().Big())
		usdBigFloat := new(big.Float).SetInt(big.NewInt(int64(used)))

		pctBigFloat := new(big.Float).Quo(usdBigFloat, cntBigFloat)
		pctBigFloat.Mul(pctBigFloat, big.NewFloat(100))
		pct, _ := pctBigFloat.Float64()

		return n.(iplib.Net6).Count().String(), pct
	}
	return "unknown IP version", -1
}

func init() {
	listRootCmd.AddCommand(listFreeCmd)
	listFreeCmd.Flags().StringVarP(&listFreeSummaryFlag, "summary", "s", "include", "summary display (include|exclude|only)")
}
