package template_file_spy

import "errors"

type TemplateFileSpy struct {
	Params     map[string]string
	CallsCount map[string]int

	SuccessResult map[string]any
	ErrorResult   map[string]error
}

func New() *TemplateFileSpy {
	return &TemplateFileSpy{
		Params:        make(map[string]string),
		CallsCount:    make(map[string]int),
		SuccessResult: make(map[string]any),
		ErrorResult:   make(map[string]error),
	}
}

func (t *TemplateFileSpy) GetRecoveryPasswordTemplate() (string, error) {
	t.CallsCount["GetRecoveryPasswordTemplate"]++

	success := t.SuccessResult["GetRecoveryPasswordTemplate"]
	if success != nil {
		return success.(string), nil
	}

	return "", t.ErrorResult["GetRecoveryPasswordTemplate"]
}

func (t *TemplateFileSpy) DefineGetRecoveryPasswordTemplateSuccess() {
	t.SuccessResult["GetRecoveryPasswordTemplate"] = "<div>{{first_name}}</div><div>{{recovery_link}}</div>"
}

func (t *TemplateFileSpy) GetRecoveryPasswordTemplateReplaced(firstName, recoveryLink string) string {
	return "<div>" + firstName + "</div><div>" + recoveryLink + "</div>"
}

func (t *TemplateFileSpy) DefineGetRecoveryPasswordTemplateError() {
	t.ErrorResult["GetRecoveryPasswordTemplate"] = errors.New("fake-error")
}

func (t *TemplateFileSpy) GetActivationAccountTemplate() (string, error) {
	t.CallsCount["GetActivationAccountTemplate"]++

	success := t.SuccessResult["GetActivationAccountTemplate"]
	if success != nil {
		return success.(string), nil
	}

	return "", t.ErrorResult["GetActivationAccountTemplate"]
}

func (t *TemplateFileSpy) DefineGetActivationAccountTemplateSuccess() {
	t.SuccessResult["GetActivationAccountTemplate"] = "<div>{{first_name}}</div><div>{{activation_link}}</div><div>{{activation_link}}</div>"
}

func (t *TemplateFileSpy) GetGetActivationAccountTemplateReplaced(firstName, activationLink string) string {
	return "<div>" + firstName + "</div><div>" + activationLink + "</div>" + "<div>" + activationLink + "</div>"
}

func (t *TemplateFileSpy) DefineGetActivationAccountTemplateError() {
	t.ErrorResult["GetActivationAccountTemplate"] = errors.New("fake-error")
}
