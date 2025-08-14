package model

import "time"

// WPayUpstreamProduct 上游产品与系统通道对应表
type WPayUpstreamProduct struct {
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UpstreamId   int64     `gorm:"column:upstream_id;not null" json:"upstreamId"` // 上游供应商ID
	Title        string    `gorm:"type:varchar(40);not null;comment:上游产品名称" json:"title"`
	SysChannelId int64     `gorm:"column:sys_channel_id;not null" json:"sysChannelId"` // 系统通道ID
	UpstreamCode string    `gorm:"type:varchar(100);not null;comment:上游通道编码" json:"upstreamCode"`
	Currency     string    `gorm:"type:varchar(100);not null;comment:国家货币" json:"currency"`
	Status       int8      `gorm:"type:tinyint(1);default:0;comment:1:开启0:关闭" json:"status"`
	DefaultRate  float64   `gorm:"type:decimal(4,2);default:0.00;comment:默认费率" json:"defaultRate"`
	AddRate      float64   `gorm:"type:decimal(4,2);default:0.00;comment:默认添加费率" json:"addRate"`
	Weight       int       `gorm:"default:1;comment:权重值" json:"weight"`
	SuccessRate  float64   `gorm:"type:decimal(5,2);default:100.00;comment:成功率" json:"successRate"`
	OrderRange   string    `gorm:"type:varchar(100);not null;comment:订单范围" json:"orderRange"`
	Remark       string    `gorm:"type:varchar(100);comment:备注" json:"remark"`
	CreateBy     string    `gorm:"type:varchar(64);default:'';comment:创建者" json:"createBy"`
	CreateTime   time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createTime"`
	UpdateBy     string    `gorm:"type:varchar(64);default:'';comment:更新者" json:"updateBy"`
	UpdateTime   time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP;comment:更新时间" json:"updateTime"`
}

func (WPayUpstreamProduct) TableName() string {
	return "w_pay_upstream_product"
}
