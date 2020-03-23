package config

import (
	"fmt"
	"net/url"
	"strings"
)

func NormalizeHost(host string) (string, error) {
	if !strings.HasPrefix("https://", host) && !strings.HasPrefix("http://", host) {
		host = "https://" + host
	}

	if _, err := url.ParseRequestURI(host); err != nil {
		return "", fmt.Errorf("invalid host: %v", err)
	}

	return host, nil
}
