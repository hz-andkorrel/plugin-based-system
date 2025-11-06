package domain

import (
	common "eventbus/common"
)

// The Bus structure represents the event bus.
// Plugins can subscribe to events and publish events to notify subscribers.
type Bus struct {
	subscribers map[string][]common.Handler
}

// NewBus creates a new instance of the event bus.
// An empty map of subscribers is initialized to keep track of event subscriptions.
func NewBus() common.Bus {
	return &Bus{
		subscribers: make(map[string][]common.Handler),
	}
}

// Subscribe allows a plugin to subscribe to a specific event by providing an event name and a handler function.
func (bus *Bus) Subscribe(eventName string, handler common.Handler) {
	bus.subscribers[eventName] = append(bus.subscribers[eventName], handler)
}

// Publish allows a plugin to publish an event to all subscribers of that event.
func (bus *Bus) Publish(event common.Event) {
	for _, handler := range bus.subscribers[event.Name] {
		handler(event)
	}
}
