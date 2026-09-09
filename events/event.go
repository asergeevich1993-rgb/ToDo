package event

import "time"

type Events struct {
	Descrition string
	Errors     string
	CreatedAt  time.Time
}

func NewEvent(description string, errors string) Events {
	return Events{
		Descrition: description,
		Errors:     errors,
		CreatedAt:  time.Now(),
	}
}
