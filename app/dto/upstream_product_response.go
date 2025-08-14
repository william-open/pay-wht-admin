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

// 根据查询条件查询供应商通道信息
type UpstreamProductResponse struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	SysChannelId int64  `json:"sysChannelId"`
	UpstreamCode string `json:"upstreamCode"`
}
