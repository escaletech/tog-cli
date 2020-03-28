package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:           "version",
	Aliases:       []string{"v"},
	Short:         "Get Tog CLI version",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("tog version %v (%v)\n", cmd.Parent().Version, buildDate)
		fmt.Printf("https://github.com/escaletech/tog-cli/releases/tag/%v\n", cmd.Parent().Version)
		return nil
	},
}
