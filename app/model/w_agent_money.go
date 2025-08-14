package model

import (
	"time"
)

type WAgentMoney struct {
	ID         uint      `gorm:"primaryKey;autoIncrement;comment:主键"`
	UID        uint      `gorm:"column:uid;not null;default:0;index:idx_uid;comment:代理用户ID"`
	Money      float64   `gorm:"column:money;not null;default:0;comment:收益金额"`
	OrderMoney float64   `gorm:"column:order_money;not null;default:0;comment:来源订单金额"`
	OrderNo    string    `gorm:"column:order_no;type:varchar(50);not null;comment:订单编码"`
	PID        uint      `gorm:"column:pid;not null;default:0;index:idx_pid;comment:下线用户ID"`
	Currency   string    `gorm:"column:currency;type:varchar(6);not null;index:idx_currency;comment:币种"`
	Type       int8      `gorm:"column:type;not null;default:11;index:idx_type;comment:收益类型:11=代收,21=代付,31=充值,41=提现"`
	Remark     string    `gorm:"column:remark;type:varchar(10);comment:备注"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;default:0;comment:创建时间戳"`
}

func (WAgentMoney) TableName() string {
	return "w_agent_money"
}
