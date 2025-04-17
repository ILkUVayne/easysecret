package cmd

import (
	"easysecret/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func init() {
	rootCmd.AddCommand(exampleCmd())
}

func exampleCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "example",
		Short:   "example",
		Long:    "example.",
		Example: "es example",
		Aliases: []string{"e", "example"},
		RunE: func(cmd *cobra.Command, args []string) error {
			p := tea.NewProgram(tui.ExampleModel{}, tea.WithOutput(os.Stdout))
			if _, err := p.Run(); err != nil {
				log.Fatal(err)
			}
			return nil
		},
	}
}
