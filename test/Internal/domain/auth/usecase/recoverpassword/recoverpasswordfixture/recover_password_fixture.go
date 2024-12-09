package recoverpasswordfixture_test

import (
	"fmt"
	"getfund-api-v2/internal/domain/auth/usecase/recoverpassword"
	sut "getfund-api-v2/internal/domain/auth/usecase/recoverpassword/application"
	"getfund-api-v2/internal/shared/resultapp"
	"getfund-api-v2/pkg/inputvalidation"
	"getfund-api-v2/test/helper/codespy"
	"getfund-api-v2/test/helper/securityspy"
	"getfund-api-v2/test/helper/settingsspy"
	"getfund-api-v2/test/helper/userrepositoryspy"
)

func NewSut() (recoverpassword.UseCase, *securityspy.HasherSpy, *settingsspy.ApplicationSettingsSpy, *userrepositoryspy.UserRepositorySpy, *codespy.CodeSpy) {
	hasherSpy := securityspy.New()
	settingsSpy := settingsspy.New()
	userRepoSpy := userrepositoryspy.New()
	codeSpy := codespy.New()

	return sut.New(hasherSpy, settingsSpy, userRepoSpy, codeSpy), hasherSpy, settingsSpy, userRepoSpy, codeSpy
}

func GetInvalidInputWithError() (*recoverpassword.Input, *resultapp.ApplicationError) {
	return &recoverpassword.Input{Username: ""},
		resultapp.New(resultapp.BAD_REQUEST_CODE, fmt.Errorf(inputvalidation.Err_Msg_PARAMETER_NOT_EMPTY.Error(), "Username"))
}

func GetValidInput() *recoverpassword.Input {
	return &recoverpassword.Input{Username: "fake-username"}
}
