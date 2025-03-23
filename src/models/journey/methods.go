// Package journey /*
// Copyright © 2025 Arash Rasoulzadeh <arashrasoulzadeh@gmail.com>
package journey

import (
	"errors"
	"fmt"
	"github.com/arashrasoulzadeh/piperdb/src/persist"
	"github.com/google/uuid"
	"time"
)

func (j *Journey) Persist() error {
	return persist.SaveModel(j, j.Path())
}

func (j *Journey) Path() string {
	return fmt.Sprintf("%s/%s/journeys/%s.json", j.NamespaceName, j.PipelineId, j.Id)
}

func (j *Journey) MarkAsPending() (err error) {
	if j.Status == Pending {
		return errors.New("already pending")
	}
	j.Status = Pending
	err = j.addStatusHistory(Pending)
	return err
}
func (j *Journey) MarkAsRunning() (err error) {
	if j.Status == Running {
		return errors.New("already running")
	}
	j.Status = Running
	err = j.addStatusHistory(Running)
	return err
}
func (j *Journey) MarkAsComplete() (err error) {
	if j.Status == Complete {
		return errors.New("already complete")
	}
	j.Status = Complete
	err = j.addStatusHistory(Complete)
	return err
}
func (j *Journey) MarkAsFailed() (err error) {
	if j.Status == Failed {
		return errors.New("already failed")
	}
	j.Status = Failed
	err = j.addStatusHistory(Failed)
	return err
}

func (j *Journey) addStatusHistory(status Status) error {
	j.StatusHistory = append(j.StatusHistory, StatusHistory{
		Status: status,
		Date:   time.Now(),
	})
	return nil
}

func (j *Journey) AssignNewUUID() {
	j.Id = uuid.New().String()
}
