// Package server /*
// Copyright © 2025 Arash Rasoulzadeh <arashrasoulzadeh@gmail.com>
package server

import "net"

func (server *Server) Listen(port string) {
	server.Listener, _ = net.Listen("tcp", ":"+port)
}
func (server *Server) L() net.Listener {
	return server.Listener
}
