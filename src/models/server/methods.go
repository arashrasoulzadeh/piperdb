// Package server /*
// Copyright © 2025 Arash Rasoulzadeh <arashrasoulzadeh@gmail.com>
package server

import "net"

func (s *Server) Listen(port string) (err error) {
	s.Listener, err = net.Listen("tcp", ":"+port)
	return err
}
func (s *Server) L() net.Listener {
	return s.Listener
}
