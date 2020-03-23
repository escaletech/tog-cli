package table

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/alecthomas/chroma/quick"
)

var yamlFormatter = regexp.MustCompile(`"(\w+)":`)

func FormatYaml(value interface{}) string {
	bytes, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	raw := strings.Replace(yamlFormatter.ReplaceAllString(string(bytes), "$1: "), ",", ", ", -1)

	writer := &strings.Builder{}
	if err := quick.Highlight(writer, raw, "yaml", "terminal16m", "solarized-dark256"); err != nil {
		panic(err)
	}

	return writer.String()
}

func IfEmpty(value, empty string) string {
	if value == "" {
		return empty
	}
	return value
}
