package dto

// SaveCollectionAddress 保存归集钱包地址
type SaveCollectionAddress struct {
	Id           int    `json:"id" binding:"required"`
	MId          int    `json:"mId"`
	Currency     string `json:"currency"`
	Symbol       string `json:"symbol"`
	CurrencyType string `json:"currency_type"`
	Protocol     string `json:"protocol"`
	Address      string `json:"address"`
	ChainSymbol  string `json:"chain_symbol"`
	Status       string `json:"status"`
	CreateBy     string `json:"createBy"`
	UpdateBy     string `json:"updateBy"`
	Remark       string `json:"remark"`
}

// CollectionAddressListRequest 归集钱包地址列表
type CollectionAddressListRequest struct {
	PageRequest
	Address string `query:"address" form:"address"`
	Status  string `query:"status" form:"status"`
}

// CreateCollectionAddressRequest 新增归集钱包地址
type CreateCollectionAddressRequest struct {
	MId          int    `json:"mId" binding:"required"`
	Currency     string `json:"currency" binding:"required"`
	Symbol       string `json:"symbol" binding:"required"`
	CurrencyType string `json:"currencyType" binding:"required"`
	Protocol     string `json:"protocol" binding:"required"`
	Address      string `json:"address" binding:"required"`
	ChainSymbol  string `json:"chainSymbol" binding:"required"`
	Status       string `json:"status"`
	Remark       string `json:"remark"`
}

// UpdateCollectionAddressRequest 更新归集钱包地址
type UpdateCollectionAddressRequest struct {
	Id           int    `json:"id" binding:"required"`
	MId          int    `json:"mId" binding:"required"`
	Currency     string `json:"currency" binding:"required"`
	Symbol       string `json:"symbol" binding:"required"`
	CurrencyType string `json:"currencyType" binding:"required"`
	Protocol     string `json:"protocol" binding:"required"`
	Address      string `json:"address" binding:"required"`
	ChainSymbol  string `json:"chainSymbol" binding:"required"`
	Status       string `json:"status"`
	Remark       string `json:"remark"`
}
