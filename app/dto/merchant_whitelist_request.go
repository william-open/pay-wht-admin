package dto

// SaveMerchantWhitelist 保存商户通道
type SaveMerchantWhitelist struct {
	MID        uint64 `json:"mId"`
	IPAddress  string `json:"iPAddress"`
	CanAdmin   uint8  `json:"canAdmin"`
	CanPayout  uint8  `json:"canPayout"`
	CanReceive uint8  `json:"canReceive"`
	CreateBy   string `json:"createBy"`
	Remark     string `json:"remark"`
}

// MerchantWhitelistRequest 商户白名单列表
type MerchantWhitelistRequest struct {
	PageRequest
	MId int8 `query:"mId" form:"mId"`
}

// CreateMerchantWhitelistRequest 新增商户白名单
type CreateMerchantWhitelistRequest struct {
	MID        uint64 `json:"mId"`
	IPAddress  string `json:"iPAddress"`
	CanAdmin   uint8  `json:"canAdmin"`
	CanPayout  uint8  `json:"canPayout"`
	CanReceive uint8  `json:"canReceive"`
	CreateBy   string `json:"createBy"`
	Remark     string `json:"remark"`
}

// RemoveMerchantWhitelistRequest 删除商户白名单
type RemoveMerchantWhitelistRequest struct {
	Id uint64 `json:"id"`
}
