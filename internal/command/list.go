package command

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/escaletech/tog-cli/internal/color"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:           "list",
	Short:         "List flags in a namespace",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := verifyHost(); err != nil {
			return err
		}

		fmt.Println("Listing flags for namespace", color.Parameter(currentContext.Namespace))
		flags, err := togClient.ListFlags(currentContext.Namespace)
		if err != nil {
			return err
		}

		if len(flags) == 0 {
			fmt.Println(color.Feedback("\n  There are no flags\n"))
			return nil
		}

		renderFlags(flags)

		return nil
	},
}
