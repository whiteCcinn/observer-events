package observer_events

import (
	"errors"
)

type IBaseEvent interface {
	GetEventName() string
}

type IListener interface {
	Handle(event interface{})
}

type ISubscriber interface {
	// A method to bind an Event and a Listener
	// If a listener is already present in the event, it is added to the queue
	Subscriber(event interface{}, listener interface{}) error
	DeclareSubscriber(eventName string, listener interface{}) error
	// Triggering event
	Fire(event interface{}) error
	// Triggers the event as a block
	FireBlock(event interface{}) error
	// Clear all events
	ClearEvents()
	// Clear a specific event
	ClearEvent(event interface{}) error
	// Return all event's name
	EventNames() []string
	// Is there an event
	HasEvents(event interface{}) (bool, error)
	// Return Amount of events
	EventCount() int
	// Returns the number of listeners for a specific event
	EventListenerCount(event interface{}) (int, error)
}

var globalSubscriber = NewSubscriber()

func Subscribe(event interface{}, listener interface{}) error {
	return globalSubscriber.Subscriber(event, listener)
}

func Event(event interface{}) error {
	ok := checkIsEvent(event)

	if !ok {
		return errors.New("event must be implement IBaseEvent")
	}

	return globalSubscriber.Fire(event)
}

func EventBlock(event interface{}) error {
	ok := checkIsEvent(event)

	if !ok {
		return errors.New("event must be implement IBaseEvent")
	}

	return globalSubscriber.FireBlock(event)
}

func ClearEvents() {
	globalSubscriber.ClearEvents()
}

func ClearEvent(event interface{}) error {
	return globalSubscriber.ClearEvent(event)
}

func EventNames() []string {
	return globalSubscriber.EventNames()
}

func HasEvents(event interface{}) (bool, error) {
	return globalSubscriber.HasEvents(event)
}

func EventCount() int {
	return globalSubscriber.EventCount()
}

func EventListenerCount(event interface{}) (int, error) {
	return globalSubscriber.EventListenerCount(event)
}

func DeclareSubscriber(eventName string, listener interface{}) error {
	return globalSubscriber.DeclareSubscriber(eventName, listener)
}
