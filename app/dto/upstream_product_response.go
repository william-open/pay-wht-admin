package dto

// 供应商通道列表
type PayUpstreamProductListResponse struct {
	Id           int     `json:"id"`
	UpstreamId   int64   `json:"upstreamId"` // 上游供应商ID
	Title        string  `json:"title"`
	SysChannelId string  `json:"sysChannelId"`
	UpstreamCode string  `json:"upstreamCode"`
	Currency     string  `json:"currency"`
	Status       int8    `json:"status"`
	DefaultRate  float64 `json:"defaultRate"`
	AddRate      float64 `json:"addRate"`
	Weight       int     `json:"weight"`
	SuccessRate  float64 `json:"successRate"`
	OrderRange   string  `json:"orderRange"`
	Remark       string  `json:"remark"` // 备注
	Coding       string  `json:"coding"` // 系统通道编码
	Type         string  `json:"type"`   // 通道类型
}

// 供应商通道详情
type PayUpstreamProductDetailResponse struct {
	Id           int     `json:"id"`
	UpstreamId   int64   `json:"upstreamId"` // 上游供应商ID
	Title        string  `json:"title"`
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

// TestPayUpstreamProductDetailResponse 测试上游供应商通道详情
type TestPayUpstreamProductDetailResponse struct {
	MerchantName      string `json:"merchantName"`      // 商户名称
	MerchantBalance   string `json:"merchantBalance"`   // 商户余额
	UpstreamName      string `json:"upstreamName"`      // 上游供应商名称
	UpstreamBalance   string `json:"upstreamBalance"`   // 上游供应商余额
	ChannelTitle      string `json:"channelTitle"`      // 系统通道名称
	ChannelCode       string `json:"channelCode"`       // 系统通道编码
	ProductTitle      string `json:"productTitle"`      // 产品名称
	ProductCode       string `json:"productCode"`       //产品编码
	ProductOrderRange string `json:"productOrderRange"` // 产品金额范围
	Currency          string `json:"currency"`          //货币符号
}

// 根据查询条件查询供应商通道信息
type UpstreamProductResponse struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	SysChannelId int64  `json:"sysChannelId"`
	UpstreamCode string `json:"upstreamCode"`
}

// UpstreamProductTypeResponse 上游供应商支付通道类型
type UpstreamProductTypeResponse struct {
	Type   int8   `json:"type"`   // 上游支付通道类型
	Coding string `json:"coding"` // 系统通道编码
}
