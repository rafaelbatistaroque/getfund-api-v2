package validate_prizedraw_coupon_fixture

import (
	"getfund-api-v2/internal/domain/prizedraw/core/dto"
	"getfund-api-v2/internal/domain/prizedraw/core/entity"
	"getfund-api-v2/internal/domain/prizedraw/core/usecase/validate_prizedraw_coupon"
	validate_coupon_application "getfund-api-v2/internal/domain/prizedraw/core/usecase/validate_prizedraw_coupon/application"
	"getfund-api-v2/test/helper/eventbus_spy"
	"getfund-api-v2/test/helper/repository_spy/prizedraw_repository_spy"
	"getfund-api-v2/test/helper/settings_spy"
	"time"
)

type ValidatePrizeDrawCouponFixture struct {
	RepoSpy     *prizedraw_repository_spy.CouponRepositorySpy
	BusSpy      *eventbus_spy.EventBusSpy
	SettingsSpy *settings_spy.ApplicationSettingsSpy
}

func NewSut() (validate_prizedraw_coupon.UseCase, *ValidatePrizeDrawCouponFixture) {
	repoSpy := prizedraw_repository_spy.New()
	busSpy := eventbus_spy.New()
	settingsSpy := settings_spy.New()

	return validate_coupon_application.New(repoSpy, busSpy, settingsSpy),
		&ValidatePrizeDrawCouponFixture{
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

func WithUserId(id int) Option {
	return func(params *validate_prizedraw_coupon.Input) {
		params.UserId = id
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

func GetCouponWithoutPrizeDrawLinked(id int) *dto.CouponDto {
	validCoupon := GetValidCoupon()
	validCoupon.PrizeDrawId = id

	return validCoupon
}

func GetCouponNotStartYet(startAt time.Duration) *dto.CouponDto {
	validCoupon := GetValidCoupon()
	validCoupon.CouponTypeApplicability.StartAt = int64(startAt)

	return validCoupon
}

func GetValidCouponWithApplication(userId int, couponTypeCode string) *dto.CouponDto {
	validCoupon := GetValidCoupon()
	validCoupon.CouponTypeApplicability.CouponTypeCode = couponTypeCode
	validCoupon.UserCouponApplies = addUserApplies(2, userId)

	return validCoupon
}

func GetValidCouponWithEmailLinked() *dto.CouponDto {
	validCoupon := GetValidCoupon()
	validCoupon.CouponTypeApplicability.CouponTypeCode = entity.UNIQUE_APPLICATION_BY_EMAIL_TYPE
	validCoupon.UserCouponApplies = addUserApplies(1, 1)

	return validCoupon
}

func GetValidCouponWithApplicationReached(limit int, couponTypeCode string) *dto.CouponDto {
	validCoupon := GetValidCoupon()
	validCoupon.CouponTypeApplicability.LimitApplication = &limit
	validCoupon.CouponTypeApplicability.CouponTypeCode = couponTypeCode
	validCoupon.UserCouponApplies = addUserApplies(limit, 0)

	return validCoupon
}

func GetExpiredCoupon() *dto.CouponDto {
	minus72Hours := time.Now().Add(-72 * time.Hour).Unix()
	minus24Hours := time.Now().Add(-24 * time.Hour).Unix()
	validCoupon := GetValidCoupon()
	validCoupon.CouponTypeApplicability.StartAt = minus72Hours
	validCoupon.CouponTypeApplicability.EndAt = &minus24Hours
	validCoupon.CouponTypeApplicability.CouponTypeCode = entity.EXPIRATION_TYPE

	return validCoupon
}

func addUserApplies(limit, userId int) []*dto.UserCouponApplyDto {
	userApplies := make([]*dto.UserCouponApplyDto, limit)
	if userId != 0 {
		userApplies[0] = &dto.UserCouponApplyDto{
			UserId: userId,
		}
	}

	for id := range limit {
		userApplies[id] = &dto.UserCouponApplyDto{
			UserId: id,
		}

	}
	return userApplies
}

func GetValidPrizeDraw() *dto.PrizeDrawDto {
	return &dto.PrizeDrawDto{Id: 5}
}

func GetProductWithDiferentIdResponse() *dto.ProductDto {
	return &dto.ProductDto{IsActive: true, Id: 3}
}

func GetProductResponse() *dto.ProductDto {
	return &dto.ProductDto{IsActive: true, Id: 10}
}

func GetInactiveProductResponse() *dto.ProductDto {
	return &dto.ProductDto{IsActive: false}
}
