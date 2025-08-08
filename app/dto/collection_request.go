package dto

// CollectionListRequest 归集列表
type CollectionListRequest struct {
	PageRequest
	ToAddress string `query:"toAddress" form:"toAddress"`
}
