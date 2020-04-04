package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "proxy",
	Short: "proxy bloomfilter",
	Long:  "...",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var bloomFilter = &cobra.Command{
	Use:   "bloomfilter",
	Short: "start bloomfilter nodes",
	Long:  "this is a test",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		for _, str := range args {
			fmt.Println(str)
		}
	},
}

func init() {
	rootCmd.AddCommand(bloomFilter)
}
func main() {
	// mp := make(map[string]struct{}, 0)
	// for i := 0; i < 32; i++ {
	// 	for j := 0; j <= 255; j++ {
	// 		for k := 0; k <= 255; k++ {
	// 			str := fmt.Sprintf("10.%d.%d.%d", i, j, k)
	// 			mp[str] = struct{}{}
	// 		}
	// 	}
	// }
	fmt.Println("finish!")
	time.Sleep(time.Minute * 3)
}
