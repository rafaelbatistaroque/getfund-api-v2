package sessionentity

import (
	"getfund-api-v2/internal/pkg/expect"
	fixtures "getfund-api-v2/internal/test/domain/auth/entity/sessionentity/sessionentityfixture"
	"testing"
)

func Test_GivenSessionNewInstance_WhenIdParamNullOrEmpty_ThenEnsureReturnError(t *testing.T) {
	// Arrange & Act
	_, err := fixtures.NewSessionWithInvalidId()

	//Assert
	expect.Error(t, err)
}

func Test_GivenSessionNewInstance_WhenFirstNameParamNullOrEmpty_ThenEnsureReturnError(t *testing.T) {
	// Arrange & Act
	_, err := fixtures.NewSessionWithInvalidId()

	//Assert
	expect.Error(t, err)
}

func Test_GivenSessionNewInstance_WhenRoleParamLowerThanZero_ThenEnsureReturnError(t *testing.T) {
	// Arrange & Act
	_, err := fixtures.NewSessionWithInvalidRole()

	//Assert
	expect.Error(t, err)
}

func Test_GivenSessionNewInstance_WhenSuccess_ThenEnsureReturnSessionWithPropertiesFilled(t *testing.T) {
	// Arrange & Act
	session, _ := fixtures.NewValidSession()

	//Assert
	expect.NotNil(t, session)
	expect.Equal(t, session.GetID(), fixtures.FAKE_ID)
	expect.Equal(t, session.GetFirstName(), fixtures.FAKE_FIRSTNAME)
	expect.Equal(t, session.GetIsAdmin(), fixtures.FAKE_ROLE == 1)
	expect.Equal(t, session.GetToken(), fixtures.EMPTY_STRING)
}

func Test_GivenSessionOnSetToken_WhenParamNullOrEmpty_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	session, _ := fixtures.NewValidSession()

	//Act
	err := session.SetToken(fixtures.EMPTY_STRING)

	//Assert
	expect.Error(t, err)
}

func Test_GivenSessionOnSetToken_WhenParamFilled_ThenEnsureSessionWithToken(t *testing.T) {
	// Arrange
	session, _ := fixtures.NewValidSession()

	//Act
	session.SetToken(fixtures.FAKE_TOKEN)

	//Assert
	expect.NotEmpty(t, session.GetToken())
}
