package send_recover_password_mail_application

import (
	"getfund-api-v2/internal/domain/notification/adapter/usecase/send_recover_password_mail"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/internal/shared/service/cache_service"
)

type sendRecoverPasswordMailApplication struct {
	cacheService cache_service.Cache
	//email service
	//settings
}

func New(cacheService cache_service.Cache) send_recover_password_mail.UseCase {
	return &sendRecoverPasswordMailApplication{
		cacheService: cacheService,
	}
}

func (uc *sendRecoverPasswordMailApplication) Execute(input *send_recover_password_mail.Input) (*send_recover_password_mail.Output, *result_app.ApplicationError) {
	input.Validate()
	if input.IsInvalid() {
		return nil, result_app.New(result_app.BAD_REQUEST_CODE, input.GetErrors())
	}

	_, errCache := uc.cacheService.Get(input.KeyCache)
	if errCache != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, errCache)
	}

	return nil, nil
}

//WIP: Handler
//TODO: recover data cached by key received
//TODO: build a email template with params to replace
//TODO: replace specific
//TODO: send email
