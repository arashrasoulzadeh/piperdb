// Package persist
// Copyright © 2025 Arash Rasoulzadeh <arashrasoulzadeh@gmail.com>
package persist

import (
	"github.com/arashrasoulzadeh/piperdb/src/models/helpers"
	"os"
	"testing"

	"github.com/arashrasoulzadeh/piperdb/src/config"
	"github.com/arashrasoulzadeh/piperdb/src/models/namespaces"
	"github.com/arashrasoulzadeh/piperdb/src/models/pipeline"
	"github.com/arashrasoulzadeh/piperdb/src/persist"

	"github.com/stretchr/testify/assert"
)

func TestPersistNameSpace(t *testing.T) {
	path := "./temp"
	config.NewConfig(config.WithInstancePath(path))
	err := os.RemoveAll(path)
	assert.NoError(t, err)

	ns := namespaces.Namespace{}

	t.Run("Create And Persist", func(t *testing.T) {
		err := ns.SetName("sample_ns")
		assert.NoError(t, err)

		//create new test pipeline
		pipelines := pipeline.Pipelines{}
		pipelineToAppend := pipeline.Pipeline{
			NamespaceName: ns.Name,
		}

		pipelineToAppend, errCreatePipeline := ns.CreateNewPipeline([]byte("test"))
		assert.NoError(t, errCreatePipeline)
		pipelineToAppend.Id = "sample_pipeline"
		pipelines = append(pipelines, pipelineToAppend)
		ns.SetPipelines(pipelines)

		err = persist.SaveModel(ns, ns.Path())
		assert.NoError(t, err)
	})

	t.Run("Check if namespace meta file is valid", func(t *testing.T) {
		metaFilePath := "./temp/sample_ns/meta.json"
		assert.FileExists(t, metaFilePath)

		model, err := helpers.LoadNamespaceFromMeta(metaFilePath)
		assert.NoError(t, err)
		assert.Equal(t, model.Name, "sample_ns")
	})

	t.Run("check if pipeline file is valid", func(t *testing.T) {
		metaFilePath := "./temp/sample_ns/meta.json"
		assert.FileExists(t, metaFilePath)

		model, err := helpers.LoadNamespaceFromMeta(metaFilePath)
		assert.NoError(t, err)
		assert.Equal(t, model.Name, "sample_ns")
	})

	t.Run("re-persist", func(t *testing.T) {
		err = persist.SaveModel(ns, ns.Path())
		assert.NoError(t, err)
	})

	t.Run("re-check if pipeline file is valid", func(t *testing.T) {
		metaFilePath := "./temp/sample_ns/meta.json"
		assert.FileExists(t, metaFilePath)

		model, err := helpers.LoadNamespaceFromMeta(metaFilePath)
		assert.NoError(t, err)
		assert.Equal(t, model.Name, "sample_ns")
	})
}
