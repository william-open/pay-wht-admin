package model

import (
	"github.com/shopspring/decimal"
	"ruoyi-go/framework/datetime"

	"gorm.io/gorm"
)

type WMerchant struct {
	MId               int `gorm:"primaryKey;autoIncrement"`
	Username          string
	Password          string
	Nickname          string
	CallbackSecretKey string
	NotifyUrl         string
	AesSecretKey      string
	PublicKey         string
	PrivateKey        string
	AppId             string
	ApiKey            string
	ApiIp             string
	LoginApiIp        string
	ApiDomain         string
	Balance           decimal.Decimal
	Status            string `gorm:"default:0"`
	PayType           string `gorm:"default:1"`
	userType          string `gorm:"default:1"`
	CreateBy          string
	CreateTime        datetime.Datetime `gorm:"autoCreateTime"`
	UpdateBy          string
	UpdateTime        datetime.Datetime `gorm:"autoUpdateTime"`
	DeleteTime        gorm.DeletedAt
	Remark            string
	UpstreamId        string
	Ways              string
	UserType          string
	PId               int64
}

func (WMerchant) TableName() string {
	return "w_merchant"
}
