package model

import (
	"github.com/shopspring/decimal"
	"ruoyi-go/framework/datetime"
)

type WAddress struct {
	AddressId       int `gorm:"primaryKey;autoIncrement"`
	MId             int
	ChainSymbol     string
	address         string
	privateKey      string
	TrxAmount       decimal.Decimal
	Trc20UsdtAmount decimal.Decimal
	Status          string `gorm:"default:0"`
	CreateBy        string
	CreateTime      datetime.Datetime `gorm:"autoCreateTime"`
	UpdateBy        string
	UpdateTime      datetime.Datetime `gorm:"autoUpdateTime"`
	Remark          string
}

func (WAddress) TableName() string {
	return "w_address"
}
