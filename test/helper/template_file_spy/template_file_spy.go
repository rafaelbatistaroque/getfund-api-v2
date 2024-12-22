package template_file_spy

import "errors"

type TemplateFileSpy struct {
	Params     map[string]string
	CallsCount map[string]int

	SuccessResult map[string]interface{}
	ErrorResult   map[string]error
}

func New() *TemplateFileSpy {
	return &TemplateFileSpy{
		Params:        make(map[string]string),
		CallsCount:    make(map[string]int),
		SuccessResult: make(map[string]interface{}),
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

func (t *TemplateFileSpy) DefineGetRecoveryPasswordTemplateError() {
	t.ErrorResult["GetRecoveryPasswordTemplate"] = errors.New("fake-error")
}
