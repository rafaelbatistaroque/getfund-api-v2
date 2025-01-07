package signin_mapper

import (
	"encoding/json"
	model "getfund-api-v2/internal/domain/auth/core/model"
	"getfund-api-v2/internal/domain/auth/core/usecase/signin"
	"getfund-api-v2/internal/shared/contract/settings"
	"getfund-api-v2/internal/shared/security"
)

type SigninMapper interface {
	ToOutput(token string, session *model.SessionModel) *signin.Output
	SessionToString(session *model.SessionModel) (string, error)
	ToSessionModel(user *model.UserModel) *model.SessionModel
}

type signinMapper struct {
	hasher   security.Hasher
	settings settings.ApplicationSettings
}

// Constructor
func New(hasher security.Hasher, settings settings.ApplicationSettings) SigninMapper {
	return &signinMapper{
		hasher:   hasher,
		settings: settings,
	}
}

func (m *signinMapper) ToOutput(token string, session *model.SessionModel) *signin.Output {
	return &signin.SigninOutput{
		Token: token,
		Session: signin.SessionOutput{
			ID:        session.ID,
			FirstName: session.FirstName,
			IsAdmin:   session.IsAdmin == 1,
		},
	}
}

func (m *signinMapper) SessionToString(session *model.SessionModel) (string, error) {
	sessionSerialized, err := json.Marshal(session)
	if err != nil {
		return "", err
	}

	return string(sessionSerialized), nil
}

func (m *signinMapper) ToSessionModel(user *model.UserModel) *model.SessionModel {
	return &model.SessionModel{
		ID:        user.Id,
		FirstName: m.hasher.DecryptMerged(user.FirstName, m.settings.GetSecretKey()),
		IsAdmin:   user.IsAdmin,
	}
}
