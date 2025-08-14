package model

import (
	"time"
)

type WAgentMerchant struct {
	ID           int64     `gorm:"primaryKey;autoIncrement"` // 主键自增
	AID          int64     `gorm:"column:a_id;not null"`     // 代理ID
	MID          int64     `gorm:"column:m_id;not null"`     // 商户ID
	Currency     string    `gorm:"column:currency;type:varchar(30)"`
	SysChannelID int64     `gorm:"column:sys_channel_id;not null"`                     // 系统通道编码ID
	UpChannelID  int64     `gorm:"column:up_channel_id;not null"`                      // 上游通道编码ID
	Status       int8      `gorm:"column:status;default:0"`                            // 状态：0关闭，1开启
	DefaultRate  float64   `gorm:"column:default_rate;type:decimal(4,2);default:0.00"` // 代理抽点费率
	SingleFee    float64   `gorm:"column:single_fee;type:decimal(4,2);default:0.00"`   // 单笔费用
	Remark       string    `gorm:"column:remark;type:varchar(100)"`
	CreateBy     string    `gorm:"column:create_by;type:varchar(64);default:''"`
	CreateTime   time.Time `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP"`
	UpdateBy     string    `gorm:"column:update_by;type:varchar(64);default:''"`
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

// 自定义表名（可选，默认是结构体名的蛇形复数）
func (WAgentMerchant) TableName() string {
	return "w_agent_merchant"
}
