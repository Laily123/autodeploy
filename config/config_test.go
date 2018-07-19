package config

import (
	"testing"
)

func TestConfig(t *testing.T) {
	filePath := "../config.toml"
	ParseConfig(filePath)
}
