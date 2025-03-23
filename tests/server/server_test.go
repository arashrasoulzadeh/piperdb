// Package server /*
// Copyright © 2025 Arash Rasoulzadeh <arashrasoulzadeh@gmail.com>
package server

import (
	"github.com/arashrasoulzadeh/piperdb/src/server"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	t.Run("create and listen", func(t *testing.T) {
		server.InitAndListen()
		conn, err := net.Dial("tcp", ":3000")
		assert.NoError(t, err)
		conn.Close()
	})
}
