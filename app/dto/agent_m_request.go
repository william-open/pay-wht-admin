package dto

// SaveAgentM 保存代理商户
type SaveAgentM struct {
	Id           int64   `json:"id"`
	MId          int64   `json:"mId" binding:"required"`
	AId          int64   `json:"aId" binding:"required"`
	Currency     string  `json:"currency"`
	SysChannelId int64   `json:"sysChannelId" binding:"required"`
	DefaultRate  float64 `json:"defaultRate"`
	SingleFee    float64 `json:"singleFee"`
	Status       int8    `json:"status"`
	CreateBy     string  `json:"createBy"`
	UpdateBy     string  `json:"updateBy"`
	Remark       string  `json:"remark"`
}

// AgentMListRequest 代理商户列表
type AgentMListRequest struct {
	PageRequest
	AId          int64  `query:"aId" form:"aId"`
	MerchantName string `query:"merchantName" form:"merchantName"`
	ChannelName  string `query:"channelName" form:"channelName"`
}

// CreateAgentMRequest 新增代理商户
type CreateAgentMRequest struct {
	MId          int64   `json:"mId" binding:"required"`
	AId          int64   `json:"aId" binding:"required"`
	Currency     string  `json:"currency"`
	SysChannelId int64   `json:"sysChannelId" binding:"required"`
	DefaultRate  float64 `json:"defaultRate"`
	SingleFee    float64 `json:"singleFee"`
	Remark       string  `json:"remark"`
}

// UpdateAgentMRequest 更新代理商户
type UpdateAgentMRequest struct {
	Id           int64   `json:"Id" binding:"required"`
	MId          int64   `json:"mId" binding:"required"`
	AId          int64   `json:"aId" binding:"required"`
	Currency     string  `json:"currency"`
	SysChannelId int64   `json:"sysChannelId" binding:"required"`
	DefaultRate  float64 `json:"defaultRate"`
	SingleFee    float64 `json:"singleFee"`
	Status       int8    `json:"status"`
	Remark       string  `json:"remark"`
}
