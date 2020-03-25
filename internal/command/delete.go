package command

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/escaletech/tog-cli/internal/client"
	"github.com/escaletech/tog-cli/internal/color"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:           "delete",
	Aliases:       []string{"del"},
	Short:         "Delete a flag's",
	Args:          cobra.ExactArgs(1),
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := verifyHost(); err != nil {
			return err
		}

		fmt.Println("Deleting flag " + color.Parameter(currentContext.Namespace) + "/" + color.Parameter(args[0]))

		err := togClient.DeleteFlag(currentContext.Namespace, args[0])
		if err == client.ErrNotFound {
			fmt.Println(color.Feedback("\n  Flag not found\n"))
			return nil
		}
		if err != nil {
			return err
		}

		fmt.Println(color.Feedback("\n  Flag deleted\n"))

		return nil
	},
}
