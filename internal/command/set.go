package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	"github.com/escaletech/tog-cli/internal/client"
	"github.com/escaletech/tog-cli/internal/color"
)

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().StringP("rollout", "r", "", "Set rollout strategy")
	setCmd.Flags().StringP("description", "d", "", "Set description")
	setCmd.Flags().Bool("on", false, "Set flag as globally on")
	setCmd.Flags().Bool("off", false, "Set flag as globally off")
}

var setCmd = &cobra.Command{
	Use:           "set",
	Short:         "Set a flag's attributes",
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
			f = &client.Flag{Namespace: currentContext.Namespace, Name: args[0]}
		} else if err != nil {
			return err
		}

		options := cmd.Flags()

		if raw, _ := options.GetString("rollout"); raw != "" {
			var rollout []*client.Rollout
			if err := yaml.Unmarshal([]byte(raw), &rollout); err != nil {
				return err
			}

			fmt.Println("Setting", color.Parameter("rollout"))
			f.Rollout = rollout
		} else if on, _ := options.GetBool("on"); on {
			fmt.Println("Setting", color.Parameter("rollout"))
			f.Rollout = []*client.Rollout{{Value: true}}
		} else if off, _ := options.GetBool("off"); off {
			fmt.Println("Setting", color.Parameter("rollout"))
			f.Rollout = []*client.Rollout{{Value: false}}
		}

		if descr, _ := options.GetString("description"); descr != "" {
			fmt.Println("Setting", color.Parameter("description"))
			f.Description = descr
		}

		res, err := togClient.SaveFlag(f)
		if err != nil {
			return err
		}

		renderFlags([]*client.Flag{res})

		return nil
	},
}
