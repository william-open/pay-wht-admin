package model

import (
	"time"
)

type MoneyLogType int8

const (
	TypeCollection       MoneyLogType = 1  // 代收
	TypePayment          MoneyLogType = 2  // 代付
	TypeRecharge         MoneyLogType = 3  // 充值
	TypeWithdrawal       MoneyLogType = 4  // 提现
	TypeFee              MoneyLogType = 5  // 手续费
	TypeCollectionProfit MoneyLogType = 11 // 代收收益
	TypePaymentProfit    MoneyLogType = 21 // 代付收益
	TypeRechargeProfit   MoneyLogType = 31 // 充值收益
	TypeWithdrawalProfit MoneyLogType = 41 // 提现收益
	TypeUnfreeze         MoneyLogType = 61 // 解冻资金
	TypeFreezeDelete     MoneyLogType = 62 // 删除冻结
)

type WMoneyLog struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	UID         int       `gorm:"column:uid;not null;comment:用户id"`
	Money       float64   `gorm:"column:money;type:decimal(13,2);not null;comment:金额"`
	OrderNo     string    `gorm:"column:order_no;size:30;comment:订单id"`
	Type        int8      `gorm:"column:type;not null;comment:1:代收,2代付..."`
	Operator    string    `gorm:"column:operator;size:15;not null;comment:操作者"`
	Currency    string    `gorm:"column:currency;size:10;not null;comment:币种"`
	Description string    `gorm:"column:description;size:30;comment:备注"`
	OldBalance  float64   `gorm:"column:old_balance;type:decimal(13,2);not null;comment:旧余额"`
	Balance     float64   `gorm:"column:balance;type:decimal(13,2);not null;comment:余额"`
	CreateTime  time.Time `gorm:"column:create_time;type:timestamp;not null;comment:时间戳"`
}

func (WMoneyLog) TableName() string {
	return "w_money_log"
}
