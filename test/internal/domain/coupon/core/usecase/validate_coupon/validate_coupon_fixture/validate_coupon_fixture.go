package validate_coupon_fixture

import (
	coupon_dto "getfund-api-v2/internal/domain/coupon/core/dto/coupon_dto"
	"getfund-api-v2/internal/domain/coupon/core/usecase/validate_coupon"
	validate_coupon_application "getfund-api-v2/internal/domain/coupon/core/usecase/validate_coupon/application"
	"getfund-api-v2/test/helper/eventbus_spy"
	"getfund-api-v2/test/helper/repository_spy/coupon_repository_spy"
	"getfund-api-v2/test/helper/settings_spy"
	"time"
)

type ValidateCouponFixture struct {
	RepoSpy     *coupon_repository_spy.CouponRepositorySpy
	BusSpy      *eventbus_spy.EventBusSpy
	SettingsSpy *settings_spy.ApplicationSettingsSpy
}

func NewSut() (validate_coupon.UseCase, *ValidateCouponFixture) {
	repoSpy := coupon_repository_spy.New()
	busSpy := eventbus_spy.New()
	settingsSpy := settings_spy.New()

	return validate_coupon_application.New(repoSpy, busSpy, settingsSpy),
		&ValidateCouponFixture{
			RepoSpy:     repoSpy,
			BusSpy:      busSpy,
			SettingsSpy: settingsSpy,
		}

}

type Option func(*validate_coupon.Input)

func GetInput(options ...Option) *validate_coupon.Input {
	input := &validate_coupon.Input{
		CouponCode: "FAKE_CPN",
	}

	for _, opt := range options {
		opt(input)
	}

	return input
}

func WithEmptyCouponCode() Option {
	return func(params *validate_coupon.Input) {
		params.CouponCode = ""
	}
}

func WithInvalidCouponCode() Option {
	return func(params *validate_coupon.Input) {
		params.CouponCode = "fake" //less than 8 characters
	}
}

func GetValidCoupon() *coupon_dto.CouponDto {
	minus72Hours := time.Now().Add(-24 * time.Hour).Unix()
	minus24Hours := time.Now().Add(24 * time.Hour).Unix()
	return &coupon_dto.CouponDto{
		StartAt:     minus72Hours,
		EndAt:       &minus24Hours,
		ProductId:   10,
		PrizeDrawId: 5,
	}
}

func GetPrizeDrawResponse() []byte {
	return []byte(`{"prize_draw": {"winner_entrance_id": 1}}`)
}

func GetProductResponse() []byte {
	return []byte(`{"product": {"total_price": 100, "is_active": true}}`)
}
