package activate_user_application

import (
	"encoding/json"
	"errors"
	"getfund-api-v2/internal/config/env"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/internal/domain/auth/core/domain_service/activate_user_mapper"
	"getfund-api-v2/internal/domain/auth/core/dto"
	"getfund-api-v2/internal/domain/auth/core/entity"
	"getfund-api-v2/internal/domain/auth/core/usecase/activate_user"
	"getfund-api-v2/internal/domain/auth/core/usecase/activate_user/event"
	shared_bus "getfund-api-v2/internal/shared/bus"
	"getfund-api-v2/internal/shared/cache"
	shared_error "getfund-api-v2/internal/shared/error"
)

type activateUserApplication struct {
	cache      cache.Service
	repository auth_contract.Repository
	mapper     activate_user_mapper.Mapper
	bus        shared_bus.EventBus
	env        env.Variable
}

func New(cache cache.Service, repository auth_contract.Repository, mapper activate_user_mapper.Mapper, bus shared_bus.EventBus, env env.Variable) activate_user.UseCase {
	return &activateUserApplication{
		cache:      cache,
		repository: repository,
		mapper:     mapper,
		bus:        bus,
		env:        env,
	}
}

func (a *activateUserApplication) Execute(input *activate_user.Input) (*activate_user.Output, *shared_error.Error) {
	validatable := input.Validate()
	if validatable.IsInvalid() {
		return nil, shared_error.New(shared_error.UNAUTHORIZED_CODE, validatable.GetErrors())
	}

	userData, errData := getUserData(input, a.cache)
	if errData != nil {
		return nil, errData
	}

	if err := checkForDuplicateUser(input, userData, a.cache, a.repository); err != nil {
		return nil, err
	}

	user := entity.NewUser(
		userData.FirstName,
		userData.LastName,
		userData.Username,
		userData.Password,
	)

	userDto := a.mapper.ToDto(user)

	userSaved, err := signup(userDto, a.repository)
	if err != nil {
		return nil, err
	}

	defer a.cache.Delete(input.ActivationDataKey)

	if userData.CouponCode != "" {
		payloadCoupon := &event.ActivateUserWithCouponConfirmedPayload{
			UserId:     userSaved.Id,
			CouponCode: userData.CouponCode,
			Email:      userData.Username,
		}

		a.bus.EmitWithPayload(&event.ActivateUserWithCouponConfirmedEvent{}, payloadCoupon)
	}

	return &activate_user.Output{
		Username: user.GetUsername(),
		Password: user.GetPassword(),
	}, nil
}

func getUserData(input *activate_user.Input, cache cache.Service) (*dto.ActivationUserData, *shared_error.Error) {
	userSerialized, errCache := cache.Get(input.ActivationDataKey)
	if errCache != nil {
		return nil, shared_error.New(shared_error.NOT_FOUND_CODE, errors.New("activation code not found"))
	}

	var userData = dto.ActivationUserData{}
	if err := json.Unmarshal([]byte(userSerialized), &userData); err != nil {
		return nil, shared_error.New(shared_error.SERVER_ERROR_CODE, errors.New("error on get user data"))
	}

	return &userData, nil
}

func checkForDuplicateUser(input *activate_user.Input, userData *dto.ActivationUserData, cache cache.Service, repository auth_contract.Repository) *shared_error.Error {
	userDuplicated, errRepo := repository.UserExists(userData.Username)
	if errRepo != nil {
		return shared_error.New(shared_error.SERVER_ERROR_CODE, errRepo)
	}

	if userDuplicated != nil {
		defer cache.Delete(input.ActivationDataKey)
		return shared_error.New(shared_error.DUPLICATED_ENTRY_CODE, errors.New("user already exists"))
	}

	return nil
}

func signup(userDto *dto.ActivationUserDto, repository auth_contract.Repository) (*dto.UserDto, *shared_error.Error) {
	userSaved, err := repository.Signup(userDto)
	if err != nil {
		return nil, shared_error.New(shared_error.SERVER_ERROR_CODE, err)
	}

	return userSaved, nil
}
