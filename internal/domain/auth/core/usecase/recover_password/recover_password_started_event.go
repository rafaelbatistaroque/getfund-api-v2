package recover_password

import "getfund-api-v2/pkg/bus"

type RecoverPasswordStartedEvent struct {
	bus.EventBase
}

func (e *RecoverPasswordStartedEvent) GetName() string {
	return "RecoverPasswordStartedEvent"
}
