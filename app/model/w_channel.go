package model

import (
	"ruoyi-go/framework/datetime"
)

type WChannel struct {
	Id          int     `gorm:"primaryKey;autoIncrement"`
	Title       string  `json:"title"`
	Status      int     `json:"status"`
	Max         string  `json:"max" gorm:"default:99999"`
	Min         string  `json:"min" gorm:"default:0"`
	DefaultRate float64 `json:"defaultRate"`
	Coding      string  `json:"coding"`
	AddRate     float64 `json:"addRate"`
	Type        int     `json:"type"`
	Charge      int     `json:"charge"`
	Currency    string  `json:"currency"`
	Remark      string  `json:"remark"`
	CreateBy    string
	CreateTime  datetime.Datetime `gorm:"autoCreateTime"`
	UpdateBy    string
	UpdateTime  datetime.Datetime `gorm:"autoUpdateTime"`
}

func (WChannel) TableName() string {
	return "w_pay_way"
}
