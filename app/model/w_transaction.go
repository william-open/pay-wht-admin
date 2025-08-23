package model

import (
	"github.com/shopspring/decimal"
	"wht-admin/framework/datetime"
)

type WTransaction struct {
	Id           int `gorm:"primaryKey;autoIncrement"`
	MId          int
	CurrencyType string
	BizNo        string
	BlockNum     string
	TxId         string
	fromAddress  string
	toAddress    string
	Amount       decimal.Decimal
	Status       string `gorm:"default:0"`
	CreateBy     string
	CreateTime   datetime.Datetime `gorm:"autoCreateTime"`
	UpdateBy     string
	UpdateTime   datetime.Datetime `gorm:"autoUpdateTime"`
	Remark       string
}

func (WTransaction) TableName() string {
	return "w_transaction"
}
