package command

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/escaletech/tog-cli/internal/client"
	"github.com/escaletech/tog-cli/internal/color"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:           "get [name]",
	Short:         "Get a feature flag",
	Args:          cobra.ExactArgs(1),
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := verifyHost(); err != nil {
			return err
		}

		fmt.Println("Getting flag " + color.Parameter(currentContext.Namespace) + "/" + color.Parameter(args[0]))

		f, err := togClient.GetFlag(currentContext.Namespace, args[0])
		if err == client.ErrNotFound {
			fmt.Println(color.Feedback("\n  Flag not found\n"))
			return nil
		} else if err != nil {
			return err
		}

		renderFlags([]*client.Flag{f})

		return nil
	},
}
