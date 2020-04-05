package cmd

import (
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/c-robinson/iplib"
)

var incBy string // string so it can be sent to big.Int.SetString()

var incrementCmd = &cobra.Command{
	Use:   "increment",
	Short: "increment an IP address by a given amount (default 1)",
	Long: `
Increment takes an IP address as input. If no arguments are given it will
increment the address by one. the --by argument is used to specify a number.`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	ValidArgs:             []string{"by"},
	Run: func(cmd *cobra.Command, args []string) {
		ip := retrieveIPAddress(args[0], v46)

		i, err := strconv.Atoi(incBy)
		if err == nil {
			fmt.Println(iplib.IncrementIPBy(ip, uint32(i)))
		} else {
			z, ok := new(big.Int).SetString(incBy, 10)
			if !ok {
				fmt.Printf("cannot convert '%s' to an integer", incBy)
				os.Exit(1)
			}
			fmt.Println(iplib.IncrementIP6By(ip, z))
		}
	},
}

func init() {
	differenceRootCmd.AddCommand(incrementCmd)
	incrementCmd.Flags().StringVar(&incBy, "by", "1", "increment address by count")
}
