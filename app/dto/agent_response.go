package dto

import (
	"github.com/shopspring/decimal"
	"ruoyi-go/framework/datetime"
)

// AgentListResponse 代理列表
type AgentListResponse struct {
	MId        int               `json:"mId"`
	Username   string            `json:"username"`
	ApiKey     string            `json:"apiKey"`
	Nickname   string            `json:"nickname"`
	AppId      string            `json:"appId"`
	UserType   string            `json:"userType"`
	Balance    decimal.Decimal   `json:"balance"`
	Status     string            `json:"status"`
	SubCount   int               `json:"subCount"`
	CreateBy   string            `json:"createBy"`
	CreateTime datetime.Datetime `json:"createTime"`
}

// AgentDetailResponse 代理详情
type AgentDetailResponse struct {
	MId               int               `json:"mId"`
	Username          string            `json:"username"`
	Nickname          string            `json:"nickname"`
	CallbackSecretKey string            `json:"callbackSecretKey"`
	NotifyUrl         string            `json:"notifyUrl"`
	AesSecretKey      string            `json:"aesSecretKey"`
	PublicKey         string            `json:"publicKey"`
	PrivateKey        string            `json:"privateKey"`
	AppId             string            `json:"appId"`
	ApiIp             string            `json:"apiIp"`
	LoginApiIp        string            `json:"loginApiIp"`
	ApiDomain         string            `json:"apiDomain"`
	PayType           string            `json:"payType"`
	Password          string            `json:"password"`
	UserType          string            `json:"userType"`
	Balance           decimal.Decimal   `json:"balance"`
	Status            string            `json:"status"`
	CreateBy          string            `json:"createBy"`
	CreateTime        datetime.Datetime `json:"createTime"`
	UpdateBy          string            `json:"updateBy"`
	UpdateTime        datetime.Datetime `json:"updateTime"`
	Remark            string            `json:"remark"`
}
