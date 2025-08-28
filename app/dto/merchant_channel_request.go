package dto

import "time"

// SaveMerchantChannel 保存商户通道
type SaveMerchantChannel struct {
	ID           int       `json:"id"`
	MId          int64     `json:"mId"`          // 商户ID
	Currency     string    `json:"currency"`     // 货币符号
	SysChannelID int64     `json:"sysChannelId"` // 系统通道编码ID
	UpChannelID  int64     `json:"upChannelId"`  // 上游通道编码ID
	Status       string    `json:"status"`       // 1:开启 0:关闭
	DefaultRate  float64   `json:"defaultRate"`  // 默认费率
	SingleFee    float64   `json:"singleFee"`    // 单笔费用
	Weight       int       `json:"weight"`       // 权重值
	SuccessRate  float64   `json:"successRate"`  // 成功率
	OrderRange   string    `json:"orderRange"`   // 订单金额范围
	Remark       string    `json:"remark"`       // 备注
	CreateBy     string    `json:"createBy"`     // 创建者
	CreateTime   time.Time `json:"createTime"`   // 创建时间
	UpdateBy     string    `json:"updateBy"`     // 更新者
	UpdateTime   time.Time `json:"updateTime"`   // 更新时间
}

// MerchantChannelListRequest 商户通道列表
type MerchantChannelListRequest struct {
	PageRequest
	Keyword  string `query:"keyword" form:"keyword"`
	Currency string `query:"currency" form:"currency"`
	Status   *int8  `query:"status" form:"status"`
	MId      int8   `query:"mId" form:"mId"`
}

// CreateMerchantChannelRequest 新增商户通道
type CreateMerchantChannelRequest struct {
	Currency         string  `json:"currency"`         // 货币符号
	SysChannelID     int64   `json:"sysChannelId"`     // 系统通道编码ID
	MId              int64   `json:"mId"`              // 商户ID
	UpChannelID      int64   `json:"upChannelId"`      // 上游通道编码ID
	Status           string  `json:"status"`           // 1:开启 0:关闭
	DefaultRate      float64 `json:"defaultRate"`      // 默认费率
	SingleFee        float64 `json:"singleFee"`        // 单笔费用
	Weight           int     `json:"weight"`           // 权重值
	SuccessRate      float64 `json:"successRate"`      // 成功率
	OrderRange       string  `json:"orderRange"`       // 订单金额范围
	Remark           string  `json:"remark"`           // 备注
	UpstreamProducts []int64 `json:"upstreamProducts"` // 绑定上游通道产品
}

// UpdateMerchantChannelRequest 更新商户通道
type UpdateMerchantChannelRequest struct {
	ID               int     `json:"id"`
	MId              int64   `json:"mId"`              // 商户ID
	Currency         string  `json:"currency"`         // 货币符号
	SysChannelID     int64   `json:"sysChannelId"`     // 系统通道编码ID
	UpChannelID      int64   `json:"upChannelId"`      // 上游通道编码ID
	Status           string  `json:"status"`           // 1:开启 0:关闭
	DefaultRate      float64 `json:"defaultRate"`      // 默认费率
	SingleFee        float64 `json:"singleFee"`        // 单笔费用
	Weight           int     `json:"weight"`           // 权重值
	SuccessRate      float64 `json:"successRate"`      // 成功率
	OrderRange       string  `json:"orderRange"`       // 订单金额范围
	UpstreamProducts []int64 `json:"upstreamProducts"` // 绑定上游通道产品
}

// UpdateMerchantChannelStatusRequest 修改商户通道状态
type UpdateMerchantChannelStatusRequest struct {
	ID     int    `json:"id" binding:"required"`
	Status string `json:"status" binding:"required"`
}
