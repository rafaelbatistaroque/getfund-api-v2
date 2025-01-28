package create_user_application

import (
	"errors"
	user_contract "getfund-api-v2/internal/domain/user/core/contract"
	"getfund-api-v2/internal/domain/user/core/usecase/create_user"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/internal/shared/security"
	"getfund-api-v2/internal/shared/service/cache_service"
	"time"
)

const (
	key_cache_prefix = "user_activation_"
)

type createUserApplication struct {
	repository user_contract.Repository
	hasher     security.Hasher
	cache      cache_service.Cache
}

func New(repository user_contract.Repository, hasher security.Hasher, cache cache_service.Cache) create_user.UseCase {
	return &createUserApplication{
		repository: repository,
		hasher:     hasher,
		cache:      cache,
	}
}

func (c *createUserApplication) Execute(input *create_user.Input) (*create_user.Output, *result_app.ApplicationError) {
	validated := input.Validate()
	if validated.IsInvalid() {
		return nil, result_app.New(result_app.BAD_REQUEST_CODE, validated.GetErrors())
	}

	userDuplicated, err := c.repository.GetUserByUsername(input.Email)
	if err != nil {
		return nil, result_app.New(result_app.NOT_FOUND_CODE, err)
	}

	if userDuplicated != nil {
		return nil, result_app.New(result_app.DUPLICATED_ENTRY_CODE, errors.New("user already exists"))
	}

	keyCache, errCode := buildActivationCode(c.hasher)
	if errCode != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, errCode)
	}

	c.cache.Set(keyCache, input, 24*time.Hour)

	return nil, nil
}

func buildActivationCode(hasher security.Hasher) (string, error) {
	activationCode, errCode := hasher.GetRandomCode(20)
	if errCode != nil {
		return "", errors.New("error to save user")
	}

	return key_cache_prefix + activationCode, nil
}
