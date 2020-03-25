package command

import (
	"fmt"

	"github.com/escaletech/tog-cli/internal/client"
	"github.com/escaletech/tog-cli/internal/table"
)

func renderFlags(flags []*client.Flag) {
	rows := make([][]string, len(flags))
	for i, f := range flags {
		rows[i] = []string{f.Name, table.IfEmpty(f.Description, "-"), table.FormatYaml(f.Rollout)}
	}
	table.Render([]string{"NAME", "DESC", "ROLLOUT"}, rows)
}

func verifyHost() error {
	if currentContext.Host == "" {
		return fmt.Errorf("host is not set")
	}

	return nil
}
