package model

import (
	"gorm.io/gorm"
	"wht-admin/framework/datetime"
)

type WMerchantWhitelist struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement;comment:主键"`
	MID        uint64 `gorm:"column:m_id;not null;comment:商户ID"`
	IPAddress  string `gorm:"size:32;not null;comment:IP地址"`
	CanAdmin   uint8  `gorm:"not null;default:0;comment:登录后台权限"`
	CanPayout  uint8  `gorm:"not null;default:0;comment:代付下单"`
	CanReceive uint8  `gorm:"not null;default:0;comment:代收下单"`
	CreateBy   string
	CreateTime datetime.Datetime `gorm:"autoCreateTime"`
	UpdateBy   string
	UpdateTime datetime.Datetime `gorm:"autoUpdateTime"`
	DeleteTime gorm.DeletedAt
	Remark     string `gorm:"size:100;comment:备注"`
}

func (WMerchantWhitelist) TableName() string {
	return "w_merchant_whitelist"
}
