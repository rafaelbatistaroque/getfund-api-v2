package activate_user_mapper_test

import (
	fixture "getfund-api-v2/test/external/domain/user/core/domain_service/active_user_mapper/active_user_mapper_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenActivateUserMapper_WhenToDto_ThenEnsureCorrectMapToActivationUserDto(t *testing.T) {
	// Arrange
	sut := fixture.NewSut()
	expectedMapped := fixture.GetUserEntity()

	// Act
	result := sut.ToDto(expectedMapped)

	// Assert
	verify.Should(t, result).NotNil()
	verify.Should(t, result.FirstName).Be(expectedMapped.GetFirstName())
	verify.Should(t, result.LastName).Be(expectedMapped.GetLastName())
	verify.Should(t, result.Email).Be(expectedMapped.GetEmail())
	verify.Should(t, result.Gender).Be(expectedMapped.GetGender())
	verify.Should(t, result.MainSocialNetwork).Be(expectedMapped.GetMainSocialNetwork())
	verify.Should(t, result.Password).Be(expectedMapped.GetPassword())
	verify.Should(t, result.RegisteredUrl).Be(expectedMapped.GetRegisteredUrl())
	verify.Should(t, result.IsActive).Be(expectedMapped.GetIsActive())
	verify.Should(t, result.IsAdmin).Be(expectedMapped.GetIsAdmin())
	verify.Should(t, result.UserCategoryId).Be(expectedMapped.GetUserCategoryId())
	verify.Should(t, result.CountryId).Be(expectedMapped.GetCountryId())
}
