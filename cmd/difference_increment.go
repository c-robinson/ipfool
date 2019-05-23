package cmd

import (
	"errors"
	"fmt"
	"math/big"
	"net"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/c-robinson/iplib"
)

var incBy string // string so it can be sent to big.Int.SetString()

var incrementCmd = &cobra.Command{
	Use:   "increment",
	Short: "increment an IP address by <n>",
	Long:  "",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1  {
			return errors.New("requires an ip address")
		}
		if ip := net.ParseIP(args[0]); ip == nil {
			return errors.New("argument is not a valid IP address")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		ip := net.ParseIP(args[0])

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
