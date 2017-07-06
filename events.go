package main

import "github.com/satori/go.uuid"

// Event : event
type Event struct {
	AccID string
	Type  string
}

// CreateEvent : create a new account
type CreateEvent struct {
	Event
	AccName string
}

// DepositEvent : Make a deposit
type DepositEvent struct {
	Event
	Amount int
}

// WithdrawEvent : Withdraw money
type WithdrawEvent struct {
	Event
	Amount int
}

// TransferEvent : Transfer money
type TransferEvent struct {
	Event
	TargetID string
	Amount   int
}

// NewCreateAccountEvent : creates a new CreateEvent
func NewCreateAccountEvent(name string) CreateEvent {
	event := new(CreateEvent)
	event.Type = "CreateEvent"
	event.AccID = uuid.NewV4().String()
	event.AccName = name
	return *event
}

// NewDepositEvent : helper to create a new DepositEvent
func NewDepositEvent(id string, amount int) DepositEvent {
	event := new(DepositEvent)
	event.Type = "DepositEvent"
	event.AccID = id
	event.Amount = amount
	return *event
}

// NewWithdrawEvent : helper to create a new WithdrawEvent
func NewWithdrawEvent(id string, amount int) WithdrawEvent {
	event := new(WithdrawEvent)
	event.Type = "WithdrawEvent"
	event.AccID = id
	event.Amount = amount
	return *event
}

// NewTransferEvent : helper to create a new TransferEvent
func NewTransferEvent(id, targetID string, amount int) TransferEvent {
	event := new(TransferEvent)
	event.Type = "TransferEvent"
	event.AccID = id
	event.TargetID = targetID
	event.Amount = amount
	return *event
}
