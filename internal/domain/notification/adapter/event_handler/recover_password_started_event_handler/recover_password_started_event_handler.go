package recover_password_started_event_handler

import (
	"encoding/json"
	"getfund-api-v2/internal/domain/notification/core/usecase/send_recover_password_mail"
	"getfund-api-v2/internal/shared/service/cache_service"
	"getfund-api-v2/pkg/bus"
	logger "getfund-api-v2/pkg/log"
)

type recoverPasswordStartedEventHandler struct {
	logger                  logger.Logger
	sendRecoverPasswordMail send_recover_password_mail.UseCase
	cache                   cache_service.Cache
}

func New(sendRecoverPasswordMail send_recover_password_mail.UseCase, cache cache_service.Cache) bus.Handler {
	return &recoverPasswordStartedEventHandler{
		logger:                  *logger.New("recoverPasswordStartedEventHandler"),
		sendRecoverPasswordMail: sendRecoverPasswordMail,
		cache:                   cache,
	}
}

func (h *recoverPasswordStartedEventHandler) Handle(event bus.Event) {
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
