package command

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/escaletech/tog-cli/internal/client"
	"github.com/escaletech/tog-cli/internal/config"
)

var currentContext config.Context
var configStore *config.Store
var togClient *client.Client
var buildDate string

var hostFlag string
var namespaceFlag string

func init() {
	rootCmd.PersistentFlags().StringVarP(&hostFlag, "host", "H", "", "target host (overrides default from config)")
	rootCmd.PersistentFlags().StringVarP(&namespaceFlag, "namespace", "n", "", "target namespace (overrides default from config)")
}

var rootCmd = &cobra.Command{
	Use:    "tog",
	Short:  "Tog CLI",
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
	SilenceErrors: true,
	SilenceUsage:  true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if hostFlag != "" {
			host, err := config.NormalizeHost(hostFlag)
			if err != nil {
				return err
			}

			currentContext.Host = host
		}

		if namespaceFlag != "" {
			currentContext.Namespace = namespaceFlag
		} else if currentContext.Namespace == "" {
			currentContext.Namespace = "default"
		}

		togClient = client.New(cmd.Version, currentContext.Host, currentContext.AuthToken)

		return nil
	},
	DisableAutoGenTag: true,
}

func Execute(cstore *config.Store) (*cobra.Command, error) {
	ctx, err := cstore.GetContext()
	if err != nil {
		fmt.Println("Error: failed to get config context:", err)
		return rootCmd, err
	}

	configStore = cstore
	currentContext = ctx
	return rootCmd.ExecuteC()
}

func SetMetadata(version, date string) {
	rootCmd.Version = version
	buildDate = date
}
