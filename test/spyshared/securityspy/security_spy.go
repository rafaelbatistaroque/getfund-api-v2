package securityspy

import "errors"

type HasherSpy struct {
	Params        map[string]interface{}
	CallsCount    map[string]int
	ErrorResult   map[string]error
	SuccessResult map[string]interface{}
}

func New() *HasherSpy {
	return &HasherSpy{Params: make(map[string]interface{}, 2), ErrorResult: make(map[string]error, 1), SuccessResult: make(map[string]interface{}, 2), CallsCount: make(map[string]int, 1)}
}

func (h *HasherSpy) HashAndMerge(input string, serverSalt []byte) string {
	h.Params["HashAndMerge:input"] = input
	h.Params["HashAndMerge:serverSalt"] = serverSalt

	h.CallsCount["HashAndMerge"]++

	success := h.SuccessResult["HashAndMerge"]
	if success != nil {
		return success.(string)
	}

	return ""
}
func (h *HasherSpy) DecryptMerged(mergedEncryptedData string, secretKey []byte) string {
	h.Params["DecryptMerged:mergedEncryptedData"] = mergedEncryptedData
	h.Params["DecryptMerged:secretKey"] = secretKey

	success := h.SuccessResult["DecryptMerged"]
	if success != nil {
		return success.(string)
	}

	return ""
}
func (h *HasherSpy) HashWithSalt(inputText string, serverSalt []byte) (string, error) {
	h.Params["HashWithSalt:inputText"] = inputText
	h.Params["HashWithSalt:serverSalt"] = serverSalt

	h.CallsCount["HashWithSalt"]++

	success := h.SuccessResult["HashWithSalt"]
	if success != nil {
		return success.(string), h.ErrorResult["HashWithSalt"]
	}

	return "", h.ErrorResult["HashWithSalt"]
}

func (h *HasherSpy) IsMatch(inputHashed, inputText string, serverSalt []byte) bool {
	h.Params["IsMatch:inputHashed"] = inputHashed
	h.Params["IsMatch:inputText"] = inputText
	h.Params["IsMatch:serverSalt"] = serverSalt

	h.CallsCount["IsMatch"]++

	result := h.SuccessResult["IsMatch"]
	if result != nil {
		return result.(bool)
	}

	return false
}

func (h *HasherSpy) Encrypt(input string, secretKey []byte) string {
	h.Params["Encrypt:input"] = input
	h.Params["Encrypt:secretKey"] = secretKey

	h.CallsCount["Encrypt"]++

	success := h.SuccessResult["Encrypt"]
	if success != nil {
		return success.(string)
	}

	return ""
}

func (h *HasherSpy) DefineHashWithSaltError() {
	h.ErrorResult["HashWithSalt"] = errors.New("fake-error")
}

func (h *HasherSpy) DefineHashWithSaltSuccess(result string) {
	h.SuccessResult["HashWithSalt"] = result
}

func (h *HasherSpy) DefineIsMatchError() {
	h.SuccessResult["IsMatch"] = false
}

func (h *HasherSpy) DefineIsMatchSuccess() {
	h.SuccessResult["IsMatch"] = true
}

func (h *HasherSpy) DefineEncryptSuccess() {
	h.SuccessResult["Encrypt"] = "FAKE_ENCRYPT_AJS6YFL284NF61305J4B"
}

func (h *HasherSpy) DefineHashAndMergeSuccess(result string) {
	h.SuccessResult["HashAndMerge"] = result
}

func (h *HasherSpy) DefineDecryptMergedSuccess(result string) {
	h.SuccessResult["DecryptMerged"] = result
}
