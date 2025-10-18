package event

import shared_bus "getfund-api-v2/internal/shared/bus"

const RECOVER_PASSWORD_STARTED = "recover.password.started"

// RecoverPasswordStartedEvent is dispatched when a user starts the password recovery process.
// The payload is the key to retrieve the recovery data from cache.
type RecoverPasswordStartedEvent struct {
	shared_bus.EventBase
}

func (e *RecoverPasswordStartedEvent) GetName() string {
	return RECOVER_PASSWORD_STARTED
}
