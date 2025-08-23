package service

import (
	"wht-admin/app/dto"
	"wht-admin/app/model"
	"wht-admin/framework/dal"
)

type CollectionAddressService struct{}

// GetCollectionAddressList 归集列表
func (s *CollectionAddressService) GetCollectionAddressList(param dto.CollectionAddressListRequest, isPaging bool) ([]dto.CollectionAddressListResponse, int) {
	var count int64
	collectionAddressList := make([]dto.CollectionAddressListResponse, 0)

	query := dal.Gorm.Model(model.WCollectionAddress{}).Order("w_collection_address.id desc")

	if param.Address != "" {
		query.Where("address LIKE ?", "%"+param.Address+"%")
	}

	if isPaging {
		query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}
	query.Find(&collectionAddressList)

	return collectionAddressList, int(count)
}

// GetCollectionAddressByCollectionAddressId 根据归集id查询归集
func (s *CollectionAddressService) GetCollectionAddressByCollectionAddressId(id int) dto.CollectionAddressDetailResponse {

	var collectionAddress dto.CollectionAddressDetailResponse

	dal.Gorm.Model(model.WCollectionAddress{}).Where("id = ?", id).Last(&collectionAddress)

	return collectionAddress
}

// GetCollectionAddressByAddress 根据归集名称查询归集
func (s *CollectionAddressService) GetCollectionAddressByAddress(address string) dto.CollectionAddressDetailResponse {

	var collectionAddressDetail dto.CollectionAddressDetailResponse

	dal.Gorm.Model(model.WCollectionAddress{}).Where("address = ?", address).Last(&collectionAddressDetail)

	return collectionAddressDetail
}

// CreateCollectionAddress 新增归集钱包地址
func (s *CollectionAddressService) CreateCollectionAddress(param dto.SaveCollectionAddress) error {

	return dal.Gorm.Model(model.WCollectionAddress{}).Create(&model.WCollectionAddress{
		MId:          param.MId,
		ChainSymbol:  param.ChainSymbol,
		Symbol:       param.Symbol,
		Protocol:     param.Protocol,
		CurrencyType: param.CurrencyType,
		Address:      param.Address,
		Status:       param.Status,
		CreateBy:     param.CreateBy,
		Remark:       param.Remark,
	}).Error
}

// UpdateCollectionAddress 更新归集钱包地址
func (s *CollectionAddressService) UpdateCollectionAddress(param dto.SaveCollectionAddress) error {

	return dal.Gorm.Model(model.WCollectionAddress{}).Where("id = ?", param.MId).Updates(&model.WCollectionAddress{
		MId:          param.MId,
		ChainSymbol:  param.ChainSymbol,
		Symbol:       param.Symbol,
		Protocol:     param.Protocol,
		CurrencyType: param.CurrencyType,
		Address:      param.Address,
		Status:       param.Status,
		UpdateBy:     param.UpdateBy,
		Remark:       param.Remark,
	}).Error
}
