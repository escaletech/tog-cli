package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/escaletech/tog-cli/internal/config"
)

func TestNormalizeHost(t *testing.T) {
	t.Run("prepends https by default", func(t *testing.T) {
		actual, err := config.NormalizeHost("mysite.com")
		assert.NoError(t, err)
		assert.Equal(t, "https://mysite.com", actual)
	})

	t.Run("keeps https prefix", func(t *testing.T) {
		actual, err := config.NormalizeHost("https://mysite.com")
		assert.NoError(t, err)
		assert.Equal(t, "https://mysite.com", actual)
	})

	t.Run("keeps http prefix", func(t *testing.T) {
		actual, err := config.NormalizeHost("http://mysite.com")
		assert.NoError(t, err)
		assert.Equal(t, "http://mysite.com", actual)
	})
}
