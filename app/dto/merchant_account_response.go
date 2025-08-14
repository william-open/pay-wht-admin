package dto

import (
	"ruoyi-go/framework/datetime"
)

// 商户账户列表
type MerchantAccountListResponse struct {
	ID               uint64            `json:"id"`
	UID              uint64            `json:"mId"`              // 商户ID
	TotalMoney       float64           `json:"totalMoney"`       // 可用余额
	TotalFreezeMoney float64           `json:"totalFreezeMoney"` // 冻结余额
	Currency         string            `json:"currency"`         // 货币符号
	MerchantTitle    string            `json:"merchantTitle"`    // 商户名称
	Country          string            `json:"country"`          // 国家货币
	CreateTime       datetime.Datetime `json:"createTime"`
	UpdateTime       datetime.Datetime `json:"updateTime"`
}

// MerchantAccountDetailResponse 商户账户详情
type MerchantAccountDetailResponse struct {
	ID          int               `json:"id"`
	UID         uint64            `json:"mId"`         // 商户ID
	Money       float64           `json:"money"`       // 可用余额
	FreezeMoney float64           `json:"freezeMoney"` // 冻结余额
	Currency    string            `json:"currency"`    // 货币符号
	CreateTime  datetime.Datetime `json:"createTime"`
	UpdateTime  datetime.Datetime `json:"updateTime"`
}

// MerchantAccountCurrencyListResponse 商户账户货币列表
type MerchantAccountCurrencyListResponse struct {
	ID          uint64  `json:"id"`
	UID         uint64  `json:"mId"`         // 商户ID
	Money       float64 `json:"money"`       // 可用余额
	FreezeMoney float64 `json:"freezeMoney"` // 冻结余额
	Currency    string  `json:"currency"`    // 货币符号
	Country     string  `json:"country"`     // 国家货币
	CreateTime  string  `json:"createTime"`
	UpdateTime  string  `json:"updateTime"`
}
