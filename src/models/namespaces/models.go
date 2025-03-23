// Package namespaces /*
// Copyright © 2025 Arash Rasoulzadeh <arashrasoulzadeh@gmail.com>
package namespaces

import (
	"github.com/arashrasoulzadeh/piperdb/src/models/pipeline"
)

type Namespace struct {
	Name      string             `json:"name"`
	Pipelines pipeline.Pipelines `json:"pipelines"`
}
