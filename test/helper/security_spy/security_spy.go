package security_spy

import (
	"errors"
	"getfund-api-v2/internal/shared/security"
)

type HasherSpy struct {
	Params              map[string]any
	ParamsByCall        map[string][]any
	CallsCount          map[string]int
	ErrorResult         map[string]error
	SuccessResult       map[string]any
	SuccessResultByCall map[string][]any
}

func New() *HasherSpy {
	return &HasherSpy{
		Params:              make(map[string]any, 2),
		ParamsByCall:        map[string][]any{},
		ErrorResult:         make(map[string]error, 1),
		SuccessResult:       make(map[string]any, 2),
		CallsCount:          make(map[string]int, 1),
		SuccessResultByCall: map[string][]any{},
	}
}

func (h *HasherSpy) HashAndMerge(input string, serverSalt []byte) string {
	h.Params["HashAndMerge:input"] = input
	h.Params["HashAndMerge:serverSalt"] = serverSalt

	h.ParamsByCall["HashAndMerge:input"] = append(h.ParamsByCall["HashAndMerge:input"], input)
	h.ParamsByCall["HashAndMerge:serverSalt"] = append(h.ParamsByCall["HashAndMerge:serverSalt"], serverSalt)

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

	h.ParamsByCall["DecryptMerged:mergedEncryptedData"] = append(h.ParamsByCall["DecryptMerged:mergedEncryptedData"], mergedEncryptedData)
	h.ParamsByCall["DecryptMerged:secretKey"] = append(h.ParamsByCall["DecryptMerged:secretKey"], secretKey)

	h.CallsCount["DecryptMerged"]++

	success := h.SuccessResult["DecryptMerged"]
	if success != nil {
		return success.(string)
	}

	return ""
}
func (h *HasherSpy) HashWithSalt(inputText string, serverSalt []byte) (string, error) {
	h.Params["HashWithSalt:inputText"] = inputText
	h.Params["HashWithSalt:serverSalt"] = serverSalt

	h.ParamsByCall["HashWithSalt:inputText"] = append(h.ParamsByCall["HashWithSalt:inputText"], inputText)
	h.ParamsByCall["HashWithSalt:serverSalt"] = append(h.ParamsByCall["HashWithSalt:serverSalt"], serverSalt)

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

	h.ParamsByCall["Encrypt:input"] = append(h.ParamsByCall["Encrypt:input"], input)
	h.ParamsByCall["Encrypt:secretKey"] = append(h.ParamsByCall["Encrypt:secretKey"], secretKey)

	h.CallsCount["Encrypt"]++

	success := h.SuccessResult["Encrypt"]
	if success != nil {
		h.SuccessResultByCall["Encrypt"] = append(h.SuccessResultByCall["Encrypt"], success.(string))
		return success.(string)
	}

	return ""
}

func (h *HasherSpy) GetRandomCode(length int) (string, error) {
	h.Params["GetRandomCode:length"] = length

	h.CallsCount["GetRandomCode"]++

	success := h.SuccessResult["GetRandomCode"]
	if success != nil {
		return success.(string), nil
	}

	return "", h.ErrorResult["GetRandomCode"]
}

func (h *HasherSpy) Hash(inputText string, serverSalt []byte) (*security.Hashing, error) {
	h.Params["Hash:inputText"] = inputText
	h.Params["Hash:serverSalt"] = serverSalt

	h.CallsCount["Hash"]++

	success := h.SuccessResult["Hash"]
	if success != nil {
		return success.(*security.Hashing), nil
	}

	return &security.Hashing{}, h.ErrorResult["Hash"]
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

func (h *HasherSpy) DefineGetRandomCodeSuccess() {
	h.SuccessResult["GetRandomCode"] = "FAKE_RANDOM_CODE"
}

func (h *HasherSpy) DefineGetRandomCodeError() {
	h.ErrorResult["GetRandomCode"] = errors.New("fake-error")
}

func (h *HasherSpy) DefineHashSuccess() {
	h.SuccessResult["Hash"] = &security.Hashing{Data: "FAKE_HASHED", Salt: "FAKE_SALT"}
}

func (h *HasherSpy) DefineHashError() {
	h.ErrorResult["Hash"] = errors.New("fake-error")
}
