// Package service /*
// Copyright © 2025 Arash Rasoulzadeh <arashrasoulzadeh@gmail.com>
package service

import (
	"github.com/arashrasoulzadeh/piperdb/src/models/commands"
	"github.com/arashrasoulzadeh/piperdb/src/service/routers"
)

func RouteCommand(command commands.Command) {

	switch command.Action {
	case commands.ActionEnumCreate:
		routers.Create(command)
		break
	case commands.ActionEnumUpdate:
		routers.Update(command)
		break
	}

}
