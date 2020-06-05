package observer_events_test

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	observerEvents "github.com/whiteCcinn/observer-events"
	"testing"
)

type MyEvent struct {
	eventName string
	name      string
	age       int
}

func (e MyEvent) GetEventName() string {
	return e.eventName
}

type MyListener struct {
}

func (l MyListener) Handle(event interface{}) {
	if myEvent, ok := event.(MyEvent); ok {
		fmt.Println(fmt.Sprintf("MyListener>> name:%s, age:%d", myEvent.name, myEvent.age))
	}
}

type MyListener2 struct {
}

func (l MyListener2) Handle(event interface{}) {
	if myEvent, ok := event.(MyEvent); ok {
		fmt.Println(fmt.Sprintf("MyListener2>> name:%s, age:%d", myEvent.name, myEvent.age))
	}
}

func TestEvent(t *testing.T) {
	event := MyEvent{"test_event", "ccinn", 18}
	err := observerEvents.Event(event)

	assert.NotNil(t, err)
}

func TestSubscribe(t *testing.T) {

	event := MyEvent{"test_event", "ccinn", 18}

	listener := MyListener{}

	err := observerEvents.Subscribe(event, listener)

	assert.Nil(t, err)

	listener2 := MyListener2{}

	err = observerEvents.Subscribe(event, listener2)

	assert.Nil(t, err)

	err = observerEvents.Event(event)

	assert.Nil(t, err)

	err = observerEvents.EventBlock(event)

	assert.Nil(t, err)
}

func TestClearEvents(t *testing.T) {

	event := MyEvent{"test_event", "ccinn", 18}

	listener := MyListener{}

	err := observerEvents.Subscribe(event, listener)

	assert.Nil(t, err)

	listener2 := MyListener2{}

	err = observerEvents.Subscribe(event, listener2)

	observerEvents.ClearEvents()

	if observerEvents.EventCount() != 0 {
		assert.Error(t, errors.New("clear events after is not empty"))
	}
}

func TestClearEvent(t *testing.T) {

	event := MyEvent{"test_event", "ccinn", 18}

	listener := MyListener{}

	err := observerEvents.Subscribe(event, listener)

	assert.Nil(t, err)

	listener2 := MyListener2{}

	err = observerEvents.Subscribe(event, listener2)

	assert.Nil(t, err)

	err = observerEvents.ClearEvent(event)

	assert.Nil(t, err)
}

func TestEventNames(t *testing.T) {

	event := MyEvent{"test_event", "ccinn", 18}

	listener := MyListener{}

	err := observerEvents.Subscribe(event, listener)

	assert.Nil(t, err)

	listener2 := MyListener2{}

	err = observerEvents.Subscribe(event, listener2)

	assert.Nil(t, err)

	names := observerEvents.EventNames()

	if len(names) != 1 {
		assert.Error(t, errors.New("error"))
	}
}

func TestHasEvents(t *testing.T) {

	event := MyEvent{"test_event", "ccinn", 18}

	listener := MyListener{}

	err := observerEvents.Subscribe(event, listener)

	assert.Nil(t, err)

	listener2 := MyListener2{}

	err = observerEvents.Subscribe(event, listener2)

	assert.Nil(t, err)

	has, err := observerEvents.HasEvents(event)

	assert.Nil(t, err)
	assert.Equal(t, true, has)
}

func TestEventListenerCount(t *testing.T) {

	event := MyEvent{"test_event", "ccinn", 18}

	listener := MyListener{}

	err := observerEvents.Subscribe(event, listener)

	assert.Nil(t, err)

	listener2 := MyListener2{}

	err = observerEvents.Subscribe(event, listener2)

	assert.Nil(t, err)

	count, err := observerEvents.EventListenerCount(event)

	assert.Nil(t, err)
	assert.Equal(t, 2, count)
}

