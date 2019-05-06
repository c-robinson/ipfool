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

var decBy string

var decrementCmd = &cobra.Command{
	Use:   "decrement",
	Short: "decrement an IP address by <n>",
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
