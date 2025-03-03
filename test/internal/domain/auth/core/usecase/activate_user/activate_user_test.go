package activate_user_test

import (
	"encoding/json"
	"fmt"
	"getfund-api-v2/internal/domain/auth/core/auth_dto"
	"getfund-api-v2/internal/domain/auth/core/auth_dto/auth_payload"
	"getfund-api-v2/internal/domain/auth/core/entity/user_entity"
	"getfund-api-v2/internal/domain/auth/core/usecase/activate_user"
	"getfund-api-v2/internal/shared/result_app"
	fixture "getfund-api-v2/test/internal/domain/auth/core/usecase/activate_user/activate_user_fixture"
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

func Test_GivenExecute_WhenActivationDataKeyEmpty_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	inputWithActivationCodeEmpty := fixture.GetInput(fixture.WithEmptyActivationDataKey())

	// Act
	_, err := sut.Execute(inputWithActivationCodeEmpty)

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAUTHORIZED_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "ActivationDataKey"))
}

func Test_GivenExecute_WhenActivationDataKeyInvalid_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	inputWithActivationCodeInvalid := fixture.GetInput(fixture.WithInvalidActivationDataKey())

	// Act
	_, err := sut.Execute(inputWithActivationCodeInvalid)

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAUTHORIZED_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_INVALID.Error(), "ActivationDataKey"))
}

func Test_GivenExecute_WhenInputValid_ThenEnsureCallCacheGetWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	validInput := fixture.GetInput()

	// Act
	sut.Execute(validInput)

	// Assert
	verify.Should(t, spies.CacheSpy.Params["Get:key"]).Be(validInput.ActivationDataKey)
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

func Test_GivenExecute_WhenUnmarshalSuccess_ThenEnsureCallUserExistsWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(fixture.GetUserDataWithCouponSerialized())
	spies.RepoSpy.DefineCreateUserSuccess()
	var expectedParam = auth_dto.ActivationUserData{}
	json.Unmarshal([]byte(spies.CacheSpy.SuccessResult["Get"].(string)), &expectedParam)

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.RepoSpy.Params["UserExists:username"]).Be(expectedParam.Username)
}

func Test_GivenExecute_WhenUserExistsInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(fixture.GetUserDataWithCouponSerialized())
	spies.RepoSpy.DefineCreateUserSuccess()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.RepoSpy.CallsCount["UserExists"]).Be(1)
}

func Test_GivenExecute_WhenUserExistsError_ThenEnsureReturnNotFoundError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess("")
	spies.RepoSpy.DefineUserExistsError()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, err.Message).Be(spies.RepoSpy.ErrorResult["UserExists"])
}

func Test_GivenExecute_WhenUserExistsFound_ThenEnsureReturnDuplicateEntryError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess("")
	spies.RepoSpy.DefineUserExistsSuccessUserFound()
	validInput := fixture.GetInput()
	expectedParam := "user_activation_" + validInput.ActivationCode

	// Act
	_, err := sut.Execute(validInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.DUPLICATED_ENTRY_CODE)
	verify.Should(t, err.Message.Error()).Be("user already exists")
	verify.Should(t, spies.CacheSpy.Params["Delete:key"]).Be(expectedParam)
	verify.Should(t, spies.CacheSpy.CallsCount["Delete"]).Be(1)
}

func Test_GivenExecute_WhenUserExistsNotFound_ThenEnsureCallMapperToDtoWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(fixture.GetUserDataWithCouponSerialized())
	spies.RepoSpy.DefineCreateUserSuccess()
	var expectedParam = auth_dto.ActivationUserData{}
	json.Unmarshal([]byte(spies.CacheSpy.SuccessResult["Get"].(string)), &expectedParam)

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	entityParam := spies.MapperSpy.Params["ToDto:entity"].(*user_entity.User)
	verify.Should(t, entityParam.GetFirstName()).Be(expectedParam.FirstName)
	verify.Should(t, entityParam.GetLastName()).Be(expectedParam.LastName)
	verify.Should(t, entityParam.GetUsername()).Be(expectedParam.Username)
	verify.Should(t, entityParam.GetPassword()).Be(expectedParam.Password)
	verify.Should(t, entityParam.GetIsActive()).BeTrue()
	verify.Should(t, entityParam.GetIsAdmin()).BeFalse()
	verify.Should(t, entityParam.GetCreatedAt()).NotNil()
	verify.Should(t, entityParam.GetUpdatedAt()).NotNil()
}

func Test_GivenExecute_WhenToDtoInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(fixture.GetUserDataWithCouponSerialized())
	spies.RepoSpy.DefineCreateUserSuccess()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.MapperSpy.CallsCount["ToDto"]).Be(1)
}

func Test_GivenExecute_WhenToDtoSuccess_ThenEnsureCallCreateUserWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(fixture.GetUserDataWithCouponSerialized())
	spies.MapperSpy.DefineToDtoSuccess(fixture.GetActivateUserEntity())
	spies.RepoSpy.DefineCreateUserSuccess()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.RepoSpy.Params["CreateUser:user"]).Be(spies.MapperSpy.SuccessResult["ToDto"])
}

func Test_GivenExecute_WhenCreateUserInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(fixture.GetUserDataWithCouponSerialized())
	spies.RepoSpy.DefineCreateUserSuccess()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.RepoSpy.CallsCount["CreateUser"]).Be(1)
}

func Test_GivenExecute_WhenCreateUserError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(fixture.GetUserDataWithCouponSerialized())
	spies.RepoSpy.DefineCreateUserError()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, err.Message).Be(spies.RepoSpy.ErrorResult["CreateUser"])
}

func Test_GivenExecute_WhenCreateUserSuccess_ThenEnsureCallCacheDeleteWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(fixture.GetUserDataWithCouponSerialized())
	spies.RepoSpy.DefineCreateUserSuccess()
	validInput := fixture.GetInput()
	expectedParam := "user_activation_" + validInput.ActivationCode

	// Act
	sut.Execute(validInput)

	// Assert
	verify.Should(t, spies.CacheSpy.Params["Delete:key"]).Be(expectedParam)
	verify.Should(t, spies.CacheSpy.CallsCount["Delete"]).Be(1)
}

func Test_GivenExecute_WhenUserSavedAndThereIsCouponCode_ThenEnsureCallPublishWithPayloadWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(fixture.GetUserDataWithCouponSerialized())
	spies.RepoSpy.DefineCreateUserSuccess()
	userData := fixture.GetUserDataWithCoupon()
	expectedPaylod := &auth_payload.ActivateUserWithCouponConfirmedPayload{
		UserId:     spies.RepoSpy.SuccessResult["CreateUser"].(*auth_dto.UserDto).Id,
		CouponCode: userData.CouponCode,
		Email:      userData.Username,
	}

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.BusSpy.Params["EmitWithPayload:event"][0]).Be(&activate_user.ActivateUserWithCouponConfirmedEvent{})
	verify.Should(t, spies.BusSpy.Params["EmitWithPayload:payload"][0]).Be(expectedPaylod)
}

func Test_GivenExecute_WhenEmitWithPayloadWithActivateUserWithCouponConfirmedEventInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(fixture.GetUserDataWithCouponSerialized())
	spies.RepoSpy.DefineCreateUserSuccess()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.BusSpy.CallsCount["EmitWithPayload"]).Be(1)
}

func Test_GivenExecute_WhenUserSaved_ThenEnsureReturnCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(fixture.GetUserDataWithoutCouponSerialized())
	spies.RepoSpy.DefineCreateUserSuccess()
	expectedOutput := fixture.GetOutput()

	// Act
	result, _ := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, result.Username).Be(expectedOutput.Username)
	verify.Should(t, result.Password).Be(expectedOutput.Password)
}
