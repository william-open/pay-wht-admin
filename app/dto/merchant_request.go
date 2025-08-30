package dto

// SaveMerchant 保存商户
type SaveMerchant struct {
	MId               int    `json:"mId" binding:"required"`
	Username          string `json:"username"`
	Password          string `json:"password"`
	Nickname          string `json:"nickname"`
	CallbackSecretKey string `json:"callbackSecretKey"`
	NotifyUrl         string `json:"notifyUrl"`
	AesSecretKey      string `json:"aesSecretKey"`
	PublicKey         string `json:"publicKey"`
	PrivateKey        string `json:"PrivateKey"`
	UserType          string `json:"userType"`
	ApiKey            string `json:"apiKey"`
	AppId             string `json:"appId"`
	Status            string `json:"status"`
	PayType           string `json:"payType"`
	CreateBy          string `json:"createBy"`
	UpdateBy          string `json:"updateBy"`
	Remark            string `json:"remark"`
	Ways              string `json:"ways"`
	UpstreamId        string `json:"upstreamId"`
	PayPwd            string `json:"payPwd"`
}

// ResetMerchantGoogleSecret 重置谷歌验证码
type ResetMerchantGoogleSecret struct {
	MId int `json:"mId" binding:"required"`
}

// MerchantListRequest 商户列表
type MerchantListRequest struct {
	PageRequest
	MerchantName string `query:"username" form:"username"`
	Status       string `query:"status" form:"status"`
}

// CreateMerchantRequest 新增商户
type CreateMerchantRequest struct {
	MId               int    `json:"mId""`
	Username          string `json:"username" binding:"required"`
	Password          string `json:"password" binding:"required"`
	Nickname          string `json:"nickname" binding:"required"`
	CallbackSecretKey string `json:"callbackSecretKey"`
	//NotifyUrl         string `json:"notifyUrl"`
	Status     string `json:"status"`
	PayType    string `json:"payType"`
	UserType   string `json:"userType"`
	UpstreamId string `json:"upstreamId"`
	Ways       string `json:"ways"`
	Remark     string `json:"remark"`
}

// UpdateMerchantRequest 更新商户
type UpdateMerchantRequest struct {
	MId      int    `json:"mId" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password"`
	Nickname string `json:"nickname" binding:"required"`
	//CallbackSecretKey string `json:"callbackSecretKey" binding:"required"`
	//NotifyUrl         string `json:"notifyUrl" binding:"required"`
	Status string `json:"status"`
	Remark string `json:"remark"`
}

// UpdateMerchantPwdRequest 更新商户密码
type UpdateMerchantPwdRequest struct {
	MId      int    `json:"mId" binding:"required"`
	LoginPwd string `json:"loginPwd"`
	PayPwd   string `json:"payPwd"`
}

// UpdateWhitelistRequest 更新白名单
type UpdateWhitelistRequest struct {
	MId        int    `json:"mId" binding:"required"`
	ApiIp      string `json:"apiIp" binding:"required"`
	LoginApiIp string `json:"loginApiIp" binding:"required"`
	ApiDomain  string `json:"apiDomain" binding:"required"`
}

// ResetMerchantGoogleSecretRequest 重置谷歌验证码
type ResetMerchantGoogleSecretRequest struct {
	MId int `json:"mId" binding:"required"`
}
