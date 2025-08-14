package service

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/model"
	"ruoyi-go/framework/dal"
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
func (s *MerchantChannelService) UpdateMerchantChannel(param dto.SaveMerchantChannel) error {

	return dal.Gorm.Model(model.WMerchantChannel{}).Where("id = ?", param.ID).Updates(&model.WMerchantChannel{
		Currency:     param.Currency,
		MId:          param.MId,
		SysChannelID: param.SysChannelID,
		UpChannelID:  param.UpChannelID,
		Status:       param.Status,
		DefaultRate:  param.DefaultRate,
		SingleFee:    param.SingleFee,
		Weight:       param.Weight,
		OrderRange:   param.OrderRange,
		UpdateBy:     param.UpdateBy,
		Remark:       param.Remark,
	}).Error
}

// GetMerchantChannelList 商户通道列表
func (s *MerchantChannelService) GetMerchantChannelList(param dto.MerchantChannelListRequest, isPaging bool) ([]dto.MerchantChannelListResponse, int) {
	var count int64
	merchantChannelList := make([]dto.MerchantChannelListResponse, 0)

	query := dal.Gorm.Table("w_merchant_channel AS a").
		Joins("JOIN w_pay_upstream_product AS b ON  a.up_channel_id = b.id").
		Joins("JOIN w_pay_upstream AS c ON b.upstream_id = c.id").
		Joins("JOIN w_pay_way AS d ON a.sys_channel_id = d.id").
		Joins("JOIN w_currency_code AS e ON a.currency = e.`code`").
		Joins("JOIN w_merchant AS f ON a.m_id = f.m_id").
		Select("a.id,a.create_by,a.create_time,a.order_range,a.m_id,a.currency,a.default_rate,a.single_fee,b.order_range as up_order_range,c.title as upstream_title,b.default_rate as up_default_rate, b.add_rate as up_add_rate,d.coding,d.type,e.country,f.nickname as merchant_title")
	query.Where("a.m_id = ?", param.MId)
	if param.Status != -1 {
		query.Where("a.status = ?", param.Status)
	}
	if isPaging {
		query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}
	query.Order("a.id desc").Find(&merchantChannelList)

	return merchantChannelList, int(count)
}

// GetMerchantByMerchantChannelId 根据商户通道id查询商户通道信息
func (s *MerchantChannelService) GetMerchantByMerchantChannelId(mChannelId int) dto.MerchantChannelDetailResponse {

	var merchantChannel dto.MerchantChannelDetailResponse

	dal.Gorm.Model(model.WMerchantChannel{}).Where("id = ?", mChannelId).Last(&merchantChannel)

	return merchantChannel
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
