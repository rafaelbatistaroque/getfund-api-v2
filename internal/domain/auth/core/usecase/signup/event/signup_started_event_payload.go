package event

import shared_bus "getfund-api-v2/internal/shared/bus"

const SIGNUP_STARTED = "signup.started"

// SignupStartedEvent is dispatched when a user starts the signup process.
type SignupStartedEvent struct {
	shared_bus.EventBase
}

func (e *SignupStartedEvent) GetName() string {
	return SIGNUP_STARTED
}

// SignupStartedPayload is the data contract for the SignupStartedEvent.
type SignupStartedPayload struct {
	FirstName      string `json:"first_name"`
	Email          string `json:"username"`
	ActivationCode string `json:"activation_code"`
	ActivationLink string `json:"activation_link"`
}
