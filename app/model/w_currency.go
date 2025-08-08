package model

import (
	"ruoyi-go/framework/datetime"

	"gorm.io/gorm"
)

type WCurrency struct {
	CurrencyId      int `gorm:"primaryKey;autoIncrement"`
	PId             int
	Currency        string
	Symbol          string
	Logo            string
	ContractAddress string
	CurrencyType    string
	ChainSymbol     string
	Protocol        string
	Decimals        int
	Status          string `gorm:"default:0"`
	CreateBy        string
	CreateTime      datetime.Datetime `gorm:"autoCreateTime"`
	UpdateBy        string
	UpdateTime      datetime.Datetime `gorm:"autoUpdateTime"`
	DeleteTime      gorm.DeletedAt
	Remark          string
}

func (WCurrency) TableName() string {
	return "w_currency"
}
