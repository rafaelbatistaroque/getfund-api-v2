package activate_user_test

import (
	"encoding/json"
	"fmt"
	"getfund-api-v2/internal/domain/user/core/entity/activate_user_entity"
	"getfund-api-v2/internal/domain/user/core/user_dto"
	"getfund-api-v2/internal/shared/result_app"
	fixture "getfund-api-v2/test/internal/domain/user/core/usecase/activate_user/activate_user_fixture"
	"testing"

	"github.com/rafaelbatistaroque/validation"
	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenExecute_WhenActivationCodeEmpty_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	inputWithActivationCodeEmpty := fixture.GetInput(fixture.WithEmptyActivationCode())

	// Act
	_, err := sut.Execute(inputWithActivationCodeEmpty)

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAUTHORIZED_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "ActivationCode"))
}

func Test_GivenExecute_WhenActivationCodeInvalid_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	inputWithActivationCodeInvalid := fixture.GetInput(fixture.WithInvalidActivationCodeLength())

	// Act
	_, err := sut.Execute(inputWithActivationCodeInvalid)

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAUTHORIZED_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_SHOULD_HAVE_EXACTLY_CHARACTER.Error(), "ActivationCode", 20))
}

func Test_GivenExecute_WhenInputValid_ThenEnsureCallCacheGetWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	validInput := fixture.GetInput()
	expectedParam := "user_activation_" + validInput.ActivationCode

	// Act
	sut.Execute(validInput)

	// Assert
	verify.Should(t, spies.CacheSpy.Params["Get:key"]).Be(expectedParam)
}

func Test_GivenExecute_WhenCacheGetInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.CacheSpy.CallsCount["Get"]).Be(1)
}

func Test_GivenExecute_WhenCacheGetError_ThenEnsureReturnNotFoundErrorWIthApropriateMessage(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetError()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.NOT_FOUND_CODE)
	verify.Should(t, err.Message.Error()).Be("activation code not found")
}

func Test_GivenExecute_WhenUnmarshalError_ThenEnsureReturnAppropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(`{error-data}`)

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, err.Message.Error()).Be("error on get user data")
}

func Test_GivenExecute_WhenUnmarshalSuccess_ThenEnsureCallGetUserByUsernameWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(`{"email": "fake-valid@mail.com"}`)
	var expectedParam = user_dto.ActivationUserData{}
	json.Unmarshal([]byte(spies.CacheSpy.SuccessResult["Get"].(string)), &expectedParam)

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.RepoSpy.Params["GetUserByUsername:username"]).Be(expectedParam.Email)
}

func Test_GivenExecute_WhenGetUserByUsernameInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess("")

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.RepoSpy.CallsCount["GetUserByUsername"]).Be(1)
}

func Test_GivenExecute_WhenGetUserByUsernameError_ThenEnsureReturnNotFoundError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess("")
	spies.RepoSpy.DefineGetUserByUsernameError()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, err.Message).Be(spies.RepoSpy.ErrorResult["GetUserByUsername"])
}

func Test_GivenExecute_WhenGetUserByUsernameFound_ThenEnsureReturnDuplicateEntryError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess("")
	spies.RepoSpy.DefineGetUserByUsernameSuccess()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.DUPLICATED_ENTRY_CODE)
	verify.Should(t, err.Message.Error()).Be("user already exists")
	verify.Should(t, spies.CacheSpy.CallsCount["Delete"]).Be(1)
}

func Test_GivenExecute_WhenGetUserByUsernameNotFound_ThenEnsureMapperCallToEntityWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(fixture.GetUserDataWithCouponSerialized())
	var expectedParam = user_dto.ActivationUserData{}
	json.Unmarshal([]byte(spies.CacheSpy.SuccessResult["Get"].(string)), &expectedParam)

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.MapperSpy.Params["ToEntity:data"]).Be(&expectedParam)
}

func Test_GivenExecute_WhenToEntityInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess("")

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.MapperSpy.CallsCount["ToEntity"]).Be(1)
}

func Test_GivenExecute_WhenGetUserByUsernameNotFound_ThenEnsureCallMapperToDtoWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(fixture.GetUserDataWithCouponSerialized())
	var expectedParam = user_dto.ActivationUserData{}
	json.Unmarshal([]byte(spies.CacheSpy.SuccessResult["Get"].(string)), &expectedParam)

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	entityParam := spies.MapperSpy.Params["ToDto:entity"].(*activate_user_entity.ActivationUser)
	verify.Should(t, entityParam.GetFirstName()).Be(expectedParam.FirstName)
	verify.Should(t, entityParam.GetLastName()).Be(expectedParam.LastName)
	verify.Should(t, entityParam.GetEmail()).Be(expectedParam.Email)
	verify.Should(t, entityParam.GetGender()).Be(expectedParam.Gender)
	verify.Should(t, entityParam.GetPassword()).Be(expectedParam.Password)
	verify.Should(t, entityParam.GetCountryId()).Be(expectedParam.CountryId)
	verify.Should(t, entityParam.GetUserCategoryId()).Be(expectedParam.UserCategoryId)
	verify.Should(t, entityParam.GetMainSocialNetwork()).Be(expectedParam.MainSocialNetwork)
	verify.Should(t, entityParam.GetRegisteredUrl()).Be(expectedParam.RegisteredUrl)
	verify.Should(t, entityParam.GetIsActive()).Be(1)
	verify.Should(t, entityParam.GetIsAdmin()).Be(0)
	verify.Should(t, entityParam.GetRegisteredAt()).NotNil()
}
