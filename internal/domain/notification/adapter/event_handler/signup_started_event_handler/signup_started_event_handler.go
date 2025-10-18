package signup_process_started_event_handler

import (
	"encoding/json"
	"getfund-api-v2/internal/domain/notification/core/usecase/send_activation_account_mail"
	shared_bus "getfund-api-v2/internal/shared/bus"
	shared_logger "getfund-api-v2/internal/shared/log"
)

type SignupStartedEventHandler struct {
	logger                    shared_logger.Logger
	sendActivationAccountMail send_activation_account_mail.UseCase
}

func New(sendActivationAccountMail send_activation_account_mail.UseCase) shared_bus.Handler {
	return &SignupStartedEventHandler{
		logger:                    *shared_logger.New("SignupProcessStartedEventHandler"),
		sendActivationAccountMail: sendActivationAccountMail,
	}
}

func (h *SignupStartedEventHandler) Handle(event shared_bus.Event) {
	var payload *send_activation_account_mail.Input
	if err := json.Unmarshal(event.GetPayload(), &payload); err != nil {
		h.logger.Error("IsOk: False | get payload failed")
		return
	}

	success, errUsecase := h.sendActivationAccountMail.Execute(payload)
	if errUsecase != nil {
		h.logger.Errorf("IsOk: False | Code: %d | Message %v", errUsecase.Code, errUsecase.Message)
		return
	}

	h.logger.Infof("IsOk: True | %v", success.Message)
}
