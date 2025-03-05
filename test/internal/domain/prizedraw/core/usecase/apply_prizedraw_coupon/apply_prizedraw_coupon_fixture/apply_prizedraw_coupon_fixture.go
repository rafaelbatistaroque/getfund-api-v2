package apply_prizedraw_coupon_fixture

import (
	"getfund-api-v2/internal/domain/prizedraw/core/usecase/apply_prizedraw_coupon"
	apply_prizedraw_coupon_application "getfund-api-v2/internal/domain/prizedraw/core/usecase/apply_prizedraw_coupon/application"
	"getfund-api-v2/test/helper/eventbus_spy"
	"getfund-api-v2/test/helper/repository_spy/prizedraw_repository_spy"
	"getfund-api-v2/test/helper/settings_spy"
)

type ApplyPrizeDrawCouponFixture struct {
	RepoSpy     *prizedraw_repository_spy.CouponRepositorySpy
	BusSpy      *eventbus_spy.EventBusSpy
	SettingsSpy *settings_spy.ApplicationSettingsSpy
}

func NewSut() (apply_prizedraw_coupon.UseCase, *ApplyPrizeDrawCouponFixture) {
	repoSpy := prizedraw_repository_spy.New()
	busSpy := eventbus_spy.New()
	settingsSpy := settings_spy.New()

	return apply_prizedraw_coupon_application.New(repoSpy, busSpy, settingsSpy),
		&ApplyPrizeDrawCouponFixture{
			RepoSpy:     repoSpy,
			BusSpy:      busSpy,
			SettingsSpy: settingsSpy,
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
