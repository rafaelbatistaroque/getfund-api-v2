package activate_user_confirmed_event_handler

import (
	"encoding/json"
	"getfund-api-v2/internal/domain/auth/core/usecase/signin"
	"getfund-api-v2/internal/shared/app_constant"
	"getfund-api-v2/pkg/bus"
	logger "getfund-api-v2/pkg/log"
)

type activateUserConfirmedEventHandler struct {
	logger logger.Logger
	signin signin.UseCase
}

func New(signin signin.UseCase) bus.Handler {
	return &activateUserConfirmedEventHandler{
		logger: *logger.New("activateUserConfirmedEventHandler"),
		signin: signin,
	}
}

type ActivateUserConfirmedPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *activateUserConfirmedEventHandler) Handle(event bus.Event) {
	var payload = ActivateUserConfirmedPayload{}
	if err := json.Unmarshal(event.GetPayload(), &payload); err != nil {
		h.logger.Error("IsOk: False | get payload failed")
		return
	}

	success, errSignin := h.signin.Execute(&signin.Input{
		UserName: payload.Username,
		Password: payload.Password,
	})

	if errSignin != nil {
		event.DefineResponse(app_constant.EMPTYB)
		h.logger.Errorf("IsOk: False | Code: %d | Message %v", errSignin.Code, errSignin.Message)
		return
	}

	result, err := json.Marshal(success)
	if err != nil {
		event.DefineResponse(app_constant.EMPTYB)
		h.logger.Error("IsOk: False | marshal failed")
		return
	}

	event.DefineResponse(result)
	h.logger.Infof("IsOk: True")
}
