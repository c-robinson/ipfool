package cmd

import (
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/c-robinson/iplib"
)

var decBy string

var decrementCmd = &cobra.Command{
	Use:   "decrement",
	Short: "decrement an IP address by a given amount (default 1)",
	Long: `decrement takes an IP address as input. If no arguments are given it will
decrement the address by one. the --by argument is used to specify a number.`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	ValidArgs:             []string{"by"},
	Run: func(cmd *cobra.Command, args []string) {
		ip := retrieveIPAddress(args[0], v46)

		i, err := strconv.Atoi(decBy)
		if err == nil {
			fmt.Println(iplib.DecrementIPBy(ip, uint32(i)))
		} else {
			z, ok := new(big.Int).SetString(decBy, 10)
			if !ok {
				fmt.Printf("cannot convert '%s' to an integer", decBy)
				os.Exit(1)
			}
			fmt.Println(iplib.DecrementIP6By(ip, z))
		}
	},
}

func init() {
	differenceRootCmd.AddCommand(decrementCmd)
	decrementCmd.Flags().StringVar(&decBy, "by", "1", "decrement address by count")
}
