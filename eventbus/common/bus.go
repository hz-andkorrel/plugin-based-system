package common

// The Bus structure represents the event bus.
// Plugins can subscribe to events and publish events to notify subscribers.
type Bus interface {
	Subscribe(eventName string, handler Handler)
	Publish(event Event)
}
