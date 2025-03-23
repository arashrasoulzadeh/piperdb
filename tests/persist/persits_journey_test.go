// Package persist /*
// Copyright © 2025 Arash Rasoulzadeh <arashrasoulzadeh@gmail.com>
package persist

import (
	"github.com/arashrasoulzadeh/piperdb/src/models/helpers"
	"github.com/arashrasoulzadeh/piperdb/src/models/journey"
	"github.com/arashrasoulzadeh/piperdb/src/models/namespaces"
	"github.com/arashrasoulzadeh/piperdb/src/models/pipeline"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPersistJourney(t *testing.T) {

	var ns *namespaces.Namespace
	var p *pipeline.Pipeline
	var j *journey.Journey
	t.Run("be sure that namespace exists and valid before full journey test", func(t *testing.T) {
		metaFilePath := "./temp/sample_ns/meta.json"
		assert.FileExists(t, metaFilePath)
		model, err := helpers.LoadNamespaceFromMeta(metaFilePath)
		assert.NoError(t, err)
		assert.Equal(t, model.Name, "sample_ns")
		ns = model
		p = &ns.GetPipelines()[0]
	})

	t.Run("create new journey model", func(t *testing.T) {
		createdJourney, err := p.CreateJourney(&journey.Journey{})
		j = createdJourney
		assert.NoError(t, err)
		assert.NotNil(t, j)
		assert.Equal(t, createdJourney, j)
	})

	t.Run("persist single journey", func(t *testing.T) {
		errPersist := j.Persist()
		assert.NoError(t, errPersist)
	})
	t.Run("make random history and status", func(t *testing.T) {
		j.MarkAsPending()
		j.MarkAsRunning()
		j.MarkAsComplete()
		j.MarkAsFailed()
		errPersist := j.Persist()
		assert.NoError(t, errPersist)
	})
	t.Run("load journey and validate data", func(t *testing.T) {
		loadedJourney, err := helpers.LoadJourneyFromMeta("./temp/" + j.Path())
		assert.NoError(t, err)
		assert.NotNil(t, loadedJourney)
		assert.Equal(t, j.Status, loadedJourney.Status)
		assert.Equal(t, j.Owner, loadedJourney.Owner)
		assert.Equal(t, j.Payload, loadedJourney.Payload)
		assert.Equal(t, j.NamespaceName, loadedJourney.NamespaceName)
		assert.Equal(t, j.PipelineId, loadedJourney.PipelineId)
		assert.Equal(t, j.StatusHistory[0].Status, loadedJourney.StatusHistory[0].Status)
		assert.Equal(t, j.StatusHistory[1].Status, loadedJourney.StatusHistory[1].Status)
		assert.Equal(t, j.StatusHistory[2].Status, loadedJourney.StatusHistory[2].Status)
		assert.Equal(t, j.StatusHistory[3].Status, loadedJourney.StatusHistory[3].Status)
	})
	t.Run("add new status and persist", func(t *testing.T) {
		j.MarkAsComplete()
		errPersist := j.Persist()
		assert.NoError(t, errPersist)
	})
	t.Run("check update", func(t *testing.T) {
		loadedJourney, err := helpers.LoadJourneyFromMeta("./temp/" + j.Path())
		assert.NoError(t, err)
		assert.NotNil(t, loadedJourney)
		assert.Equal(t, j.Status, loadedJourney.Status)
		assert.Equal(t, j.Owner, loadedJourney.Owner)
		assert.Equal(t, j.Payload, loadedJourney.Payload)
		assert.Equal(t, j.NamespaceName, loadedJourney.NamespaceName)
		assert.Equal(t, j.PipelineId, loadedJourney.PipelineId)
		assert.Equal(t, j.StatusHistory[4].Status, loadedJourney.StatusHistory[4].Status)
	})
}
