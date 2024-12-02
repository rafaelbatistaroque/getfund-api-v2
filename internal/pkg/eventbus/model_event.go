package eventbus

type ModelEvent struct{}

func (e *ModelEvent) GetName() string {
	return "ModelEvent"
}
