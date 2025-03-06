package apply_prizedraw_coupon_application

import (
	"bytes"
	"errors"
	prizedraw_contract "getfund-api-v2/internal/domain/prizedraw/core/contract"
	"getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_dto"
	"getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_payload"
	"getfund-api-v2/internal/domain/prizedraw/core/entity"
	"getfund-api-v2/internal/domain/prizedraw/core/usecase/apply_prizedraw_coupon"
	"getfund-api-v2/internal/settings"
	"getfund-api-v2/internal/shared/app_constant"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/internal/shared/security"
	"getfund-api-v2/pkg/bus"
	"strconv"
	"time"
)

const (
	_TIME_OUT         = "timeout waiting for coupon apply"
	_EMPTY_RESPONSE   = "empty response from coupon apply"
	_INVALID_RESPONSE = "invalid response from coupon apply"
	_INVALID_PURCHASE = "invalid purchase id"
)

type applyPrizeDrawCouponApplication struct {
	repository prizedraw_contract.Repository
	bus        bus.EventBus
	settings   settings.ApplicationSettings
	hasher     security.Hasher
}

func New(repository prizedraw_contract.Repository, bus bus.EventBus, settings settings.ApplicationSettings, hasher security.Hasher) apply_prizedraw_coupon.UseCase {
	return &applyPrizeDrawCouponApplication{
		repository: repository,
		bus:        bus,
		settings:   settings,
		hasher:     hasher,
	}
}

func (a *applyPrizeDrawCouponApplication) Execute(input *apply_prizedraw_coupon.Input) (*apply_prizedraw_coupon.Output, *result_app.ApplicationError) {
	validatable := input.Validate()
	if validatable.IsInvalid() {
		return nil, result_app.New(result_app.UNPROCESSABLE_CONTENT_CODE, validatable.GetErrors())
	}
	channelResponse := make(chan []byte, 1)
	payload := &prizedraw_payload.ApplyPrizeDrawCouponStartedPayload{
		UserId:       input.UserId,
		ProductId:    input.ProductId,
		PrizeDrawId:  input.PrizeDrawId,
		CouponId:     input.CouponId,
		ItemQuantity: 1,
	}

	a.bus.EmitWithPayloadAndResponse(&apply_prizedraw_coupon.ApplyPrizeDrawCouponStartedEvent{}, payload, channelResponse)
	purchaseId, err := a.getPurchaseIdFromResponse(channelResponse)
	if err != nil {
		return nil, err
	}

	luckyCode, errRandom := a.hasher.GetRandomCode(8)
	if errRandom != nil || luckyCode == "" {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, errors.New("erro on build lucky number"))
	}

	entrance := entity.NewEntrance(luckyCode, input.UserId, input.PrizeDrawId, *purchaseId, false)
	entrance_dto := &prizedraw_dto.EntranceDto{
		LuckyCode:   entrance.GetLuckyCode(),
		UserId:      entrance.GetUserId(),
		PrizeDrawId: entrance.GetPrizeDrawId(),
		PurchaseId:  entrance.GetPurchaseId(),
		IsDonation:  entrance.GetIsDonation(),
		CreatedAt:   entrance.GetCreatedAt().Unix(),
		UpdatedAt:   entrance.GetUpdatedAt().Unix(),
	}

	_, errEntrance := a.repository.CreateEntrance(entrance_dto)
	if errEntrance != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, errors.New("erro on create entrance"))
	}

	//evento para enviar email
	return nil, nil
}

func (v *applyPrizeDrawCouponApplication) getPurchaseIdFromResponse(responseChannel chan []byte) (*int, *result_app.ApplicationError) {
	var purchaseId int
	var err error

	select {
	case response := <-responseChannel:
		if bytes.Equal(response, app_constant.EMPTYB) {
			return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_EMPTY_RESPONSE))
		}

		purchaseId, err = strconv.Atoi(string(response))
		if err != nil {
			return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_INVALID_RESPONSE))
		}

		if purchaseId == 0 {
			return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_INVALID_PURCHASE))
		}

	case <-time.After(time.Duration(v.settings.GetTimeoutResponseEvent()) * time.Second):
		return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_TIME_OUT))
	}

	return &purchaseId, nil
}
