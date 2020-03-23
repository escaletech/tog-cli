package config

import "time"

type Context struct {
	Config
	AuthToken string
}

type Config struct {
	Namespace string `yaml:"namespace"`
	Host      string `yaml:"host"`
}

type state struct {
	Config   `yaml:"config"`
	Contexts map[string]context `yaml:"contexts"`
}

type context struct {
	Host      string    `yaml:"host"`
	AuthToken string    `yaml:"auth_token"`
	TTL       time.Time `yaml:"ttl"`
}
