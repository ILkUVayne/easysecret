package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"log"
)

const (
	VERSION = "0.0.1"
)

var rootCmd = &cobra.Command{
	Use:     "est",
	Long:    "EasySecret is a simple secret key management tools.",
	Example: "est",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.ExecuteContext(context.Background()); err != nil {
		log.Fatal(err)
	}
}
