package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var banner = `
____ ____ ____ _   _ ____ ____ ____ ____ ____ ___ 
|___ |__| [__   \_/  [__  |___ |    |__/ |___  |  
|___ |  | ___]   |   ___] |___ |___ |  \ |___  |
`

func init() {
	rootCmd.AddCommand(versionCmd())
}

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Short:   "Display EasySecret version",
		Example: "es version",
		Aliases: []string{"v", "version"},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(banner)
			fmt.Println("EasySecret Version " + VERSION)
		},
	}
}
