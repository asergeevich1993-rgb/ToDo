package handlers

import (
	"errors"
	"time"
)

type TaskDTO struct {
	Title      string `json:"title"`
	Descrition string `json:"description"`
}

func (t *TaskDTO) ValidateForCreate() error {
	if t.Title == "" {
		return errors.New("title is empty")
	}
	if t.Descrition == "" {
		return errors.New("description is empty")
	}

	return nil

}

type ErrorDTO struct {
	Err  string    `json:"err"`
	Time time.Time `json:"time"`
}
