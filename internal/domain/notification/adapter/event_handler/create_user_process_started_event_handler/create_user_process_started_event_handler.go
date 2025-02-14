package create_user_process_started_event_handler

import (
	"encoding/json"
	"getfund-api-v2/internal/domain/notification/core/usecase/send_activation_account_mail"
	"getfund-api-v2/internal/shared/service/cache_service"
	"getfund-api-v2/pkg/bus"
	logger "getfund-api-v2/pkg/log"
)

type createUserProcessStartedEventHandler struct {
	logger                    logger.Logger
	cache                     cache_service.Cache
	sendActivationAccountMail send_activation_account_mail.UseCase
}

func New(cache cache_service.Cache, sendActivationAccountMail send_activation_account_mail.UseCase) bus.Handler {
	return &createUserProcessStartedEventHandler{
		logger:                    *logger.New("createUserProcessStartedEventHandler"),
		cache:                     cache,
		sendActivationAccountMail: sendActivationAccountMail,
	}
}

type CreateUserProcessPayload struct {
	ActivationCode string `json:"activation_code"`
	ActivationLink string `json:"activation_link"`
}

func (h *createUserProcessStartedEventHandler) Handle(event bus.Event) {
	var payload CreateUserProcessPayload
	if err := json.Unmarshal(event.GetPayload(), &payload); err != nil {
		h.logger.Error("IsOk: False | get payload failed")
		return
	}

	resultCache, err := h.cache.Get(payload.ActivationCode)
	if err != nil {
		h.logger.Error("IsOk: False | get cache failed")
		return
	}

	var input = &send_activation_account_mail.Input{
		ActivationLink: payload.ActivationLink,
	}
	if err := json.Unmarshal([]byte(resultCache), input); err != nil {
		h.logger.Errorf("IsOk: False | %v", err)
		return
	}

	h.sendActivationAccountMail.Execute(input)
}
