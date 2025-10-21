package recover_password_started_event_handler

import (
	"encoding/json"
	"getfund-api-v2/internal/domain/notification/core/usecase/send_recover_password_mail"

	shared_bus "getfund-api-v2/internal/shared/bus"
	"getfund-api-v2/internal/shared/cache"
	shared_logger "getfund-api-v2/internal/shared/log"
)

type recoverPasswordStartedEventHandler struct {
	logger                  shared_logger.Logger
	sendRecoverPasswordMail send_recover_password_mail.UseCase
	cache                   cache.Service
}

func New(sendRecoverPasswordMail send_recover_password_mail.UseCase, cache cache.Service) shared_bus.Handler {
	return &recoverPasswordStartedEventHandler{
		logger:                  *shared_logger.New("recoverPasswordStartedEventHandler"),
		sendRecoverPasswordMail: sendRecoverPasswordMail,
		cache:                   cache,
	}
}

func (h *recoverPasswordStartedEventHandler) Handle(event shared_bus.Event) {
	payload := string(event.GetPayload())
	if payload == "" {
		h.logger.Error("IsOk: False | get payload failed")
		return
	}

	resultCache, errCache := h.cache.Get(payload)
	if errCache != nil {
		h.logger.Error("IsOk: False | get cache failed")
		return
	}

	var input = &send_recover_password_mail.Input{}
	if err := json.Unmarshal([]byte(resultCache), input); err != nil {
		h.logger.Errorf("IsOk: False | %v", err)
		return
	}

	success, err := h.sendRecoverPasswordMail.Execute(input)
	if err != nil {
		h.logger.Errorf("IsOk: False | Code: %d | Message %v", err.Code, err.Message)
		return
	}

	h.logger.Infof("IsOk: True | %v", success.Messagem)
}
