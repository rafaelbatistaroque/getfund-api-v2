package auth_contract

import model "getfund-api-v2/internal/domain/auth/model"

type UserRepository interface {
	GetByUserName(username string) (*model.UserModel, error)
}
