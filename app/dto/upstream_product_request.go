package dto

// 保存供应商通道
type SavePayUpstreamProduct struct {
	Id           int     `json:"id"`
	UpstreamId   int64   `json:"upstreamId"` // 上游供应商ID
	Title        string  `json:"title"`
	SysChannelId int64   `json:"sysChannelId"`
	Currency     string  `json:"currency"`
	UpstreamCode string  `json:"upstreamCode"`
	Status       int8    `json:"status"`
	DefaultRate  float64 `json:"defaultRate"`
	AddRate      float64 `json:"addRate"`
	Weight       int     `json:"weight"`
	SuccessRate  float64 `json:"successRate"`
	OrderRange   string  `json:"orderRange"`
	Remark       string  `json:"remark"` // 备注
}

// 供应商通道列表
type PayUpstreamProductListRequest struct {
	PageRequest
	UpstreamId   int64  `query:"upstreamId" form:"upstreamId"`
	Title        string `query:"title" form:"title"`
	ChannelCode  string `query:"channelCode" form:"channelCode"`
	UpstreamCode string `query:"upstreamCode" form:"upstreamCode"`
	Currency     string `query:"currency" form:"currency"`
	Status       string `query:"status" form:"status"`
	Keyword      string `query:"roleName" form:"roleName"`
}

// 新增供应商通道
type CreatePayUpstreamProductRequest struct {
	UpstreamId   int64   `json:"upstreamId"` // 上游供应商ID
	Title        string  `json:"title"`
	Currency     string  `json:"currency"`
	SysChannelId int64   `json:"sysChannelId"`
	UpstreamCode string  `json:"upstreamCode"`
	Status       int8    `json:"status"`
	DefaultRate  float64 `json:"defaultRate"`
	AddRate      float64 `json:"addRate"`
	Weight       int     `json:"weight"`
	SuccessRate  float64 `json:"successRate"`
	OrderRange   string  `json:"orderRange"`
	Remark       string  `json:"remark"` // 备注
}

// 更新供应商通道
type UpdatePayUpstreamProductRequest struct {
	Id           int     `json:"id"`
	UpstreamId   int64   `json:"UpstreamId"` // 上游供应商ID
	Title        string  `json:"title"`
	Currency     string  `json:"currency"`
	SysChannelId int64   `json:"sysChannelId"`
	UpstreamCode string  `json:"upstreamCode"`
	Status       int8    `json:"status"`
	DefaultRate  float64 `json:"defaultRate"`
	AddRate      float64 `json:"addRate"`
	Weight       int     `json:"weight"`
	SuccessRate  float64 `json:"successRate"`
	OrderRange   string  `json:"orderRange"`
	Remark       string  `json:"remark"` // 备注
}

// 供应商通道
type UpstreamProductRequest struct {
	UpstreamId   int64  `query:"upstreamId" form:"upstreamId"` // 上游供应商ID
	SysChannelId string `query:"sysChannelId" form:"sysChannelId"`
	Status       int8   `query:"status" form:"status"`
	Currency     string `query:"currency" form:"currency"` // 货币符号
	Keyword      string `query:"keyword" form:"keyword"`   // 关键字
}

// 供应商通道产品是否存在
type ExistPayUpstreamProductRequest struct {
	UpstreamId   int64  `json:"UpstreamId"` // 上游供应商ID
	Title        string `json:"title"`
	Currency     string `json:"currency"`
	SysChannelId int64  `json:"sysChannelId"`
	UpstreamCode string `json:"upstreamCode"`
}

// TestCreatePayUpstreamProductRequest 测试上游供应商通道产品
type TestCreatePayUpstreamProductRequest struct {
	Id           int64  `json:"id"`           // 上游供应商通道产品ID
	Amount       string `json:"amount"`       // 订单金额
	AccName      string `json:"accName"`      // 姓名
	PayType      string `json:"payType"`      // 通道编码
	AccNo        string `json:"accNo"`        // 账号
	PayEmail     string `json:"payEmail"`     // 邮箱
	BankName     string `json:"bankName"`     // 银行名称
	BankCode     string `json:"bankCode"`     // 银行编码
	PayPhone     string `json:"payPhone"`     //手机号码
	PayMethod    string `json:"payMethod"`    // 支付方式
	IdentityType string `json:"identityType"` // 证件类型
	IdentityNum  string `json:"identityNum"`  // 证件号码
}
