package observer_events

import (
	"errors"
	"sync"
)

type Subscriber struct {
	// [event => [listener.handle]]
	EventListeners sync.Map

	mutex sync.Mutex
}

// A method to bind an Event and a Listener
// If a listener is already present in the event, it is added to the queue
func (s *Subscriber) Subscriber(event interface{}, listener interface{}) error {
	ok := checkIsEvent(event)

	if !ok {
		return errors.New("event must be implement IBaseEvent")
	}

	ok = checkIsListener(listener)

	if !ok {
		return errors.New("listener must be implement IListener")
	}

	listeners, ok := s.EventListeners.Load(event.(IBaseEvent).GetEventName())

	var listenersList []IListener

	if ok {
		listenersList = append(listeners.([]IListener), listener.(IListener))
	} else {
		listenersList = append(listenersList, listener.(IListener))
	}

	s.EventListeners.Store(event.(IBaseEvent).GetEventName(), listenersList)

	return nil
}

func (s *Subscriber) DeclareSubscriber(eventName string, listener interface{}) {

	listeners, ok := s.EventListeners.Load(eventName)

	var listenersList []IListener

	if ok {
		listenersList = append(listeners.([]IListener), listener.(IListener))
	} else {
		listenersList = append(listenersList, listener.(IListener))
	}

	s.EventListeners.Store(eventName, listenersList)
}

// Triggering event
func (s *Subscriber) Fire(event interface{}) error {
	has, err := s.HasEvents(event)

	if !has {
		return err
	}

	listeners, ok := s.EventListeners.Load(event.(IBaseEvent).GetEventName())

	if ok {
		wg := sync.WaitGroup{}
		s.mutex.Lock()
		for _, listener := range listeners.([]IListener) {
			wg.Add(1)
			go func(handler IListener) {
				defer wg.Done()
				handler.Handle(event)
			}(listener)
		}
		s.mutex.Unlock()
		wg.Wait()
	} else {
		return errors.New(" There are no listeners ")
	}

	return nil
}

// Triggers the event as a block
func (s *Subscriber) FireBlock(event interface{}) error {
	has, err := s.HasEvents(event)

	if !has {
		return err
	}

	listeners, ok := s.EventListeners.Load(event.(IBaseEvent).GetEventName())

	if ok {
		s.mutex.Lock()
		for _, listener := range listeners.([]IListener) {
			listener.Handle(event)
		}
		s.mutex.Unlock()
	} else {
		return errors.New(" There are no listeners ")
	}

	return nil
}

// Clear all events
func (s *Subscriber) ClearEvents() {
	s.EventListeners = sync.Map{}
}

// Clear a specific event
func (s *Subscriber) ClearEvent(event interface{}) error {
	has, err := s.HasEvents(event)

	if !has {
		return err
	}

	s.EventListeners.Delete(event.(IBaseEvent).GetEventName())

	return nil
}

// Return all event's name
func (s *Subscriber) EventNames() []string {
	events := make([]string, 0)
	s.EventListeners.Range(func(k, v interface{}) bool {
		events = append(events, k.(string))
		return true
	})

	return events
}

// Is there an event
func (s *Subscriber) HasEvents(event interface{}) (bool, error) {
	is := checkIsEvent(event)

	if !is {
		return false, errors.New("event must be implement IBaseEvent")
	}

	_, ok := s.EventListeners.Load(event.(IBaseEvent).GetEventName())

	return ok, nil
}

// Return Amount of events
func (s *Subscriber) EventCount() int {
	i := 0

	s.EventListeners.Range(func(k, v interface{}) bool {
		i++
		return true
	})

	return i
}

// Returns the number of listeners for a specific event
func (s *Subscriber) EventListenerCount(event interface{}) (int, error) {
	has, err := s.HasEvents(event)

	if !has {
		return 0, err
	}

	listeners, _ := s.EventListeners.Load(event.(IBaseEvent).GetEventName())

	return len(listeners.([]IListener)), nil
}

// Instantiate a SUBSCRIBE object
func NewSubscriber() *Subscriber {
	return new(Subscriber)
}

// Check if the event is valid
func checkIsEvent(event interface{}) bool {
	_, ok := event.(IBaseEvent)

	return ok
}

// Check if the listener is valid
func checkIsListener(listener interface{}) bool {
	_, ok := listener.(IListener)

	return ok
}
