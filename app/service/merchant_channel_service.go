package service

import (
	"log"
	"wht-admin/app/dto"
	"wht-admin/app/model"
	"wht-admin/framework/dal"
)

type MerchantChannelService struct{}

// CreateMerchantChannel 新增商户通道
func (s *MerchantChannelService) CreateMerchantChannel(param dto.SaveMerchantChannel) (int, error) {

	m := &model.WMerchantChannel{
		Currency:     param.Currency,
		MId:          param.MId,
		SysChannelID: param.SysChannelID,
		UpChannelID:  param.UpChannelID,
		DefaultRate:  param.DefaultRate,
		SingleFee:    param.SingleFee,
		Weight:       param.Weight,
		OrderRange:   param.OrderRange,
		UpdateBy:     param.UpdateBy,
		Status:       param.Status,
		CreateBy:     param.CreateBy,
		Remark:       param.Remark,
	}

	result := dal.Gorm.Model(model.WMerchantChannel{}).Create(m)
	if result.Error != nil {
		return 0, result.Error
	}
	return m.ID, nil // 返回新增的 ID
}

// UpdateMerchantChannel 更新商户通道
func (s *MerchantChannelService) UpdateMerchantChannel(param dto.UpdateMerchantChannel) error {

	return dal.Gorm.Model(model.WMerchantChannel{}).Where("id = ?", param.ID).Updates(&model.WMerchantChannel{
		DefaultRate: param.DefaultRate,
		SingleFee:   param.SingleFee,
		Weight:      param.Weight,
		OrderRange:  param.OrderRange,
	}).Error
}

// GetMerchantChannelList 商户通道列表
func (s *MerchantChannelService) GetMerchantChannelList(param dto.MerchantChannelListRequest, isPaging bool) ([]dto.MerchantChannelListResponse, int) {

	log.Printf("查询商户通道请求参数:%+v", param)
	var count int64
	merchantChannelList := make([]dto.MerchantChannelListResponse, 0)

	query := dal.Gorm.Table("w_merchant_channel AS a").
		Joins("JOIN w_pay_upstream_product AS b ON  a.up_channel_id = b.id").
		Joins("JOIN w_pay_upstream AS c ON b.upstream_id = c.id").
		Joins("JOIN w_pay_way AS d ON a.sys_channel_id = d.id").
		Joins("JOIN w_currency_code AS e ON a.currency = e.`code`").
		Joins("JOIN w_merchant AS f ON a.m_id = f.m_id").
		Select("a.status,a.weight,a.id,a.create_by,a.create_time,a.order_range,a.m_id,a.currency,a.default_rate,a.single_fee,b.order_range as up_order_range,c.title as upstream_title,b.default_rate as up_default_rate, b.add_rate as up_add_rate,d.coding,d.type,e.country,f.nickname as merchant_title,b.title as up_product_title,b.upstream_code,d.type")
	query.Where("a.m_id = ?", param.MId)
	if param.Status != nil {
		query.Where("a.status = ?", param.Status)
	}
	if param.Keyword != "" {
		query.Where("d.coding like ? OR d.title like ? OR b.title like ? OR b.upstream_code like ?", "%"+param.Keyword+"%", "%"+param.Keyword+"%", "%"+param.Keyword+"%", "%"+param.Keyword+"%")
	}
	if param.Type != nil {
		query.Where("d.type = ?", param.Type)
	}
	if isPaging {
		query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}
	query.Order("a.id desc").Find(&merchantChannelList)

	return merchantChannelList, int(count)
}

// GetDetailByMerchantChannelId 根据商户通道id查询商户通道信息
func (s *MerchantChannelService) GetDetailByMerchantChannelId(mChannelId int) dto.MerchantChannelDetailResponse {

	var detail dto.MerchantChannelDetailResponse

	dal.Gorm.Table("w_merchant_channel AS a").
		Joins("LEFT JOIN w_pay_way AS b ON a.sys_channel_id = b.id").
		Joins("LEFT JOIN w_pay_upstream_product AS c ON c.id = a.up_channel_id").
		Joins("LEFT JOIN w_currency_code as d ON d.code = a.currency").
		Joins("LEFT JOIN w_pay_upstream as e ON e.id = c.upstream_id").
		Select("a.currency,a.id,a.default_rate,a.single_fee,a.weight,a.order_range,b.coding,c.upstream_code,d.country,b.title,e.title as sys_title,e.title as upstream_title,c.title as upstream_channel_title").
		Where("a.id = ?", mChannelId).Find(&detail)

	return detail
}

// GetMerchantChannelBySysAndUp 根据商户ID+系统通道ID+上游通道ID+查询商户通道
func (s *MerchantChannelService) GetMerchantChannelBySysAndUp(mId, sysChannelId, upChannelId int64) dto.MerchantChannelDetailResponse {

	var merchantChannel dto.MerchantChannelDetailResponse

	dal.Gorm.Model(model.WMerchantChannel{}).Where("m_id = ?", mId).Where("sys_channel_id = ?", sysChannelId).Where("up_channel_id = ?", upChannelId).Last(&merchantChannel)

	return merchantChannel
}

// ExistMerchantChannel 查询商户产品通道是否存在
func (s *MerchantChannelService) ExistMerchantChannel(mId, sysChannelId, upChannelId int64, currency string) dto.MerchantChannelDetailResponse {

	var merchantChannel dto.MerchantChannelDetailResponse

	dal.Gorm.Model(model.WMerchantChannel{}).Where("currency = ?", currency).
		Where("m_id = ?", mId).
		Where("sys_channel_id = ?", sysChannelId).
		Where("up_channel_id = ?", upChannelId).
		Last(&merchantChannel)

	return merchantChannel
}

// UpdateMerchantChannelStatus 更新商户通道状态
func (s *MerchantChannelService) UpdateMerchantChannelStatus(param dto.UpdateMerchantChannelStatus) error {

	return dal.Gorm.Model(model.WMerchantChannel{}).Where("id = ?", param.ID).Updates(&model.WMerchantChannel{
		Status: param.Status,
	}).Error
}

// RemoveMerchantChannel  删除商户通道
func (s *MerchantChannelService) RemoveMerchantChannel(id int) error {

	return dal.Gorm.Unscoped().Model(model.WMerchantChannel{}).Where("id = ?", id).Delete(&model.WMerchantChannel{}).Error
}
