// Package helpers /*
// Copyright © 2025 Arash Rasoulzadeh <arashrasoulzadeh@gmail.com>
package helpers

import (
	"github.com/arashrasoulzadeh/piperdb/src/models/journey"
)

func LoadJourneyFromMeta(path string) (*journey.Journey, error) {
	return LoadJSON[journey.Journey](path)
}
