// Package service /*
// Copyright © 2025 Arash Rasoulzadeh <arashrasoulzadeh@gmail.com>
package service

import "github.com/arashrasoulzadeh/piperdb/src/models/server"

func InitAndListen() (s server.Server, err error) {
	s = server.Server{}
	err = s.Listen("3000")
	go Accept(&s)
	return s, err
}
