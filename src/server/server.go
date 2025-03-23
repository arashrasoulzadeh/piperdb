// Package server /*
// Copyright © 2025 Arash Rasoulzadeh <arashrasoulzadeh@gmail.com>
package server

import "github.com/arashrasoulzadeh/piperdb/src/models/server"

func InitAndListen() server.Server {
	s := server.Server{}
	s.Listen("3000")
	return s
}
