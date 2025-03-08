package apply_prizedraw_coupon_application

import (
	"errors"
	prizedraw_contract "getfund-api-v2/internal/domain/prizedraw/core/contract"
	"getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_dto"
	"getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_payload"
	"getfund-api-v2/internal/domain/prizedraw/core/entity"
	"getfund-api-v2/internal/domain/prizedraw/core/usecase/apply_prizedraw_coupon"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/internal/shared/security"
	"getfund-api-v2/pkg/bus"
)

const (
	_INVALID_PURCHASE = "invalid purchase id"
	_COUPON_APPLY     = "[coupon apply] "
)

type applyPrizeDrawCouponApplication struct {
	repository prizedraw_contract.Repository
	bus        bus.EventBus
	hasher     security.Hasher
}

func New(repository prizedraw_contract.Repository, bus bus.EventBus, hasher security.Hasher) apply_prizedraw_coupon.UseCase {
	return &applyPrizeDrawCouponApplication{
		repository: repository,
		bus:        bus,
		hasher:     hasher,
	}
}

func (a *applyPrizeDrawCouponApplication) Execute(input *apply_prizedraw_coupon.Input) (*apply_prizedraw_coupon.Output, *result_app.ApplicationError) {
	validatable := input.Validate()
	if validatable.IsInvalid() {
		return nil, result_app.New(result_app.UNPROCESSABLE_CONTENT_CODE, validatable.GetErrors())
	}

	payload := &prizedraw_payload.ApplyPrizeDrawCouponStartedPayload{
		UserId:       input.UserId,
		ProductId:    input.ProductId,
		PrizeDrawId:  input.PrizeDrawId,
		CouponId:     input.CouponId,
		ItemQuantity: 1,
	}

	var purchaseId int
	promise := a.bus.EmitAndWaitPromise(&apply_prizedraw_coupon.ApplyPrizeDrawCouponStartedEvent{}, payload, &purchaseId)
	if !promise.IsErrorNil() {
		return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_COUPON_APPLY+promise.ErrorToString()))
	}

	if purchaseId == 0 {
		return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_COUPON_APPLY+_INVALID_PURCHASE))
	}

	luckyCode, errRandom := a.hasher.GetRandomCode(8)
	if errRandom != nil || luckyCode == "" {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, errors.New("erro on build lucky number"))
	}

	entrance := entity.NewEntrance(luckyCode, input.UserId, input.PrizeDrawId, 1, false)
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

	//salvar entrance em cache com chave prize_draw_coupon_applied_key_entrance
	//evento para enviar email com entranceId
	return nil, nil
}
