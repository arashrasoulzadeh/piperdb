package persist

import (
	"fmt"
	"github.com/arashrasoulzadeh/piperdb/src/config"
	"github.com/arashrasoulzadeh/piperdb/src/models/helpers"
	"github.com/arashrasoulzadeh/piperdb/src/models/namespaces"
	"github.com/arashrasoulzadeh/piperdb/src/models/pipeline"
	"github.com/arashrasoulzadeh/piperdb/src/persist"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"os"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestPersistBenchmark(t *testing.T) {
	t.Run("create 100 ns", func(t *testing.T) {
		RunBenchmark(t, 100)
		RunBenchmark(t, 1000)
		clear(t)
	})
}

func RunBenchmark(t *testing.T, total int) {
	workerCount := runtime.NumCPU()
	fmt.Printf("Running benchmark with %d workers\n", workerCount)
	jobs := make(chan int, total)
	var wg sync.WaitGroup

	clear(t)
	start := time.Now()

	for w := 0; w < workerCount; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range jobs {
				processTask(t)
			}
		}()
	}

	for i := 0; i < total; i++ {
		jobs <- i
	}
	close(jobs)

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("#%d ==> elapsed: %s\n", total, elapsed)
}

func processTask(t *testing.T) {
	name := uuid.New().String()

	ns := namespaces.Namespace{}
	err := ns.SetName(name)
	assert.NoError(t, err)

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

	metaFilePath := "./temp/" + name + "/meta.json"
	assert.FileExists(t, metaFilePath)

	model, err := helpers.LoadNamespaceFromMeta(metaFilePath)
	assert.NoError(t, err)
	assert.Equal(t, model.Name, name)
}

func clear(t *testing.T) {
	path := "./temp"
	config.NewConfig(config.WithInstancePath(path))
	err := os.RemoveAll(path)
	assert.NoError(t, err)
}
