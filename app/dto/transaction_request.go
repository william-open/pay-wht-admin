package dto

// TransactionListRequest 钱包地址列表
type TransactionListRequest struct {
	PageRequest
	ToAddress string `query:"toAddress" form:"toAddress"`
}
