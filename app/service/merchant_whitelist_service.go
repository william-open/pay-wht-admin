package service

import (
	"wht-admin/app/dto"
	"wht-admin/app/model"
	"wht-admin/common/utils"
	"wht-admin/framework/dal"
)

type MerchantWhitelistService struct{}

// CreateMerchantWhitelist 新增商户白名单
func (s *MerchantWhitelistService) CreateMerchantWhitelist(param dto.SaveMerchantWhitelist, ipList []string) error {

	var records []model.WMerchantWhitelist
	resultList := utils.DeduplicateIPs(ipList)
	for _, ip := range resultList {
		records = append(records, model.WMerchantWhitelist{
			MID:        param.MID,
			IPAddress:  ip,
			CanAdmin:   param.CanAdmin,
			CanPayout:  param.CanPayout,
			CanReceive: param.CanReceive,
			CreateBy:   param.CreateBy,
		})
	}

	result := dal.Gorm.Create(&records)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetMerchantWhitelist 商户白名单列表
func (s *MerchantWhitelistService) GetMerchantWhitelist(param dto.MerchantWhitelistRequest, isPaging bool) ([]dto.MerchantWhitelistResponse, int) {

	var count int64
	merchantChannelList := make([]dto.MerchantWhitelistResponse, 0)

	query := dal.Gorm.Table("w_merchant_whitelist AS a").
		Joins("JOIN w_merchant AS b ON  a.m_id = b.m_id").
		Select("a.id,a.ip_address,a.can_admin,a.can_payout,a.can_receive,a.create_time")
	query.Where("a.m_id = ?", param.MId)
	if isPaging {
		query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}
	query.Order("a.id desc").Find(&merchantChannelList)

	return merchantChannelList, int(count)
}

// ExistMerchantIpAddress 查询商户白名单IP是否存在
func (s *MerchantWhitelistService) ExistMerchantIpAddress(mId int64, ipAddress string) bool {

	var count int64

	dal.Gorm.Model(model.WMerchantWhitelist{}).Where("m_id = ?", mId).
		Where("ip_address = ?", ipAddress).
		Count(&count)

	return count > 0
}

// DelMerchantIpAddress 删除商户白名单IP
func (s *MerchantWhitelistService) DelMerchantIpAddress(ids []int) error {

	err := dal.Gorm.Unscoped().Model(model.WMerchantWhitelist{}).Where("id in ?", ids).
		Delete(&model.WMerchantWhitelist{}).Error

	return err
}
