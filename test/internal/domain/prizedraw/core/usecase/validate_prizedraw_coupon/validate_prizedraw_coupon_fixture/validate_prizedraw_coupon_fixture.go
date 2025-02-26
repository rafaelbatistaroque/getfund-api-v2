package validate_prizedraw_coupon_fixture

import (
	"getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_dto"
	"getfund-api-v2/internal/domain/prizedraw/core/entity"
	"getfund-api-v2/internal/domain/prizedraw/core/usecase/validate_prizedraw_coupon"
	validate_coupon_application "getfund-api-v2/internal/domain/prizedraw/core/usecase/validate_prizedraw_coupon/application"
	"getfund-api-v2/test/helper/eventbus_spy"
	"getfund-api-v2/test/helper/repository_spy/prizedraw_repository_spy"
	"getfund-api-v2/test/helper/settings_spy"
	"time"
)

type ValidateCouponFixture struct {
	RepoSpy     *prizedraw_repository_spy.CouponRepositorySpy
	BusSpy      *eventbus_spy.EventBusSpy
	SettingsSpy *settings_spy.ApplicationSettingsSpy
}

func NewSut() (validate_prizedraw_coupon.UseCase, *ValidateCouponFixture) {
	repoSpy := prizedraw_repository_spy.New()
	busSpy := eventbus_spy.New()
	settingsSpy := settings_spy.New()

	return validate_coupon_application.New(repoSpy, busSpy, settingsSpy),
		&ValidateCouponFixture{
			RepoSpy:     repoSpy,
			BusSpy:      busSpy,
			SettingsSpy: settingsSpy,
		}

}

type Option func(*validate_prizedraw_coupon.Input)

func GetInput(options ...Option) *validate_prizedraw_coupon.Input {
	input := &validate_prizedraw_coupon.Input{
		CouponCode:          "FAKE_CPN",
		SelectedProductId:   10,
		SelectedPrizeDrawId: 5,
		UserId:              1,
		Email:               "fake@email.com",
	}

	for _, opt := range options {
		opt(input)
	}

	return input
}

func WithEmptyCouponCode() Option {
	return func(params *validate_prizedraw_coupon.Input) {
		params.CouponCode = ""
	}
}

func WithInvalidCouponCode() Option {
	return func(params *validate_prizedraw_coupon.Input) {
		params.CouponCode = "fake" //less than 8 characters
	}
}

func WithEmptyEmail() Option {
	return func(params *validate_prizedraw_coupon.Input) {
		params.Email = ""
	}
}

func WithInvalidEmail() Option {
	return func(params *validate_prizedraw_coupon.Input) {
		params.Email = "fake" //not mail
	}
}

func WithSelectedProductId(id int) Option {
	return func(params *validate_prizedraw_coupon.Input) {
		params.SelectedProductId = id
	}
}

func WithSelectedPrizeDrawId(id int) Option {
	return func(params *validate_prizedraw_coupon.Input) {
		params.SelectedPrizeDrawId = id
	}
}

func GetValidCoupon() *prizedraw_dto.CouponDto {
	less72Hours := time.Now().Add(-24 * time.Hour).Unix()
	more24Hours := time.Now().Add(24 * time.Hour).Unix()
	email := "fake@mail.com"
	return &prizedraw_dto.CouponDto{
		ProductId:   10,
		PrizeDrawId: 5,
		Id:          1,
		CouponTypeApplicability: &prizedraw_dto.CouponTypeApplicabilityDto{
			StartAt:     less72Hours,
			EndAt:       &more24Hours,
			LinkedEmail: &email,
		},
	}
}

func GetCouponNotStartYet(startAt time.Duration) *prizedraw_dto.CouponDto {
	validCoupon := GetValidCoupon()
	validCoupon.CouponTypeApplicability.StartAt = int64(startAt)

	return validCoupon
}

func GetValidCouponWithApplication(userId int, couponTypeCode string) *prizedraw_dto.CouponDto {
	validCoupon := GetValidCoupon()
	validCoupon.CouponTypeApplicability.CouponTypeCode = couponTypeCode
	validCoupon.UserCouponApplies = addUserApplies(2, userId)

	return validCoupon
}

func GetValidCouponWithEmailLinked() *prizedraw_dto.CouponDto {
	validCoupon := GetValidCoupon()
	validCoupon.CouponTypeApplicability.CouponTypeCode = entity.UNIQUE_APPLICATION_BY_EMAIL_TYPE
	validCoupon.UserCouponApplies = addUserApplies(1, 1)

	return validCoupon
}

func GetValidCouponWithApplicationReached(limit int, couponTypeCode string) *prizedraw_dto.CouponDto {
	validCoupon := GetValidCoupon()
	validCoupon.CouponTypeApplicability.LimitApplication = &limit
	validCoupon.CouponTypeApplicability.CouponTypeCode = couponTypeCode
	validCoupon.UserCouponApplies = addUserApplies(limit, 0)

	return validCoupon
}

func GetExpiredCoupon() *prizedraw_dto.CouponDto {
	minus72Hours := time.Now().Add(-72 * time.Hour).Unix()
	minus24Hours := time.Now().Add(-24 * time.Hour).Unix()
	validCoupon := GetValidCoupon()
	validCoupon.CouponTypeApplicability.StartAt = minus72Hours
	validCoupon.CouponTypeApplicability.EndAt = &minus24Hours
	validCoupon.CouponTypeApplicability.CouponTypeCode = entity.EXPIRATION_TYPE

	return validCoupon
}

func addUserApplies(limit, userId int) []prizedraw_dto.UserCouponApplyDto {
	userApplies := make([]prizedraw_dto.UserCouponApplyDto, limit)
	if userId != 0 {
		userApplies[0] = prizedraw_dto.UserCouponApplyDto{
			UserId: userId,
		}
	}

	for id := range limit {
		userApplies[id] = prizedraw_dto.UserCouponApplyDto{
			UserId: id,
		}

	}
	return userApplies
}

func GetValidPrizeDraw() *prizedraw_dto.PrizeDrawDto {
	return &prizedraw_dto.PrizeDrawDto{Id: 5}
}

func GetProductResponse() []byte {
	return []byte(`{"product": {"id":10, "total_price": 100, "is_active": true, "entrance_qty": 1}}`)
}

func GetInactiveProductResponse() []byte {
	return []byte(`{"product": {"id":10, "total_price": 100, "is_active": false, "entrance_qty": 1}}`)
}

func GetNullResponse() []byte {
	return []byte(`{}`)
}
