package create_user

import "getfund-api-v2/pkg/bus"

type CreateUserProcessWithCouponStartedEvent struct {
	bus.EventBase
}

func (e *CreateUserProcessWithCouponStartedEvent) GetName() string {
	return "CreateUserProcessWithCouponStartedEvent"
}
