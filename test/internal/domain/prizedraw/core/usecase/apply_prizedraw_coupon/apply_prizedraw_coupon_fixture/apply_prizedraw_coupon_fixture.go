package apply_prizedraw_coupon_fixture

import (
	"getfund-api-v2/internal/domain/prizedraw/core/dto"
	"getfund-api-v2/internal/domain/prizedraw/core/usecase/apply_prizedraw_coupon"
	apply_prizedraw_coupon_application "getfund-api-v2/internal/domain/prizedraw/core/usecase/apply_prizedraw_coupon/application"
	event_prizedraw "getfund-api-v2/internal/domain/prizedraw/core/usecase/apply_prizedraw_coupon/event"
	"getfund-api-v2/test/helper/eventbus_spy"
	"getfund-api-v2/test/helper/repository_spy/prizedraw_repository_spy"
	"getfund-api-v2/test/helper/security_spy"
	"getfund-api-v2/test/helper/settings_spy"
	"time"
)

type ApplyPrizeDrawCouponFixture struct {
	RepoSpy     *prizedraw_repository_spy.CouponRepositorySpy
	BusSpy      *eventbus_spy.EventBusSpy
	SettingsSpy *settings_spy.ApplicationSettingsSpy
	HasherSpy   *security_spy.HasherSpy
}

func NewSut() (apply_prizedraw_coupon.UseCase, *ApplyPrizeDrawCouponFixture) {
	repoSpy := prizedraw_repository_spy.New()
	busSpy := eventbus_spy.New()
	hasherSpy := security_spy.New()

	return apply_prizedraw_coupon_application.New(repoSpy, busSpy, hasherSpy),
		&ApplyPrizeDrawCouponFixture{
			RepoSpy:   repoSpy,
			BusSpy:    busSpy,
			HasherSpy: hasherSpy,
		}
}

type Option func(*apply_prizedraw_coupon.Input)

func GetInput(options ...Option) *apply_prizedraw_coupon.Input {
	input := &apply_prizedraw_coupon.Input{
		CouponId:    1,
		PrizeDrawId: 1,
		ProductId:   1,
		UserId:      1,
	}

	for _, opt := range options {
		opt(input)
	}

	return input
}

func WithCouponId(couponId int) Option {
	return func(input *apply_prizedraw_coupon.Input) {
		input.CouponId = couponId
	}
}

func WithPrizeDrawId(prizeDrawId int) Option {
	return func(input *apply_prizedraw_coupon.Input) {
		input.PrizeDrawId = prizeDrawId
	}
}

func WithProductId(productId int) Option {
	return func(input *apply_prizedraw_coupon.Input) {
		input.ProductId = productId
	}
}

func WithUserId(userId int) Option {
	return func(input *apply_prizedraw_coupon.Input) {
		input.UserId = userId
	}
}

func (f *ApplyPrizeDrawCouponFixture) GetEntranceDto() *dto.EntranceDto {
	event := &event_prizedraw.ApplyPrizeDrawCouponStartedEvent{}
	return &dto.EntranceDto{
		LuckyCode:   f.HasherSpy.SuccessResult["GetRandomCode"].(string),
		UserId:      1,
		PrizeDrawId: 1,
		PurchaseId:  f.BusSpy.ReferenceResult["EmitAndWaitPromise"+event.GetName()].(int),
		IsDonation:  false,
	}
}

func GetValidCoupon() *dto.CouponDto {
	less72Hours := time.Now().Add(-24 * time.Hour).Unix()
	more24Hours := time.Now().Add(24 * time.Hour).Unix()
	email := "fake@mail.com"
	return &dto.CouponDto{
		ProductId:   10,
		PrizeDrawId: 5,
		Id:          1,
		CouponTypeApplicability: &dto.CouponTypeApplicabilityDto{
			StartAt:     less72Hours,
			EndAt:       &more24Hours,
			LinkedEmail: &email,
		},
	}
}

func ApplyCoupon(couponDto *dto.CouponDto, userId, couponId int) {
	couponDto.UserCouponApplies = make([]*dto.UserCouponApplyDto, 0)
	couponDto.UserCouponApplies = append(couponDto.UserCouponApplies, &dto.UserCouponApplyDto{
		CouponId: couponId,
		UserId:   userId,
	})
}
