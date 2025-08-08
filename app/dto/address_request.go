package dto

// AddressListRequest 钱包地址列表
type AddressListRequest struct {
	PageRequest
	Address string `query:"address" form:"address"`
	Status  string `query:"status" form:"status"`
}
