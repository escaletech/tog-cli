package command

import (
	"fmt"

	"github.com/escaletech/tog-cli/internal/color"
	"github.com/escaletech/tog-cli/internal/config"
	"github.com/escaletech/tog-cli/internal/table"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(configCmd)
}

var configCmd = &cobra.Command{
	Use:           "config [key [value]]",
	Short:         "Deal with configuration options",
	Args:          cobra.MaximumNArgs(2),
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		switch len(args) {
		case 1:
			return printConfigKey(args[0])

		case 2:
			if err := setConfigKey(args[0], args[1]); err != nil {
				return err
			}

			renderConfig()

		default:
			fmt.Println("Current configuration")
			renderConfig()
		}

		return nil
	},
}

func printConfigKey(key string) error {
	switch key {
	case "host":
		fmt.Println(currentContext.Host)
	case "namespace":
		fmt.Println(currentContext.Namespace)
	default:
		return fmt.Errorf("unknown config key %v", key)
	}

	return nil
}

func setConfigKey(key, value string) error {
	switch key {
	case "host":
		host, err := config.NormalizeHost(value)
		if err != nil {
			return err
		}
		currentContext.Host = host
	case "namespace":
		currentContext.Namespace = value
	default:
		return fmt.Errorf("unknown config key %v", key)
	}

	fmt.Println("Updating", color.Parameter(key))
	if err := configStore.SetConfig(currentContext.Config); err != nil {
		return fmt.Errorf("failed to save config: %v", err)
	}

	return nil
}

func renderConfig() {
	table.Render([]string{"NAME", "VALUE"}, [][]string{
		{"host", table.IfEmpty(currentContext.Config.Host, "-")},
		{"namespace", table.IfEmpty(currentContext.Config.Namespace, "-")},
	})
}
