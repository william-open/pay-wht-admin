package dto

import (
	"github.com/shopspring/decimal"
	"wht-admin/framework/datetime"
)

// MoneyLogResponse 资金日志列表
type MoneyLogResponse struct {
	ID            int             `json:"id"`
	Money         decimal.Decimal `json:"money"`
	OrderNo       string          `json:"orderNo"`
	Type          int             `json:"type"`
	Operator      string          `json:"operator"`
	Country       string          `json:"country"`
	MerchantTitle string          `json:"merchantTitle"`
	Currency      string          `json:"currency"`
	Balance       decimal.Decimal `json:"balance"`
	OldBalance    decimal.Decimal `json:"oldBalance"`
	CreateTime    string          `json:"createTime"`
	Description   string          `json:"description"`
}

// MoneyLogDetailResponse 资金日志详情
type MoneyLogDetailResponse struct {
	ID            int               `json:"id"`
	Money         decimal.Decimal   `json:"money"`
	OrderNo       string            `json:"orderNo"`
	Type          int               `json:"type"`
	Operator      string            `json:"operator"`
	Country       string            `json:"country"`
	MerchantTitle string            `json:"merchantTitle"`
	Currency      string            `json:"currency"`
	Balance       decimal.Decimal   `json:"balance"`
	OldBalance    decimal.Decimal   `json:"oldBalance"`
	CreateTime    datetime.Datetime `json:"createTime"`
}
