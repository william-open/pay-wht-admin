package model

import (
	"time"
)

type MerchantMoneyStatus int8

const (
	StatusFrozen MerchantMoneyStatus = 0 // 全冻结
	StatusActive MerchantMoneyStatus = 1 // 可用
)

type WMerchantMoney struct {
	ID          uint                `gorm:"primaryKey;autoIncrement;comment:主键"`
	UID         uint                `gorm:"column:uid;not null;index:idx_uid;comment:用户ID"`
	Currency    string              `gorm:"column:currency;type:varchar(10);not null;index:idx_currency;comment:货币"`
	Money       int                 `gorm:"column:money;not null;default:0;comment:可用余额"`
	FreezeMoney int                 `gorm:"column:freeze_money;not null;default:0;comment:冻结余额"`
	CreateTime  time.Time           `gorm:"column:createtime;type:timestamp;not null;comment:创建时间"`
	UpdateTime  time.Time           `gorm:"column:updatetime;type:timestamp;not null;comment:更新时间"`
	Status      MerchantMoneyStatus `gorm:"column:status;not null;default:1;comment:状态:0=全冻结,1=可用"`
}

// 设置复合唯一索引
func (WMerchantMoney) TableIndex() [][]string {
	return [][]string{
		{"uid", "currency"},
	}
}

func (WMerchantMoney) TableName() string {
	return "w_merchant_money"
}
