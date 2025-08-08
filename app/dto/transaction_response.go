package dto

import (
	"ruoyi-go/framework/datetime"
)

// TransactionListResponse  交易列表
type TransactionListResponse struct {
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

// TransactionDetailResponse 交易详情
type TransactionDetailResponse struct {
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
