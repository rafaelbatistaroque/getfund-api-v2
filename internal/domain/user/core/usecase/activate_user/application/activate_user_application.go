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
	"getfund-api-v2/pkg/bus"
	"getfund-api-v2/pkg/bus/event"
)

type activateUserApplication struct {
	cache      cache_service.Cache
	repository user_contract.Repository
	mapper     activate_user_mapper.Mapper
	bus        bus.EventBus
}

func New(cache cache_service.Cache, repository user_contract.Repository, mapper activate_user_mapper.Mapper, bus bus.EventBus) activate_user.UseCase {
	return &activateUserApplication{
		cache:      cache,
		repository: repository,
		mapper:     mapper,
		bus:        bus,
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

	userSaved, err := saveUser(userData, a.mapper, a.repository)
	if err != nil {
		return nil, err
	}

	defer a.cache.Delete(input.ActivationKey)

	if userData.CouponCode != "" {
		payloadCoupon := &user_dto.UserActivationWithCouponPayloadDto{
			ActivationCode: input.ActivationCode,
			UserId:         userSaved.Id,
		}
		a.bus.EmitWithPayload(&event.UserActivationWithCouponConfirmed{}, payloadCoupon)
	}

	return nil, nil
}

func getUserData(input *activate_user.Input, cache cache_service.Cache) (*user_dto.ActivationUserData, *result_app.ApplicationError) {
	userSerialized, errCache := cache.Get(input.ActivationKey)
	if errCache != nil {
		return nil, result_app.New(result_app.NOT_FOUND_CODE, errors.New("activation code not found"))
	}

	var userData = user_dto.ActivationUserData{}
	if err := json.Unmarshal([]byte(userSerialized), &userData); err != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, errors.New("error on get user data"))
	}

	return &userData, nil
}

func checkForDuplicateUser(input *activate_user.Input, userData *user_dto.ActivationUserData, cache cache_service.Cache, repository user_contract.Repository) *result_app.ApplicationError {
	userDuplicated, errRepo := repository.GetUserByUsername(userData.Email)
	if errRepo != nil {
		return result_app.New(result_app.SERVER_ERROR_CODE, errRepo)
	}

	if userDuplicated != nil {
		defer cache.Delete(input.ActivationKey)
		return result_app.New(result_app.DUPLICATED_ENTRY_CODE, errors.New("user already exists"))
	}

	return nil
}

func saveUser(userData *user_dto.ActivationUserData, mapper activate_user_mapper.Mapper, repository user_contract.Repository) (*user_dto.UserDto, *result_app.ApplicationError) {
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

	userDto := mapper.ToDto(user)

	userSaved, err := repository.SaveUser(userDto)
	if err != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, err)
	}

	return userSaved, nil
}
