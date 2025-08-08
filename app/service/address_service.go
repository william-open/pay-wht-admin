package service

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/model"
	"ruoyi-go/framework/dal"
)

type AddressService struct{}

// GetAddressList 币种列表
func (s *AddressService) GetAddressList(param dto.AddressListRequest, isPaging bool) ([]dto.AddressListResponse, int) {
	var count int64
	addressList := make([]dto.AddressListResponse, 0)

	query := dal.Gorm.Model(model.WAddress{}).Order("w_address.address_id desc")

	if param.Address != "" {
		query.Where("address LIKE ?", "%"+param.Address+"%")
	}

	if param.Status != "" {
		query.Where("status = ?", param.Status)
	}
	if isPaging {
		query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}
	query.Find(&addressList)

	return addressList, int(count)
}

// GetAddressByAddressId 根据币种id查询币种
func (s *AddressService) GetAddressByAddressId(addressId int) dto.AddressDetailResponse {

	var address dto.AddressDetailResponse

	dal.Gorm.Model(model.WAddress{}).Where("address_id = ?", addressId).Last(&address)

	return address
}

// GetAddressByAddressName 根据币种名称查询币种
func (s *AddressService) GetAddressByAddressName(address string) dto.AddressDetailResponse {

	var addressDetail dto.AddressDetailResponse

	dal.Gorm.Model(model.WAddress{}).Where("address = ?", address).Last(&addressDetail)

	return addressDetail
}
