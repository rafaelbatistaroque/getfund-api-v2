package recoverpasswordfixture_test

import (
	"fmt"
	"getfund-api-v2/internal/domain/auth/usecase/recoverpassword"
	sut "getfund-api-v2/internal/domain/auth/usecase/recoverpassword/application"
	"getfund-api-v2/internal/shared/resultapp"
	"getfund-api-v2/pkg/inputvalidation"
	"getfund-api-v2/test/helper/cachespy"
	"getfund-api-v2/test/helper/codespy"
	"getfund-api-v2/test/helper/eventbusspy"
	"getfund-api-v2/test/helper/securityspy"
	"getfund-api-v2/test/helper/settingsspy"
	"getfund-api-v2/test/helper/userrepositoryspy"
)

type RecoverPasswordFixture struct {
	HasherSpy   *securityspy.HasherSpy
	SettingsSpy *settingsspy.ApplicationSettingsSpy
	UserRepoSpy *userrepositoryspy.UserRepositorySpy
	CodeSpy     *codespy.CodeSpy
	CacheSpy    *cachespy.RedisCacheSpy
	EventBusSpy *eventbusspy.EventBusSpy
}

func NewSut() (recoverpassword.UseCase, *RecoverPasswordFixture) {
	hasherSpy := securityspy.New()
	settingsSpy := settingsspy.New()
	userRepoSpy := userrepositoryspy.New()
	codeSpy := codespy.New()
	cacheSpy := cachespy.New()
	eventBusSpy := eventbusspy.New()

	return sut.New(hasherSpy, settingsSpy, userRepoSpy, codeSpy, cacheSpy, eventBusSpy),
		&RecoverPasswordFixture{
			HasherSpy:   hasherSpy,
			SettingsSpy: settingsSpy,
			UserRepoSpy: userRepoSpy,
			CodeSpy:     codeSpy,
			CacheSpy:    cacheSpy,
			EventBusSpy: eventBusSpy,
		}
}

func GetInvalidInputWithError() (*recoverpassword.Input, *resultapp.ApplicationError) {
	return &recoverpassword.Input{Username: ""},
		resultapp.New(resultapp.BAD_REQUEST_CODE, fmt.Errorf(inputvalidation.Err_Msg_PARAMETER_NOT_EMPTY.Error(), "Username"))
}

func GetValidInput() *recoverpassword.Input {
	return &recoverpassword.Input{Username: "fake-username"}
}
