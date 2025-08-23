package service

import (
	"wht-admin/app/dto"
	"wht-admin/app/model"
	"wht-admin/framework/dal"
)

type CollectionService struct{}

// GetCollectionList 归集列表
func (s *CollectionService) GetCollectionList(param dto.CollectionListRequest, isPaging bool) ([]dto.CollectionListResponse, int) {
	var count int64
	collectionList := make([]dto.CollectionListResponse, 0)

	query := dal.Gorm.Model(model.WCollection{}).Order("w_collection.id desc")

	if param.ToAddress != "" {
		query.Where("to_address LIKE ?", "%"+param.ToAddress+"%")
	}

	if isPaging {
		query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}
	query.Find(&collectionList)

	return collectionList, int(count)
}

// GetCollectionByCollectionId 根据归集id查询归集
func (s *CollectionService) GetCollectionByCollectionId(id int) dto.CollectionDetailResponse {

	var collection dto.CollectionDetailResponse

	dal.Gorm.Model(model.WCollection{}).Where("id = ?", id).Last(&collection)

	return collection
}

// GetCollectionByCollectionName 根据归集名称查询归集
func (s *CollectionService) GetCollectionByTxId(txId string) dto.CollectionDetailResponse {

	var collectionDetail dto.CollectionDetailResponse

	dal.Gorm.Model(model.WCollection{}).Where("tx_id = ?", txId).Last(&collectionDetail)

	return collectionDetail
}
