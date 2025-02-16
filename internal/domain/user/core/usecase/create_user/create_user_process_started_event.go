package create_user

import "getfund-api-v2/pkg/bus"

type CreateUserProcessStartedEvent struct {
	bus.EventBase
}

func (e *CreateUserProcessStartedEvent) GetName() string {
	return "CreateUserProcessStartedEvent"
}
