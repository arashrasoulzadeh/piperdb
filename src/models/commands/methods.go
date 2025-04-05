// Package commands /*
// Copyright © 2025 Arash Rasoulzadeh <arashrasoulzadeh@gmail.com>
package commands

func (c *Command) SetClientId(clientId string) {
	c.clientId = clientId
}
func (c *Command) GetClientId() string {
	return c.clientId
}
