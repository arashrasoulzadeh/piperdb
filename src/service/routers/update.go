// Package routers /*
// Copyright © 2025 Arash Rasoulzadeh <arashrasoulzadeh@gmail.com>
package routers

import (
	"github.com/arashrasoulzadeh/piperdb/src/models/commands"
	"github.com/arashrasoulzadeh/piperdb/src/service/actions/namespace"
)

func Update(command commands.Command) {
	switch command.Model {
	case commands.ModelsEnumNamespace:
		namespace.Update(command.Value)
		break
	}
}
