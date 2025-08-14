package model

import "time"

type WMerchantChannel struct {
	ID           int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Currency     string    `gorm:"column:currency;size:30" json:"currency"`                                 // 货币符号
	MId          int64     `gorm:"column:m_id;not null" json:"mId"`                                         // 商户ID
	SysChannelID int64     `gorm:"column:sys_channel_id;not null" json:"sysChannelId"`                      // 系统通道编码ID
	UpChannelID  int64     `gorm:"column:up_channel_id;not null" json:"upChannelId"`                        // 上游通道编码ID
	Status       string    `gorm:"column:status;default:0" json:"status"`                                   // 1:开启 0:关闭
	DefaultRate  float64   `gorm:"column:default_rate;type:decimal(4,2);default:0.00" json:"defaultRate"`   // 默认费率
	SingleFee    float64   `gorm:"column:single_fee;type:decimal(4,2);default:0.00" json:"singleFee"`       // 单笔费用
	Weight       int       `gorm:"column:weight;default:1" json:"weight"`                                   // 权重值
	SuccessRate  float64   `gorm:"column:success_rate;type:decimal(5,2);default:100.00" json:"successRate"` // 成功率
	OrderRange   string    `gorm:"column:order_range;size:100;not null" json:"orderRange"`                  // 订单金额范围
	Remark       string    `gorm:"column:remark;size:100" json:"remark"`                                    // 备注
	CreateBy     string    `gorm:"column:create_by;size:64;default:''" json:"createBy"`                     // 创建者
	CreateTime   time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`                     // 创建时间
	UpdateBy     string    `gorm:"column:update_by;size:64;default:''" json:"updateBy"`                     // 更新者
	UpdateTime   time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`                     // 更新时间
}

// TableName 指定表名
func (WMerchantChannel) TableName() string {
	return "w_merchant_channel"
}
