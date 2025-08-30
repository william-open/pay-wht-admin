package dto

import (
	"wht-admin/framework/datetime"
)

// 商户通道列表
type MerchantChannelListResponse struct {
	ID             int               `json:"id"`
	MId            int64             `json:"mId"`            // 商户ID
	Currency       string            `json:"currency"`       // 货币符号
	UpChannelID    int64             `json:"upChannelId"`    // 上游通道编码ID
	Status         string            `json:"status"`         // 1:开启 0:关闭
	DefaultRate    float64           `json:"defaultRate"`    // 默认费率
	SingleFee      float64           `json:"singleFee"`      // 单笔费用
	Type           int               `json:"type"`           // 通道类型
	UpstreamTitle  string            `json:"upstreamTitle"`  // 上游供应商名称
	UpProductTitle string            `json:"upProductTitle"` // 上游产品名称
	UpstreamCode   string            `json:"upstreamCode"`   // 上游产品名称
	UpDefaultRate  float64           `json:"upDefaultRate"`  // 上游供应商通道默认费率
	UpAddRate      float64           `json:"upAddRate"`      // 上游供应商通道单笔费用
	Coding         string            `json:"coding"`         // 系统通道编码
	Country        string            `json:"country"`        // 国家
	Weight         string            `json:"weight"`         // 国家
	OrderRange     string            `json:"orderRange"`     // 商户金额范围
	UpOrderRange   string            `json:"upOrderRange"`   // 上游通道金额范围
	MerchantTitle  string            `json:"merchantTitle"`  // 商户名称
	CreateBy       string            `json:"createBy"`
	CreateTime     datetime.Datetime `json:"createTime"`
	UpdateBy       string            `json:"updateBy"`
	UpdateTime     datetime.Datetime `json:"updateTime"`
	Remark         string            `json:"remark"`
}

// MerchantChannelDetailResponse 商户通道详情
type MerchantChannelDetailResponse struct {
	ID                   int     `json:"id"`
	Country              string  `json:"country"`              // 国家
	Coding               string  `json:"coding"`               // 系统通道编码
	UpstreamCode         string  `json:"upstreamCode"`         // 上游通道编码
	DefaultRate          float64 `json:"defaultRate"`          // 默认费率
	SingleFee            float64 `json:"singleFee"`            // 单笔费用
	Weight               int     `json:"weight"`               // 权重值
	OrderRange           string  `json:"orderRange"`           // 订单金额范围
	SysTitle             string  `json:"sysTitle"`             //系统通道名称
	UpstreamTitle        string  `json:"upstreamTitle"`        //上游通道名称
	Currency             string  `json:"currency"`             //货币符号
	UpstreamChannelTitle string  `json:"upstreamChannelTitle"` //上游通道产品
}
