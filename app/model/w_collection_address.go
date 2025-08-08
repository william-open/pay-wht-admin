package model

import (
	"ruoyi-go/framework/datetime"
)

type WCollectionAddress struct {
	Id           int `gorm:"primaryKey;autoIncrement"`
	MId          int
	ChainSymbol  string
	Symbol       string
	Protocol     string
	CurrencyType string
	Address      string
	Status       string `gorm:"default:0"`
	CreateBy     string
	CreateTime   datetime.Datetime `gorm:"autoCreateTime"`
	UpdateBy     string
	UpdateTime   datetime.Datetime `gorm:"autoUpdateTime"`
	Remark       string
}

func (WCollectionAddress) TableName() string {
	return "w_collection_address"
}
