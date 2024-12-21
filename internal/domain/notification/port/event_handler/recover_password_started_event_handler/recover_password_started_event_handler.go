package recover_password_started_event_handler

import (
	"encoding/json"
	"getfund-api-v2/internal/domain/notification/adapter/usecase/send_recover_password_mail"
	"getfund-api-v2/pkg/eventbus"
	logger "getfund-api-v2/pkg/log"
)

type recoverPasswordStartedEventHandler struct {
	logger                  logger.Logger
	sendRecoverPasswordMail send_recover_password_mail.UseCase
}

func New(sendRecoverPasswordMail send_recover_password_mail.UseCase) eventbus.Handler {
	return &recoverPasswordStartedEventHandler{
		sendRecoverPasswordMail: sendRecoverPasswordMail,
		logger:                  *logger.New("recoverPasswordStartedEventHandler"),
	}
}

func (h *recoverPasswordStartedEventHandler) Handle(event eventbus.Event) {
	var input send_recover_password_mail.Input

	err := json.Unmarshal(event.GetPayload(), &input.KeyCache)
	if err != nil || input.KeyCache == "" {
		panic("Unmarshal failed")
	}

	//WIP: Handler
	//TODO: recover data cached by key received
	//TODO: build a email template with params to replace
	//TODO: replace specific
	//TODO: send email
}
