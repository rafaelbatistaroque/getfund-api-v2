package bus

type Handler interface {
	Handle(event Event)
}
