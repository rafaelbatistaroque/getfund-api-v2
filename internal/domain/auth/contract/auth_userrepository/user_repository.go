package auth_userrepository

import model "getfund-api-v2/internal/domain/auth/authmodel"

type UserRepository interface {
	GetByUserName(username string) (*model.UserModel, error)
}
