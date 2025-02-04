package user_gateway_test

import (
	"getfund-api-v2/internal/shared/result_app"
	fixture "getfund-api-v2/test/internal/domain/user/adapter/gateway/activate_user_gateway_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenActivateUser_WhenDecodeError_ThenEnsureReturnBadRequestWithError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	_, code, err := sut.ActivateUser(res, req)

	// Assert
	verify.Should(t, code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err.Error()).Be("activation code is required")
}
