package command

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/browser"
	"github.com/r3labs/sse"
	"github.com/spf13/cobra"

	"github.com/escaletech/tog-cli/internal/color"
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:           "login",
	Short:         "Authenticate with management server",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := verifyHost(); err != nil {
			return err
		}

		ch := make(chan *sse.Event)

		token := uuid.New().String()
		listener := sse.NewClient(currentContext.Host + "/auth/cli-notify/" + token)
		go listener.SubscribeChanRaw(ch)

		loginURL := currentContext.Host + "/auth/login?rd=/auth/cli-return&cli_token=" + token
		fmt.Printf("Opening login page:\n  %v\n", color.URL(loginURL))

		if err := browser.OpenURL(loginURL); err != nil {
			return err
		}

		fmt.Println("Waiting for credentials...")
		var data []byte
		for ev := range ch {
			if len(ev.Data) > 0 {
				data = ev.Data
				close(ch)
			}
		}

		var credentials struct{ AuthToken string }
		if err := json.Unmarshal(data, &credentials); err != nil {
			return err
		}

		var (
			authToken = credentials.AuthToken
			ttl       = time.Now().Add(24 * time.Hour)
		)

		if err := configStore.SetContext(currentContext.Host, authToken, ttl); err != nil {
			return err
		}

		fmt.Println("You are logged in.")
		return nil
	},
}
