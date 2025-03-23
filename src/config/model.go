// Package config /*
// Copyright © 2025 Arash Rasoulzadeh <arashrasoulzadeh@gmail.com>
package config

type Instance struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Encode bool   `json:"encode"`
}

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}
type Config struct {
	Instance Instance `json:"instance"`
	Server   Server   `json:"server"`
}
