// Package journey /*
// Copyright © 2025 Arash Rasoulzadeh <arashrasoulzadeh@gmail.com>
package journey

import "time"

type Status string

const (
	Pending  Status = "pending"
	Running  Status = "running"
	Complete Status = "complete"
	Failed   Status = "failed"
)

type StatusHistory struct {
	Status Status    `json:"status"`
	Date   time.Time `json:"date"`
}

type Journey struct {
	Id            string          `json:"id"`
	NamespaceName string          `json:"namespace_name"`
	PipelineId    string          `json:"pid"`
	Payload       []byte          `json:"payload"`
	Owner         string          `json:"called_by"`
	Status        Status          `json:"status"`
	StatusHistory []StatusHistory `json:"status_history"`
}

type Journeys []Journey
