package create_user_process_with_coupon_started_event_handler_test

import (
	coupon_common "getfund-api-v2/internal/domain/coupon/adapter/common"
	coupon_dto "getfund-api-v2/internal/domain/coupon/core/dto/coupon_dto"
	fixture "getfund-api-v2/test/internal/domain/coupon/adapter/event_handler/create_user_process_with_coupon_started_event_handler/create_user_process_with_coupon_started_event_handler_fixture"
	"testing"
	"time"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenHandler_WhenPayloadParseError_ThenEnsureNeverCallValidadeCoupon(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Handle(fixture.GetInvalidCreateUserProcessWithCouponStartedEvent())

	//Assert
	verify.Should(t, spies.RepoSpy.CallsCount["Execute"]).Be(0)
}

func Test_GivenHandler_WhenPayloadParseSuccess_ThenEnsureCallGetCouponByCodeWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Handle(fixture.GetValidCreateUserProcessWithCouponStartedEvent())

	//Assert
	verify.Should(t, spies.RepoSpy.Params["GetCouponByCode:couponCode"]).Be("fake-coupon-code")
}

func Test_GivenHandler_WhenGetCouponByCodeInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Handle(fixture.GetValidCreateUserProcessWithCouponStartedEvent())

	//Assert
	verify.Should(t, spies.RepoSpy.CallsCount["GetCouponByCode"]).Be(1)
}

func Test_GivenHandler_WhenGetCouponByCodeError_ThenEnsureCallCacheSetWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeError()
	expectedCacheKey := "fake-activation-data-key_coupon"
	expectedMessageError := &coupon_common.ErrorData{Code: "COUPON_REPOSITORY", Message: spies.RepoSpy.ErrorResult["GetCouponByCode"].Error()}
	expectedCacheValue := fixture.GetCachedCouponData(expectedMessageError, nil)

	// Act
	sut.Handle(fixture.GetValidCreateUserProcessWithCouponStartedEvent())

	//Assert
	verify.Should(t, spies.CacheSpy.Params["Set:key"]).Be(expectedCacheKey)
	verify.Should(t, spies.CacheSpy.Params["Set:value"]).Be(expectedCacheValue)
	verify.Should(t, spies.CacheSpy.Params["Set:time"]).Be(24 * time.Hour)
}

func Test_GivenHandler_WhenGetCouponByCodeSuccess_ThenEnsureCallCacheSetWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(nil)
	expectedCacheKey := "fake-activation-data-key_coupon"
	expectedSuccess := fixture.GetCouponDataFromSuccessDB(spies.RepoSpy.SuccessResult["GetCouponByCode"].(*coupon_dto.CouponDto))
	expectedCacheValue := fixture.GetCachedCouponData(nil, expectedSuccess)

	// Act
	sut.Handle(fixture.GetValidCreateUserProcessWithCouponStartedEvent())

	//Assert
	verify.Should(t, spies.CacheSpy.Params["Set:key"]).Be(expectedCacheKey)
	verify.Should(t, spies.CacheSpy.Params["Set:value"]).Be(expectedCacheValue)
	verify.Should(t, spies.CacheSpy.Params["Set:time"]).Be(24 * time.Hour)
}

func Test_GivenHandler_WhenCacheSetInvoked_ThenEnsureCallsWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeError()
	expectedCacheKey := "fake-activation-data-key_coupon"
	expectedMessageError := &coupon_common.ErrorData{Code: "COUPON_REPOSITORY", Message: spies.RepoSpy.ErrorResult["GetCouponByCode"].Error()}
	expectedCacheValue := fixture.GetCachedCouponData(expectedMessageError, nil)

	// Act
	sut.Handle(fixture.GetValidCreateUserProcessWithCouponStartedEvent())

	//Assert
	verify.Should(t, spies.CacheSpy.Params["Set:key"]).Be(expectedCacheKey)
	verify.Should(t, spies.CacheSpy.Params["Set:value"]).Be(expectedCacheValue)
	verify.Should(t, spies.CacheSpy.Params["Set:time"]).Be(24 * time.Hour)
}
