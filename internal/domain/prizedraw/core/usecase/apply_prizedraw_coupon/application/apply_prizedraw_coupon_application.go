package apply_prizedraw_coupon_application

import (
	"errors"
	prizedraw_contract "getfund-api-v2/internal/domain/prizedraw/core/contract"
	"getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_dto"
	"getfund-api-v2/internal/domain/prizedraw/core/entity"
	"getfund-api-v2/internal/domain/prizedraw/core/usecase/apply_prizedraw_coupon"
	"getfund-api-v2/internal/domain/prizedraw/core/usecase/apply_prizedraw_coupon/event"
	shared_bus "getfund-api-v2/internal/shared/bus"
	shared_error "getfund-api-v2/internal/shared/error"
	"getfund-api-v2/internal/shared/security"
)

const (
	_INVALID_PURCHASE = "invalid purchase id"
	_COUPON_APPLY     = "[coupon apply] "
)

type applyPrizeDrawCouponApplication struct {
	repository prizedraw_contract.Repository
	bus        shared_bus.EventBus
	hasher     security.Hasher
}

func New(repository prizedraw_contract.Repository, bus shared_bus.EventBus, hasher security.Hasher) apply_prizedraw_coupon.UseCase {
	return &applyPrizeDrawCouponApplication{
		repository: repository,
		bus:        bus,
		hasher:     hasher,
	}
}

func (a *applyPrizeDrawCouponApplication) Execute(input *apply_prizedraw_coupon.Input) (*apply_prizedraw_coupon.Output, *shared_error.Error) {
	if validatable := input.Validate(); validatable.IsInvalid() {
		return nil, shared_error.New(shared_error.UNPROCESSABLE_CONTENT_CODE, validatable.GetErrors())
	}

	purchaseIdReceived, err := a.emitAndAwaitPurchaseId(input)
	if err != nil {
		return nil, err
	}

	luckyCode, err := a.generateLuckyCode()
	if err != nil {
		return nil, err
	}

	coupon, err := a.getCoupon(input.CouponId)
	if err != nil {
		return nil, err
	}

	entrance := entity.NewEntrance(luckyCode, input.UserId, input.PrizeDrawId, purchaseIdReceived, false)
	coupon.ApplyCoupon(input.UserId)
	if err := a.saveEntranceAndCoupon(entrance, coupon); err != nil {
		a.emitFailureEvent(purchaseIdReceived)
		return nil, err
	}

	return &apply_prizedraw_coupon.Output{
		Message: "coupon successfully applied",
	}, nil
}

func (a *applyPrizeDrawCouponApplication) emitAndAwaitPurchaseId(input *apply_prizedraw_coupon.Input) (int, *shared_error.Error) {
	payload := &event.ApplyPrizeDrawCouponStartedPayload{
		UserId:       input.UserId,
		ProductId:    input.ProductId,
		PrizeDrawId:  input.PrizeDrawId,
		CouponId:     input.CouponId,
		ItemQuantity: 1,
	}

	var purchaseIdReceived int
	promise := a.bus.EmitAndWaitPromise(&event.ApplyPrizeDrawCouponStartedEvent{}, payload, &purchaseIdReceived)

	if promise.HasError() {
		return 0, shared_error.New(shared_error.UNAVAILABLE_CODE, errors.New(_COUPON_APPLY+promise.ErrorToString()))
	}

	if purchaseIdReceived == 0 {
		return 0, shared_error.New(shared_error.UNAVAILABLE_CODE, errors.New(_COUPON_APPLY+_INVALID_PURCHASE))
	}

	return purchaseIdReceived, nil
}

func (a *applyPrizeDrawCouponApplication) generateLuckyCode() (string, *shared_error.Error) {
	luckyCode, err := a.hasher.GetRandomCode(8)
	if err != nil || luckyCode == "" {
		return "", shared_error.New(shared_error.SERVER_ERROR_CODE, errors.New("erro on build lucky number"))
	}
	return luckyCode, nil
}

func (a *applyPrizeDrawCouponApplication) getCoupon(couponId int) (*entity.Coupon, *shared_error.Error) {
	couponDto, err := a.repository.GetCouponById(couponId)
	if err != nil || couponDto == nil {
		return nil, shared_error.New(shared_error.UNAVAILABLE_CODE, err)
	}

	return couponDto.ToEntity(), nil
}

func (a *applyPrizeDrawCouponApplication) saveEntranceAndCoupon(entrance *entity.Entrance, coupon *entity.Coupon) *shared_error.Error {
	if err := a.repository.SaveEntranceWithCouponApplied(prizedraw_dto.ToEntranceDto(entrance), prizedraw_dto.ToCouponDto(coupon)); err != nil {
		return shared_error.New(shared_error.UNAVAILABLE_CODE, errors.New("erro on apply coupon"))
	}
	return nil
}

func (a *applyPrizeDrawCouponApplication) emitFailureEvent(purchaseIdReceived int) {
	var promiseResolvedWithSuccess bool
	a.bus.EmitAndWaitPromise(&event.ApplyPrizeDrawCouponFailedEvent{}, &event.ApplyPrizeDrawCouponFailedPayload{
		PurchaseId: purchaseIdReceived,
	}, &promiseResolvedWithSuccess)
}
