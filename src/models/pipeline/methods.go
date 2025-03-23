// Package pipeline /*
// Copyright © 2025 Arash Rasoulzadeh <arashrasoulzadeh@gmail.com>
package pipeline

import (
	"encoding/json"
	"github.com/arashrasoulzadeh/piperdb/src/models/journey"
	"github.com/google/uuid"
)

func (p *Pipeline) AssignNewUUID() {
	p.Id = uuid.New().String()
}

func (p *Pipeline) SetID(id string) {
	p.Id = id
}

func (p *Pipeline) Json() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Pipeline) CreateJourney(j *journey.Journey) (*journey.Journey, error) {
	j.PipelineId = p.Id
	j.MarkAsPending()
	j.NamespaceName = p.NamespaceName
	j.AssignNewUUID()
	p.Journeys = append(p.Journeys, *j) // Store the journey in the namespace
	return j, nil
}
