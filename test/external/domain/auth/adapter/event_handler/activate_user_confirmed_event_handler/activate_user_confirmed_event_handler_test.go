package activate_user_confirmed_event_handler_test

import (
	"bytes"
	"encoding/json"
	"getfund-api-v2/internal/domain/auth/core/usecase/signin"
	"getfund-api-v2/internal/shared/app_constant"
	fixture "getfund-api-v2/test/external/domain/auth/adapter/event_handler/activate_user_confirmed_event_handler/activate_user_confirmed_event_handler_fixture"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenHandler_WhenPayloadParseError_ThenEnsureNeverCallGetCache(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Handle(fixture.GetInvalidActivateUserConfirmedEvent())

	//Assert
	verify.Should(t, spies.SigninSpy.CallsCount["Execute"]).Be(0)
}

func Test_GivenHandler_WhenPayloadParseSuccess_ThenEnsureCallSigninWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedInput := &signin.Input{
		UserName: uuid.NewString(),
		Password: uuid.NewString(),
	}

	// Act
	sut.Handle(fixture.GetValidActivateUserConfirmedEvent(expectedInput.UserName, expectedInput.Password))

	//Assert
	verify.Should(t, spies.SigninSpy.Params["Execute:input"]).Be(expectedInput)
}

func Test_GivenHandler_WhenUsecaseInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Handle(fixture.GetValidActivateUserConfirmedEvent("", ""))

	//Assert
	verify.Should(t, spies.SigninSpy.CallsCount["Execute"]).Be(1)
}

func Test_GivenHandler_WhenUsecaseError_ThenEnsureDefineChannelResponseWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.SigninSpy.DefineSigninError()
	validEvent := fixture.GetValidActivateUserConfirmedEvent("", "")
	channelResponse := make(chan []byte, 1)
	validEvent.SetChannel(channelResponse)

	// Act
	go sut.Handle(validEvent)

	//Assert
	select {
	case response := <-channelResponse:
		verify.Should(t, bytes.Equal(response, app_constant.EMPTYB)).BeTrue()
	case <-time.After(time.Duration(3) * time.Second):
		verify.Should(t, true).BeFalse().Message("time out channel response")
	}
}

func Test_GivenHandler_WhenUsecaseSuccess_ThenEnsureDefineChannelResponseWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedSigninOutput := fixture.GetSigninOutput()
	spies.SigninSpy.DefineSigninSuccessWithValue(expectedSigninOutput)
	validEvent := fixture.GetValidActivateUserConfirmedEvent("", "")
	channelResponse := make(chan []byte, 1)
	validEvent.SetChannel(channelResponse)

	// Act
	go sut.Handle(validEvent)

	//Assert
	select {
	case response := <-channelResponse:
		verify.Should(t, bytes.Equal(response, app_constant.EMPTYB)).BeFalse()
		var output = &signin.Output{}
		if err := json.Unmarshal(response, output); err != nil {
			verify.Should(t, err).NotNil()
		}
		verify.Should(t, output.Token).Be(expectedSigninOutput.Token)
		verify.Should(t, output.Session.ID).Be(expectedSigninOutput.Session.ID)
		verify.Should(t, output.Session.FirstName).Be(expectedSigninOutput.Session.FirstName)
		verify.Should(t, output.Session.IsAdmin).Be(expectedSigninOutput.Session.IsAdmin)
	case <-time.After(time.Duration(3) * time.Second):
		verify.Should(t, true).BeFalse().Message("time out channel response")
	}
}
