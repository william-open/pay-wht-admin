package dto

import (
	"wht-admin/framework/datetime"
)

// CollectionAddressListResponse 归集钱包地址列表
type CollectionAddressListResponse struct {
	Id           int               `json:"id"`
	MId          int               `json:"mId"`
	Currency     string            `json:"currency"`
	Symbol       string            `json:"symbol"`
	CurrencyType string            `json:"currency_type"`
	Protocol     string            `json:"protocol"`
	Address      string            `json:"address"`
	ChainSymbol  string            `json:"chain_symbol"`
	Status       string            `json:"status"`
	CreateBy     string            `json:"createBy"`
	CreateTime   datetime.Datetime `json:"createTime"`
}

// CollectionAddressDetailResponse 归集钱包地址详情
type CollectionAddressDetailResponse struct {
	Id           int               `json:"id"`
	MId          int               `json:"mId"`
	Currency     string            `json:"currency"`
	Symbol       string            `json:"symbol"`
	CurrencyType string            `json:"currency_type"`
	Protocol     string            `json:"protocol"`
	Address      string            `json:"address"`
	ChainSymbol  string            `json:"chain_symbol"`
	Status       string            `json:"status"`
	CreateBy     string            `json:"createBy"`
	CreateTime   datetime.Datetime `json:"createTime"`
	UpdateBy     string            `json:"updateBy"`
	UpdateTime   datetime.Datetime `json:"updateTime"`
	Remark       string            `json:"remark"`
}
