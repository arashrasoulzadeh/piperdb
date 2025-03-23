// Package config /*
// Copyright © 2025 Arash Rasoulzadeh <arashrasoulzadeh@gmail.com>
package config

import (
	"github.com/arashrasoulzadeh/piperdb/src/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitConfig(t *testing.T) {
	c := config.NewConfig(config.WithInstancePath("data"), config.WithInstanceName("test"), config.WithServerHost("127.0.0.1"), config.WithServerPort("1234"))
	assert.Equal(t, c.Instance.Path, "data")
	assert.Equal(t, c.Instance.Name, "test")
	assert.Equal(t, c.Server.Host, "127.0.0.1")
	assert.Equal(t, c.Server.Port, "1234")
}

func TestSampleConfig(t *testing.T) {
	c := config.NewConfig(config.WithConfigFile("test.toml"))
	assert.Equal(t, c.Instance.Path, "data")
	assert.Equal(t, c.Instance.Name, "test")
	assert.Equal(t, c.Server.Host, "127.0.0.1")
	assert.Equal(t, c.Server.Port, "1234")
}
