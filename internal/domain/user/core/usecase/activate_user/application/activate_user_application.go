package activate_user_application

import (
	"encoding/json"
	"errors"
	user_contract "getfund-api-v2/internal/domain/user/core/contract"
	"getfund-api-v2/internal/domain/user/core/domain_service/activate_user_mapper"
	"getfund-api-v2/internal/domain/user/core/entity/activate_user_entity"
	"getfund-api-v2/internal/domain/user/core/usecase/activate_user"
	"getfund-api-v2/internal/domain/user/core/user_dto"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/internal/shared/service/cache_service"
)

const (
	_KEY_USER_ACTIVATION_PREFIX = "user_activation_"
)

type activateUserApplication struct {
	cache      cache_service.Cache
	repository user_contract.Repository
	mapper     activate_user_mapper.Mapper
}

func New(cache cache_service.Cache, repository user_contract.Repository, mapper activate_user_mapper.Mapper) activate_user.UseCase {
	return &activateUserApplication{
		cache:      cache,
		repository: repository,
		mapper:     mapper,
	}
}

func (a *activateUserApplication) Execute(input *activate_user.Input) (*activate_user.Output, *result_app.ApplicationError) {
	validatable := input.Validate()
	if validatable.IsInvalid() {
		return nil, result_app.New(result_app.UNAUTHORIZED_CODE, validatable.GetErrors())
	}

	keyCache := _KEY_USER_ACTIVATION_PREFIX + input.ActivationCode
	userSerialized, errCache := a.cache.Get(keyCache)
	if errCache != nil {
		return nil, result_app.New(result_app.NOT_FOUND_CODE, errors.New("activation code not found"))
	}

	var userData = user_dto.ActivationUserData{}
	if err := json.Unmarshal([]byte(userSerialized), &userData); err != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, errors.New("error on get user data"))
	}

	userDuplicated, errRepo := a.repository.GetUserByUsername(userData.Email)
	if errRepo != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, errRepo)
	}

	if userDuplicated != nil {
		defer a.cache.Delete(keyCache)
		return nil, result_app.New(result_app.DUPLICATED_ENTRY_CODE, errors.New("user already exists"))
	}

	user := activate_user_entity.New(
		userData.FirstName,
		userData.LastName,
		userData.Email,
		userData.Gender,
		userData.Password,
		userData.MainSocialNetwork,
		userData.RegisteredUrl,
		userData.CountryId,
		userData.UserCategoryId)

	a.mapper.ToDto(user)

	return nil, nil
}
