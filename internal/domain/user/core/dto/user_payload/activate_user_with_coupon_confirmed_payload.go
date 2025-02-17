package payload

type ActivateUserWithCouponConfirmedPayload struct {
	UserId            int    `json:"iser_id"`
	ActivationCode    string `json:"activation_code"`
	ActivationDataKey string `json:"activation_data_key"`
}
