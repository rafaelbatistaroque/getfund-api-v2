package activate_user_test

import (
	"encoding/json"
	"fmt"
	"getfund-api-v2/internal/domain/user/core/dto/user_dto"
	payload "getfund-api-v2/internal/domain/user/core/dto/user_payload"
	"getfund-api-v2/internal/domain/user/core/entity/user_entity"
	"getfund-api-v2/internal/domain/user/core/usecase/activate_user"
	"getfund-api-v2/internal/shared/result_app"
	fixture "getfund-api-v2/test/internal/domain/user/core/usecase/activate_user/activate_user_fixture"
	"testing"
	"time"

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

func Test_GivenExecute_WhenUnmarshalSuccess_ThenEnsureCallUserExistsByUsernameWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(fixture.GetUserDataWithCouponSerialized())
	spies.RepoSpy.DefineCreateUserSuccess()
	var expectedParam = user_dto.ActivationUserData{}
	json.Unmarshal([]byte(spies.CacheSpy.SuccessResult["Get"].(string)), &expectedParam)

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.RepoSpy.Params["UserExistsByUsername:username"]).Be(expectedParam.Email)
}

func Test_GivenExecute_WhenUserExistsByUsernameInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(fixture.GetUserDataWithCouponSerialized())
	spies.RepoSpy.DefineCreateUserSuccess()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.RepoSpy.CallsCount["UserExistsByUsername"]).Be(1)
}

func Test_GivenExecute_WhenUserExistsByUsernameError_ThenEnsureReturnNotFoundError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess("")
	spies.RepoSpy.DefineUserExistsByUsernameError()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, err.Message).Be(spies.RepoSpy.ErrorResult["UserExistsByUsername"])
}

func Test_GivenExecute_WhenUserExistsByUsernameFound_ThenEnsureReturnDuplicateEntryError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess("")
	spies.RepoSpy.DefineUserExistsByUsernameSuccess()
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

func Test_GivenExecute_WhenUserExistsByUsernameNotFound_ThenEnsureCallMapperToDtoWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(fixture.GetUserDataWithCouponSerialized())
	spies.RepoSpy.DefineCreateUserSuccess()
	var expectedParam = user_dto.ActivationUserData{}
	json.Unmarshal([]byte(spies.CacheSpy.SuccessResult["Get"].(string)), &expectedParam)

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	entityParam := spies.MapperSpy.Params["ToDto:entity"].(*user_entity.User)
	verify.Should(t, entityParam.GetFirstName()).Be(expectedParam.FirstName)
	verify.Should(t, entityParam.GetLastName()).Be(expectedParam.LastName)
	verify.Should(t, entityParam.GetEmail()).Be(expectedParam.Email)
	verify.Should(t, entityParam.GetUsername()).Be(expectedParam.Email)
	verify.Should(t, entityParam.GetGender()).Be(expectedParam.Gender)
	verify.Should(t, entityParam.GetPassword()).Be(expectedParam.Password)
	verify.Should(t, entityParam.GetCountryId()).Be(expectedParam.CountryId)
	verify.Should(t, entityParam.GetUserCategoryId()).Be(expectedParam.UserCategoryId)
	verify.Should(t, entityParam.GetMainSocialNetwork()).Be(expectedParam.MainSocialNetwork)
	verify.Should(t, entityParam.GetRegisteredUrl()).Be(expectedParam.RegisteredUrl)
	verify.Should(t, entityParam.GetIsActive()).BeTrue()
	verify.Should(t, entityParam.GetIsAdmin()).BeFalse()
	verify.Should(t, entityParam.GetRegisteredAt()).NotNil()
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
	inputValid := fixture.GetInput()
	expectedPaylod := &payload.ActivateUserWithCouponConfirmedPayload{
		ActivationDataKey: inputValid.ActivationDataKey,
		UserId:            spies.RepoSpy.SuccessResult["CreateUser"].(*user_dto.UserDto).Id,
	}

	// Act
	sut.Execute(inputValid)

	// Assert
	verify.Should(t, spies.BusSpy.Params["EmitWithPayload:event"][0]).Be(&activate_user.ActivateUserWithCouponConfirmedEvent{})
	verify.Should(t, spies.BusSpy.Params["EmitWithPayload:payload"][0]).Be(expectedPaylod)
}

func Test_GivenExecute_WhenEmitWithPayloadWithUserActivationWithCouponConfirmedEventInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(fixture.GetUserDataWithCouponSerialized())
	spies.RepoSpy.DefineCreateUserSuccess()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.BusSpy.CallsCount["EmitWithPayload"]).Be(2)
}

func Test_GivenExecute_WhenUserSaved_ThenEnsureCallPublishWithPayloadWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(fixture.GetUserDataWithoutCouponSerialized())
	spies.RepoSpy.DefineCreateUserSuccess()
	var expectedParam = user_dto.ActivationUserData{}
	json.Unmarshal([]byte(spies.CacheSpy.SuccessResult["Get"].(string)), &expectedParam)
	expectedPayload := &payload.ActivateUserConfirmedPayload{
		Username: expectedParam.Email,
		Password: expectedParam.Password,
		Id:       spies.RepoSpy.SuccessResult["CreateUser"].(*user_dto.UserDto).Id,
	}

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	payloadReceived := spies.BusSpy.Params["EmitWithPayloadAndResponse:payload"][0].(*payload.ActivateUserConfirmedPayload)
	verify.Should(t, spies.BusSpy.Params["EmitWithPayloadAndResponse:event"][0]).Be(&activate_user.ActivateUserConfirmedEvent{})
	verify.Should(t, payloadReceived.Id).Be(expectedPayload.Id)
	verify.Should(t, payloadReceived.Username).Be(expectedPayload.Username)
	verify.Should(t, payloadReceived.Password).Be(expectedPayload.Password)
}

func Test_GivenExecute_WhenEmitWithPayloadWithUserCreatedEventInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(fixture.GetUserDataWithoutCouponSerialized())
	spies.RepoSpy.DefineCreateUserSuccess()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.BusSpy.CallsCount["EmitWithPayload"]).Be(1)
}

func Test_GivenExecute_WhenResponsePublishWithPayloadTimeout_ThenEnsureReturnServerErrorWithAppropriateMessage(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(fixture.GetUserDataWithoutCouponSerialized())
	spies.RepoSpy.DefineCreateUserSuccess()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, err.Message.Error()).Be("error on get session [timeout]")
}

func Test_GivenExecute_WhenEmitWithPayloadResponseEmpty_ThenEnsureReturnServerErrorWithAppropriateMessage(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(fixture.GetUserDataWithoutCouponSerialized())
	spies.RepoSpy.DefineCreateUserSuccess()
	spies.SettingsSpy.SetTimeoutResponseEvent(2)
	errChannel := make(chan *result_app.ApplicationError, 1)

	// Act
	go func() {
		_, err := sut.Execute(fixture.GetInput())
		errChannel <- err
	}()
	time.Sleep(1 * time.Second)
	responseChannel := spies.BusSpy.Params["EmitWithPayloadAndResponse:responseChannel"][0].(chan []byte)
	responseChannel <- []byte("")
	close(responseChannel)

	// Assert
	errUnwrapped := <-errChannel
	verify.Should(t, errUnwrapped.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, errUnwrapped.Message.Error()).Be("error on get session [response empty]")
}

func Test_GivenExecute_WhenEmitWithPayloadResponseInvalid_ThenEnsureReturnServerErrorWithAppropriateMessage(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(fixture.GetUserDataWithoutCouponSerialized())
	spies.RepoSpy.DefineCreateUserSuccess()
	spies.SettingsSpy.SetTimeoutResponseEvent(2)
	errChannel := make(chan *result_app.ApplicationError, 1)

	// Act
	go func() {
		_, err := sut.Execute(fixture.GetInput())
		errChannel <- err
	}()
	time.Sleep(1 * time.Second)
	responseChannel := spies.BusSpy.Params["EmitWithPayloadAndResponse:responseChannel"][0].(chan []byte)
	responseChannel <- []byte("{invalid-response}")
	close(responseChannel)

	// Assert
	errUnwrapped := <-errChannel
	verify.Should(t, errUnwrapped.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, errUnwrapped.Message.Error()).Be("error on get session [response invalid]")
}

func Test_GivenExecute_WhenEmitWithPayloadResponseSuccess_ThenEnsureReturnSession(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(fixture.GetUserDataWithoutCouponSerialized())
	spies.RepoSpy.DefineCreateUserSuccess()
	spies.SettingsSpy.SetTimeoutResponseEvent(2)
	resultChannel := make(chan *activate_user.Output, 1)
	responseEvent, expectedOutput := fixture.GetResponseSession()

	// Act
	go func() {
		result, _ := sut.Execute(fixture.GetInput())
		resultChannel <- result
	}()
	time.Sleep(1 * time.Second)
	responseChannel := spies.BusSpy.Params["EmitWithPayloadAndResponse:responseChannel"][0].(chan []byte)
	responseChannel <- responseEvent
	close(responseChannel)

	// Assert
	resultUnwrapped := <-resultChannel
	verify.Should(t, resultUnwrapped).NotNil()
	verify.Should(t, resultUnwrapped.Token).Be(expectedOutput.Token)
	verify.Should(t, resultUnwrapped.Session.ID).Be(expectedOutput.Session.ID)
	verify.Should(t, resultUnwrapped.Session.FirstName).Be(expectedOutput.Session.FirstName)
	verify.Should(t, resultUnwrapped.Session.IsAdmin).Be(expectedOutput.Session.IsAdmin)
}
