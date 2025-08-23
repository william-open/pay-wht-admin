package dto

import (
	"wht-admin/framework/datetime"
)

// AddressListResponse 钱包地址列表
type AddressListResponse struct {
	AddressId   int               `json:"addressId"`
	MId         int               `json:"mId"`
	ChainSymbol string            `json:"chainSymbol"`
	Address     string            `json:"address"`
	Status      string            `json:"status"`
	CreateBy    string            `json:"createBy"`
	CreateTime  datetime.Datetime `json:"createTime"`
}

// AddressDetailResponse 钱包地址详情
type AddressDetailResponse struct {
	AddressId   int               `json:"addressId"`
	MId         int               `json:"mId"`
	ChainSymbol string            `json:"currency"`
	Address     string            `json:"address"`
	Status      string            `json:"status"`
	CreateBy    string            `json:"createBy"`
	CreateTime  datetime.Datetime `json:"createTime"`
	UpdateBy    string            `json:"updateBy"`
	UpdateTime  datetime.Datetime `json:"updateTime"`
	Remark      string            `json:"remark"`
}
