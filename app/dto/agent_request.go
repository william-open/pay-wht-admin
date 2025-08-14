package dto

// SaveAgent 保存代理
type SaveAgent struct {
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
}

// AgentListRequest 代理列表
type AgentListRequest struct {
	PageRequest
	AgentName string `query:"username" form:"username"`
	Status    string `query:"status" form:"status"`
	UserType  int64  `query:"userType" form:"userType"`
}

// CreateAgentRequest 新增代理
type CreateAgentRequest struct {
	MId               int    `json:"mId""`
	Username          string `json:"username" binding:"required"`
	Password          string `json:"password" binding:"required"`
	Nickname          string `json:"nickname" binding:"required"`
	CallbackSecretKey string `json:"callbackSecretKey"`
	NotifyUrl         string `json:"notifyUrl"`
	Status            string `json:"status"`
	PayType           string `json:"payType"`
	UserType          string `json:"userType"`
	UpstreamId        string `json:"upstreamId"`
	Ways              string `json:"ways"`
	Remark            string `json:"remark"`
}

// UpdateAgentRequest 更新代理
type UpdateAgentRequest struct {
	MId               int    `json:"mId" binding:"required"`
	Username          string `json:"username" binding:"required"`
	Password          string `json:"password"`
	Nickname          string `json:"nickname" binding:"required"`
	CallbackSecretKey string `json:"callbackSecretKey" binding:"required"`
	NotifyUrl         string `json:"notifyUrl" binding:"required"`
	Status            string `json:"status"`
	Remark            string `json:"remark"`
}

// UpdateAgentWhitelistRequest 更新白名单
type UpdateAgentWhitelistRequest struct {
	MId        int    `json:"mId" binding:"required"`
	ApiIp      string `json:"apiIp" binding:"required"`
	LoginApiIp string `json:"loginApiIp" binding:"required"`
	ApiDomain  string `json:"apiDomain" binding:"required"`
}

// SaveAgentWhitelist 保存代理白名单
type SaveAgentWhitelist struct {
	MId        int    `json:"mId" binding:"required"`
	ApiIp      string `json:"apiIp" binding:"required"`
	LoginApiIp string `json:"loginApiIp" binding:"required"`
	ApiDomain  string `json:"apiDomain" binding:"required"`
}
