# observer-events
🌈This is an observer mode based event component, similar to the events module in PHP's laravel framework

# How to install

```shell
go get github.com/whiteCcinn/observer-events
```

# How To Use

## Event

We need to implement our own event, which needs to implement the interface `IBaseEvent`
And you need to implement the method `GetEventName()`, which requires the output event name

For example:

```go
package main

type MyEvent struct {
	eventName string
	name      string
	age       int
}

func (e *MyEvent) GetEventName() string {
	return e.eventName
}
```

## Listener

When we want to observe an event through a Listener to trigger the observer's behavior, we need to define our own Listener. 
This requires us to implement the interface `IListener`, which has an interface method `Handle(Event interface{})`, which will receive the event object
you can display the transformation of your custom events through the`.(Event)` syntax

For example:

```go
package main

import "fmt"

type MyEvent struct {
	eventName string
	name      string
	age       int
}

func (e *MyEvent) GetEventName() string {
	return e.eventName
}

type MyListener struct {
}

func (l *MyListener) Handle(event interface{}) {
	if myEvent, ok := event.(*MyEvent); ok {
		fmt.Println(fmt.Sprintf("MyListener>> name:%s, age:%d", myEvent.name, myEvent.age))
	}
}

type MyListener2 struct {
}

func (l *MyListener2) Handle(event interface{}) {
	if myEvent, ok := event.(*MyEvent); ok {
		fmt.Println(fmt.Sprintf("MyListener2>> name:%s, age:%d", myEvent.name, myEvent.age))
	}
}
```

## Subscriber

The subscriber will be the core processing structure of our entire behavior
It provides a number of apis。

```go
type ISubscriber interface {
	// A method to bind an Event and a Listener
	// If a listener is already present in the event, it is added to the queue
	Subscriber(event interface{}, listener interface{}) error
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
```

You can instantiate your own subscribers using the global subscriber or the `NewSubscriber()`

Global method `Event(Event Interface {}) = Fire(Event)`, `EventBlock(Event Interface {}) = FireBlock(Event)`

```go
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
```