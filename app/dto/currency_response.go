package dto

import (
	"ruoyi-go/framework/datetime"
)

// CurrencyListResponse 币种列表
type CurrencyListResponse struct {
	CurrencyId      int               `json:"currencyId"`
	PId             int               `json:"pId"`
	Currency        string            `json:"currency"`
	Symbol          string            `json:"symbol"`
	Logo            string            `json:"logo"`
	ContractAddress string            `json:"contractAddress"`
	CurrencyType    string            `json:"currencyType"`
	Protocol        string            `json:"protocol"`
	Decimals        int               `json:"decimals"`
	Status          string            `json:"status"`
	CreateBy        string            `json:"createBy"`
	CreateTime      datetime.Datetime `json:"createTime"`
}

// CurrencyDetailResponse 币种详情
type CurrencyDetailResponse struct {
	PId             int               `json:"pId"`
	Currency        string            `json:"currency"`
	Symbol          string            `json:"symbol"`
	Logo            string            `json:"logo"`
	ContractAddress string            `json:"contractAddress"`
	CurrencyType    string            `json:"currencyType"`
	Protocol        string            `json:"protocol"`
	Precision       int               `json:"precision"`
	Status          string            `json:"status"`
	CreateBy        string            `json:"createBy"`
	CreateTime      datetime.Datetime `json:"createTime"`
	UpdateBy        string            `json:"updateBy"`
	UpdateTime      datetime.Datetime `json:"updateTime"`
	Remark          string            `json:"remark"`
}
