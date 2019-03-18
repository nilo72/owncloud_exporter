package config

import (
	"time"
)

// Server defines the general server configuration.
type Server struct {
	Addr string
	Path string
}

// Logs defines the level and color for log configuration.
type Logs struct {
	Level  string
	Pretty bool
}

// Target defines the target specific configuration.
type Target struct {
	Address  string
	Timeout  time.Duration
}

// Collector defines the collector specific configuration.
type Collector struct {
	Users bool
}

// Config is a combination of all available configurations.
type Config struct {
	Server    Server
	Logs      Logs
	Target    Target
	Collector Collector
}

// Load initializes a default configuration struct.
func Load() *Config {
	return &Config{}
}
