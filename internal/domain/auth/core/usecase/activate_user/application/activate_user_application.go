package activate_user_application

import (
	"encoding/json"
	"errors"
	"getfund-api-v2/internal/domain/auth/core/auth_dto"
	payload "getfund-api-v2/internal/domain/auth/core/auth_dto/auth_payload"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/internal/domain/auth/core/domain_service/activate_user_mapper"
	"getfund-api-v2/internal/domain/auth/core/entity/user_entity"
	"getfund-api-v2/internal/domain/auth/core/usecase/activate_user"
	"getfund-api-v2/internal/settings"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/internal/shared/service/cache_service"
	"getfund-api-v2/pkg/bus"
)

type activateUserApplication struct {
	cache      cache_service.Cache
	repository auth_contract.Repository
	mapper     activate_user_mapper.Mapper
	bus        bus.EventBus
	settings   settings.ApplicationSettings
}

func New(cache cache_service.Cache, repository auth_contract.Repository, mapper activate_user_mapper.Mapper, bus bus.EventBus, settings settings.ApplicationSettings) activate_user.UseCase {
	return &activateUserApplication{
		cache:      cache,
		repository: repository,
		mapper:     mapper,
		bus:        bus,
		settings:   settings,
	}
}

func (a *activateUserApplication) Execute(input *activate_user.Input) (*activate_user.Output, *result_app.ApplicationError) {
	validatable := input.Validate()
	if validatable.IsInvalid() {
		return nil, result_app.New(result_app.UNAUTHORIZED_CODE, validatable.GetErrors())
	}

	userData, errData := getUserData(input, a.cache)
	if errData != nil {
		return nil, errData
	}

	if err := checkForDuplicateUser(input, userData, a.cache, a.repository); err != nil {
		return nil, err
	}

	user := user_entity.New(
		userData.FirstName,
		userData.LastName,
		userData.Username,
		userData.Password,
	)

	userDto := a.mapper.ToDto(user)

	userSaved, err := createUser(userDto, a.repository)
	if err != nil {
		return nil, err
	}

	defer a.cache.Delete(input.ActivationDataKey)

	if userData.CouponCode != "" {
		payloadCoupon := &payload.ActivateUserWithCouponConfirmedPayload{
			UserId:     userSaved.Id,
			CouponCode: userData.CouponCode,
			Email:      userData.Username,
		}

		a.bus.EmitWithPayload(&activate_user.ActivateUserWithCouponConfirmedEvent{}, payloadCoupon)
	}

	return &activate_user.Output{
		Username: user.GetUsername(),
		Password: user.GetPassword(),
	}, nil
}

func getUserData(input *activate_user.Input, cache cache_service.Cache) (*auth_dto.ActivationUserData, *result_app.ApplicationError) {
	userSerialized, errCache := cache.Get(input.ActivationDataKey)
	if errCache != nil {
		return nil, result_app.New(result_app.NOT_FOUND_CODE, errors.New("activation code not found"))
	}

	var userData = auth_dto.ActivationUserData{}
	if err := json.Unmarshal([]byte(userSerialized), &userData); err != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, errors.New("error on get user data"))
	}

	return &userData, nil
}

func checkForDuplicateUser(input *activate_user.Input, userData *auth_dto.ActivationUserData, cache cache_service.Cache, repository auth_contract.Repository) *result_app.ApplicationError {
	userDuplicated, errRepo := repository.UserExists(userData.Username)
	if errRepo != nil {
		return result_app.New(result_app.SERVER_ERROR_CODE, errRepo)
	}

	if userDuplicated != nil {
		defer cache.Delete(input.ActivationDataKey)
		return result_app.New(result_app.DUPLICATED_ENTRY_CODE, errors.New("user already exists"))
	}

	return nil
}

func createUser(userDto *auth_dto.ActivationUserDto, repository auth_contract.Repository) (*auth_dto.UserDto, *result_app.ApplicationError) {
	userSaved, err := repository.CreateUser(userDto)
	if err != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, err)
	}

	return userSaved, nil
}
