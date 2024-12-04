package authadapter

import (
	adapter "getfund-api-v2/internal/domain/auth/port/adapter"
	"getfund-api-v2/internal/pkg/verify"
	fixture "getfund-api-v2/test/domain/auth/port/adapter/recoverpasswordfixture"
	"net/http"
	"testing"
)

func Test_GivenRecoverPassword_WhenDecodeError_ThenEnsureReturnStatusBadRequestWithError(t *testing.T) {
	// Arrange
	sut := adapter.New(nil, nil, nil)
	res, req := fixture.GetHttpRequestResponse("with-body-error")

	// Act
	_, code, err := sut.RecoverPassword(res, req)

	// Assert
	verify.Should(t, code).Be(http.StatusBadRequest)
	verify.Should(t, err).NotNil()
}
