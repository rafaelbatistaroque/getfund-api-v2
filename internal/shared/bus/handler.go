package shared_bus

type Handler interface {
	// Handle handles an event.
	Handle(event Event)
}
