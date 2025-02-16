package activate_user

import "getfund-api-v2/pkg/bus"

type ActivateUserConfirmedEvent struct {
	bus.EventBase
}

func (e *ActivateUserConfirmedEvent) GetName() string {
	return "ActivateUserConfirmedEvent"
}
