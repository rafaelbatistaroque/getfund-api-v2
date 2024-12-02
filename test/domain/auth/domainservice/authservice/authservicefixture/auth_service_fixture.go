package authservicefixture

import (
	"errors"
	"getfund-api-v2/internal/domain/auth/domainservice/authservice"
	model "getfund-api-v2/internal/domain/auth/model"
	"getfund-api-v2/test/spyshared/mapperspy/signinmapperspy"
	"getfund-api-v2/test/spyshared/securityspy"
	"getfund-api-v2/test/spyshared/settingsspy"
)

func NewSut() (authservice.AuthService, *settingsspy.ApplicationSettingsSpy, *userRepositorySpy, *securityspy.HasherSpy, *signinmapperspy.SigninMapperSpy) {
	settingsSpy := settingsspy.New()
	userRepositorySpy := &userRepositorySpy{Params: make(map[string]string, 1), ErrorResult: make(map[string]error), SuccessResult: make(map[string]interface{}, 1), CallsCount: make(map[string]int, 1)}
	hasherSpy := securityspy.New()
	mapperSpy := signinmapperspy.New()

	return authservice.New(userRepositorySpy, settingsSpy, hasherSpy, mapperSpy),
		settingsSpy,
		userRepositorySpy,
		hasherSpy,
		mapperSpy
}

type userRepositorySpy struct {
	Params        map[string]string
	CallsCount    map[string]int
	ErrorResult   map[string]error
	SuccessResult map[string]interface{}
}

func (r *userRepositorySpy) GetByUserName(username string) (*model.UserModel, error) {
	r.Params["GetByUserName:username"] = username

	r.CallsCount["GetByUserName"]++

	sucess := r.SuccessResult["GetByUserName"]
	if sucess != nil {
		return sucess.(*model.UserModel), r.ErrorResult["GetByUserName"]
	}
	return nil, r.ErrorResult["GetByUserName"]
}

func (r *userRepositorySpy) DefineError() {
	r.ErrorResult["GetByUserName"] = errors.New("fake-error")
}

func (r *userRepositorySpy) DefineSuccess() {
	r.SuccessResult["GetByUserName"] = &model.UserModel{Password: "fake-password-hashed"}
}
