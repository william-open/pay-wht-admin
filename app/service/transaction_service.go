package service

import (
	"wht-admin/app/dto"
	"wht-admin/app/model"
	"wht-admin/framework/dal"
)

type TransactionService struct{}

// GetTransactionList 交易列表
func (s *TransactionService) GetTransactionList(param dto.TransactionListRequest, isPaging bool) ([]dto.TransactionListResponse, int) {
	var count int64
	transactionList := make([]dto.TransactionListResponse, 0)

	query := dal.Gorm.Model(model.WTransaction{}).Order("w_transaction.id desc")

	if param.ToAddress != "" {
		query.Where("to_address LIKE ?", "%"+param.ToAddress+"%")
	}

	if isPaging {
		query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}
	query.Find(&transactionList)

	return transactionList, int(count)
}

// GetTransactionByTransactionId 根据交易id查询交易
func (s *TransactionService) GetTransactionByTransactionId(id int) dto.TransactionDetailResponse {

	var transaction dto.TransactionDetailResponse

	dal.Gorm.Model(model.WTransaction{}).Where("id = ?", id).Last(&transaction)

	return transaction
}

// GetTransactionByTransactionName 根据交易名称查询交易
func (s *TransactionService) GetTransactionByTxId(txId string) dto.TransactionDetailResponse {

	var transactionDetail dto.TransactionDetailResponse

	dal.Gorm.Model(model.WTransaction{}).Where("tx_id = ?", txId).Last(&transactionDetail)

	return transactionDetail
}
