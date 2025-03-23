// Package pipeline
// Copyright © 2025 Arash Rasoulzadeh <arashrasoulzadeh@gmail.com>
package pipeline

import "github.com/arashrasoulzadeh/piperdb/src/models/journey"

type Pipeline struct {
	Id            string `json:"id"`
	Payload       []byte `json:"payload"`
	Journeys      journey.Journeys
	NamespaceName string `json:"namespace_name"`
}

type Pipelines []Pipeline
