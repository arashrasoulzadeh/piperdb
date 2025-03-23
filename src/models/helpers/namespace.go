// Package namespaces /*
// Copyright © 2025 Arash Rasoulzadeh <arashrasoulzadeh@gmail.com>
package helpers

import (
	"github.com/arashrasoulzadeh/piperdb/src/models/namespaces"
)

func LoadNamespaceFromMeta(path string) (*namespaces.Namespace, error) {
	return LoadJSON[namespaces.Namespace](path)
}
