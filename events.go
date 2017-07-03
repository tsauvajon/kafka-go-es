package main

import "github.com/satori/go.uuid"

// Event : event
type Event struct {
	AccID string
	Type  string
}

// CreateEvent : create event
type CreateEvent struct {
	Event
	AccName string
}

// NewCreateAccountEvent : create a new account event
func NewCreateAccountEvent(name string) CreateEvent {
	event := new(CreateEvent)
	event.Type = "CreateEvent"
	event.AccID = uuid.NewV4().String()
	event.AccName = name
	return *event
}
