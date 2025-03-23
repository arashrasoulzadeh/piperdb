// Package namespaces /*
// Copyright © 2025 Arash Rasoulzadeh <arashrasoulzadeh@gmail.com>
package namespaces

import (
	"errors"
	"fmt"
	"github.com/arashrasoulzadeh/piperdb/src/models/pipeline"
)

func (namespace *Namespace) SetName(name string) error {
	if namespace.Name != "" {
		return errors.New("namespace name already set")
	}
	namespace.Name = name
	return nil
}
func (namespace *Namespace) SetPipelines(pipelines pipeline.Pipelines) {
	namespace.Pipelines = pipelines
}

func (namespace *Namespace) GetPipelines() []pipeline.Pipeline {
	return namespace.Pipelines
}
func (namespace *Namespace) CreateNewPipeline(payload []byte) (pipeline.Pipeline, error) {
	p := pipeline.Pipeline{
		NamespaceName: namespace.Name,
		Payload:       payload,
	}
	p.AssignNewUUID()

	return p, nil
}
func (namespace *Namespace) Path() string {
	return fmt.Sprintf("%s/meta.json", namespace.Name)
}
