package model

import (
	"gorm.io/gorm"
	"time"
)

type WMerchantAccount struct {
	ID          uint64  `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	UID         uint64  `gorm:"column:uid;not null;index:idx_uid" json:"uid"`
	Status      int8    `gorm:"column:status;default:1" json:"status"`
	Currency    string  `gorm:"column:currency;type:varchar(10);not null;index:idx_currency" json:"currency"`
	Money       float64 `gorm:"column:money;type:decimal(18,4);default:0.0000" json:"money"`
	FreezeMoney float64 `gorm:"column:freeze_money;type:decimal(18,4);default:0.0000" json:"freeze_money"`
	CreateTime  int64   `gorm:"column:create_time;not null" json:"create_time"`
	UpdateTime  int64   `gorm:"column:update_time;not null" json:"update_time"`

	// 添加格式化后的时间字段（不映射到数据库）
	CreateTimeFormatted string `gorm:"-" json:"create_time_formatted,omitempty"`
	UpdateTimeFormatted string `gorm:"-" json:"update_time_formatted,omitempty"`
}

func (WMerchantAccount) TableName() string {
	return "w_merchant_money"
}

// AfterFind 钩子 - 查询后自动格式化时间
func (m *WMerchantAccount) AfterFind(tx *gorm.DB) (err error) {
	m.CreateTimeFormatted = time.Unix(m.CreateTime, 0).Format("2006-01-02 15:04:05")
	m.UpdateTimeFormatted = time.Unix(m.UpdateTime, 0).Format("2006-01-02 15:04:05")
	return nil
}

// BeforeCreate 钩子 - 设置创建和更新时间
func (m *WMerchantAccount) BeforeCreate(tx *gorm.DB) error {
	now := time.Now().Unix()
	m.CreateTime = now
	m.UpdateTime = now
	return nil
}

// BeforeUpdate 钩子 - 设置更新时间
func (m *WMerchantAccount) BeforeUpdate(tx *gorm.DB) error {
	m.UpdateTime = time.Now().Unix()
	return nil
}
