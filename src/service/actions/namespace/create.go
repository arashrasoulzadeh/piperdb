// Package namespace /*
// Copyright © 2025 Arash Rasoulzadeh <arashrasoulzadeh@gmail.com>
package namespace

import (
	"fmt"
	"github.com/arashrasoulzadeh/piperdb/src/models/namespaces"
	"github.com/arashrasoulzadeh/piperdb/src/persist"
)

func Create(name string) {
	ns := namespaces.Namespace{
		Name: name,
	}
	persist.SaveModel(ns, ns.Path())
	fmt.Println(fmt.Sprintf("created namespace %s", name))
}
