package dto

// MerchantAccountRequest 商户账户列表
type MerchantAccountRequest struct {
	PageRequest
	MerchantName string `query:"username" form:"username"`
	Status       string `query:"status" form:"status"`
}

type MerchantAccountCurrencyRequest struct {
	PageRequest
	MId      string `query:"mId" form:"mId"`
	Currency string `query:"currency" form:"currency"`
}
