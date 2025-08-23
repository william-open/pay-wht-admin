package dto

import (
	"wht-admin/framework/datetime"
)

// CollectionListResponse  归集列表
type CollectionListResponse struct {
	Id           int               `json:"id"`
	MId          int               `json:"mId"`
	CurrencyType string            `json:"currencyType"`
	BizNo        string            `json:"bizNo"`
	BlockNum     string            `json:"blockNum"`
	TxId         string            `json:"txId"`
	FromAddress  string            `json:"fromAddress"`
	ToAddress    string            `json:"toAddress"`
	Amount       string            `json:"amount"`
	Status       string            `json:"status"`
	CreateBy     string            `json:"createBy"`
	CreateTime   datetime.Datetime `json:"createTime"`
}

// CollectionDetailResponse 归集详情
type CollectionDetailResponse struct {
	Id           int               `json:"id"`
	MId          int               `json:"mId"`
	CurrencyType string            `json:"currencyType"`
	BizNo        string            `json:"bizNo"`
	BlockNum     string            `json:"blockNum"`
	TxId         string            `json:"txId"`
	FromAddress  string            `json:"fromAddress"`
	ToAddress    string            `json:"toAddress"`
	Status       string            `json:"status"`
	CreateBy     string            `json:"createBy"`
	CreateTime   datetime.Datetime `json:"createTime"`
	UpdateBy     string            `json:"updateBy"`
	UpdateTime   datetime.Datetime `json:"updateTime"`
	Remark       string            `json:"remark"`
}
