package model

import (
	"github.com/shopspring/decimal"
	"wht-admin/framework/datetime"
)

type WCollectionAddressAmount struct {
	Id             int `gorm:"primaryKey;autoIncrement"`
	MId            int
	CollectionId   int
	CurrencyType   string
	CurrencySymbol string
	ChainSymbol    string
	Protocol       string
	Address        string
	Amount         decimal.Decimal
	Precision      int
	CreateBy       string
	CreateTime     datetime.Datetime `gorm:"autoCreateTime"`
	UpdateBy       string
	UpdateTime     datetime.Datetime `gorm:"autoUpdateTime"`
	Remark         string
}

func (WCollectionAddressAmount) TableName() string {
	return "w_collection_address_amount"
}
