package dto

import (
	"wht-admin/framework/datetime"
)

// MerchantWhitelistResponse 商户白名单列表
type MerchantWhitelistResponse struct {
	ID         int               `json:"id"`
	MID        uint64            `json:"mId"`
	IPAddress  string            `json:"iPAddress"`
	CanAdmin   uint8             `json:"canAdmin"`
	CanPayout  uint8             `json:"canPayout"`
	CanReceive uint8             `json:"canReceive"`
	CreateTime datetime.Datetime `json:"createTime"`
	Remark     string            `json:"remark"`
}
