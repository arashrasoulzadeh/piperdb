// Package commands /*
// Copyright © 2025 Arash Rasoulzadeh <arashrasoulzadeh@gmail.com>
package commands

type ActionEnums string

const (
	ActionEnumCreate ActionEnums = "create"
	ActionEnumUpdate ActionEnums = "update"
)

type ModelsEnum string

const (
	ModelsEnumPipeline  ModelsEnum = "pipeline"
	ModelsEnumNamespace ModelsEnum = "namespace"
)
