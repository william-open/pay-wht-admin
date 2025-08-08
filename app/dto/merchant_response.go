package dto

import (
	"github.com/shopspring/decimal"
	"ruoyi-go/framework/datetime"
)

// 商户列表
type MerchantListResponse struct {
	MId        int               `json:"mId"`
	Username   string            `json:"username"`
	ApiKey     string            `json:"apiKey"`
	Nickname   string            `json:"nickname"`
	AppId      string            `json:"appId"`
	Balance    decimal.Decimal   `json:"balance"`
	Status     string            `json:"status"`
	CreateBy   string            `json:"createBy"`
	CreateTime datetime.Datetime `json:"createTime"`
}

type MerchantMetaResponse struct {
	Title   string `json:"title"`
	Icon    string `json:"icon"`
	Link    string `json:"link"`
	NoCache bool   `json:"noCache"`
}

// MerchantDetailResponse 商户详情
type MerchantDetailResponse struct {
	MId               int               `json:"mId"`
	Username          string            `json:"username"`
	Nickname          string            `json:"nickname"`
	CallbackSecretKey string            `json:"callbackSecretKey"`
	NotifyUrl         string            `json:"notifyUrl"`
	AesSecretKey      string            `json:"aesSecretKey"`
	PublicKey         string            `json:"publicKey"`
	PrivateKey        string            `json:"privateKey"`
	AppId             string            `json:"appId"`
	Password          string            `json:"password"`
	Balance           decimal.Decimal   `json:"balance"`
	Status            string            `json:"status"`
	CreateBy          string            `json:"createBy"`
	CreateTime        datetime.Datetime `json:"createTime"`
	UpdateBy          string            `json:"updateBy"`
	UpdateTime        datetime.Datetime `json:"updateTime"`
	Remark            string            `json:"remark"`
}
