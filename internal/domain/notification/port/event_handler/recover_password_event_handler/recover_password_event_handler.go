package recover_password_event_handler

import (
	"encoding/json"
	"fmt"
	"getfund-api-v2/internal/domain/notification/adapter/usecase/send_recover_password_mail"
	"getfund-api-v2/pkg/eventbus"
	logger "getfund-api-v2/pkg/log"
)

type recoverPasswordEventHandler struct {
	logger                  logger.Logger
	sendRecoverPasswordMail send_recover_password_mail.UseCase
}

func New(sendRecoverPasswordMail send_recover_password_mail.UseCase) eventbus.Handler {
	return &recoverPasswordEventHandler{
		sendRecoverPasswordMail: sendRecoverPasswordMail,
		logger:                  *logger.New("RecoverPasswordEventHandler"),
	}
}

// Implementa o Handler genérico para RecoverPasswordStartedEvent
func (h *recoverPasswordEventHandler) Handle(event eventbus.Event) {
	var input send_recover_password_mail.Input
	err := json.Unmarshal(event.GetPayload(), &input.KeyCache)
	if err != nil || input.KeyCache == "" {
		h.logger.Errorf("Unmarshal failed: %v", err)
		panic("Unmarshal failed")
	}

	fmt.Printf("teste %v", input.KeyCache)
	//WIP: Handler
	//TODO: recover data cached by key received
	//TODO: build a email template with params to replace
	//TODO: replace specific
	//TODO: send email
}
