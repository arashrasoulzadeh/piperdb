// Package server /*
// Copyright © 2025 Arash Rasoulzadeh <arashrasoulzadeh@gmail.com>
package server

import (
	"encoding/json"
	"github.com/arashrasoulzadeh/piperdb/src/config"
	"github.com/arashrasoulzadeh/piperdb/src/models/commands"
	"github.com/arashrasoulzadeh/piperdb/src/service"
	"github.com/stretchr/testify/assert"
	"net"
	"os"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	path := "./temp"
	config.NewConfig(config.WithInstancePath(path))
	err := os.RemoveAll(path)
	assert.NoError(t, err)

	t.Run("create and listen", func(t *testing.T) {
		listen, errInitAndListen := service.InitAndListen()
		assert.NoError(t, errInitAndListen)
		assert.NotNil(t, listen)
		conn, err := net.Dial("tcp", ":3000")
		assert.NoError(t, err)

		createMessage := commands.Command{
			Action: commands.ActionEnumCreate,
			Args:   make(map[string]string),
			Model:  commands.ModelsEnumNamespace,
			Value:  "piper",
		}

		jMessage, _ := json.Marshal(createMessage)

		write, writeErr := conn.Write([]byte(string(jMessage) + "\n"))

		updateMessage := commands.Command{
			Action: commands.ActionEnumUpdate,
			Args:   make(map[string]string),
			Model:  commands.ModelsEnumNamespace,
			Value:  "test_namespace",
		}

		jMessage, _ = json.Marshal(updateMessage)

		write, writeErr = conn.Write([]byte(string(jMessage) + "\n"))

		assert.NoError(t, writeErr)
		assert.NotEqual(t, write, 0)
		assert.NoError(t, conn.Close())
		time.Sleep(2 * time.Second)
	})
}
