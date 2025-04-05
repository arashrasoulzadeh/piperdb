// Package commands /*
// Copyright © 2025 Arash Rasoulzadeh <arashrasoulzadeh@gmail.com>
package commands

type Command struct {
	clientId string
	Action   ActionEnums       `json:"action"`
	Model    ModelsEnum        `json:"model"`
	Value    string            `json:"value"`
	Args     map[string]string `json:"args"`
}
