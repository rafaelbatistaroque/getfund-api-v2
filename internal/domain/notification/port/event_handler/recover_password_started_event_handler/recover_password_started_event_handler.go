package recover_password_started_event_handler

import (
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
	payload := string(event.GetPayload())
	if payload == "" {
		panic("get payload failed")
	}

	input := &send_recover_password_mail.Input{KeyCache: payload}
	success, err := h.sendRecoverPasswordMail.Execute(input)
	if err != nil {
		h.logger.Errorf("IsOk: False | Code: %d | Message: %s", err.Code, err.Message)
		return
	}

	h.logger.Infof("IsOk: True | Message: %s", success.Messagem)

	//WIP: Handler
	//TODO: recover data cached by key received
	//TODO: build a email template with params to replace
	//TODO: replace specific
	//TODO: send email
}
