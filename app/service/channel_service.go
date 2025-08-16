package service

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/model"
	"ruoyi-go/framework/dal"
)

type ChannelService struct{}

// 新增
func (s *ChannelService) CreateChannel(param dto.SaveChannel) error {

	tx := dal.Gorm.Begin()

	channel := model.WChannel{
		Title:       param.Title,
		Status:      param.Status,
		Max:         param.Max,
		Min:         param.Min,
		DefaultRate: param.DefaultRate,
		AddRate:     param.AddRate,
		Coding:      param.Coding,
		Type:        param.Type,
		Currency:    param.Currency,
		Remark:      param.Remark,
	}

	if err := tx.Model(model.WChannel{}).Create(&channel).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// 更新通道
func (s *ChannelService) UpdateChannel(param dto.SaveChannel) error {

	tx := dal.Gorm.Begin()

	if err := tx.Model(model.WChannel{}).Where("id = ?", param.Id).Updates(&model.WChannel{
		Title:       param.Title,
		Status:      param.Status,
		Max:         param.Max,
		Min:         param.Min,
		DefaultRate: param.DefaultRate,
		AddRate:     param.AddRate,
		Coding:      param.Coding,
		Type:        param.Type,
		Currency:    param.Currency,
		Remark:      param.Remark,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// 获取通道列表
func (s *ChannelService) GetChannelList(param dto.ChannelListRequest, isPaging bool) ([]dto.ChannelListResponse, int) {

	var count int64
	channels := make([]dto.ChannelListResponse, 0)

	query := dal.Gorm.Table("w_pay_way AS a").
		Joins("JOIN w_currency_code AS b ON a.currency = b.code").
		Select("a.*,b.country")

	if param.Keyword != "" {
		query.Where("a.title LIKE ? OR a.coding LIKE ?", "%"+param.Keyword+"%", "%"+param.Keyword+"%")
	}

	if param.Status != "" {
		query.Where("a.status = ?", param.Status)
	}

	if param.Type > 0 {
		query.Where("a.type = ?", param.Type)
	}

	if param.Currency != "" {
		query.Where("a.currency = ?", param.Currency)
	}

	if isPaging {
		query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}

	query.Order("a.id desc").Find(&channels)

	return channels, int(count)
}

// 获取通道详情
func (s *ChannelService) GetChannelById(id int) dto.ChannelDetailResponse {

	var country dto.ChannelDetailResponse

	dal.Gorm.Model(model.WChannel{}).Where("id = ?", id).Last(&country)

	return country
}

// 根据通道名称查询通道
func (s *ChannelService) GetChannelByTitle(title string) dto.ChannelDetailResponse {

	var channel dto.ChannelDetailResponse

	dal.Gorm.Model(model.WChannel{}).Where("title = ?", title).Last(&channel)

	return channel
}

// 根据通道查询通道
func (s *ChannelService) GetChannelByCoding(coding string) dto.ChannelDetailResponse {

	var channel dto.ChannelDetailResponse

	dal.Gorm.Model(model.WChannel{}).Where("coding = ?", coding).Last(&channel)

	return channel
}

// 获取所有下拉列表根据状态
func (s *ChannelService) GetAllDropDownList(param dto.QueryChannelByStatusRequest) []dto.DropDownListResponse {

	channelList := make([]dto.DropDownListResponse, 0)

	query := dal.Gorm.Model(model.WChannel{})

	if param.Status != "" {
		query.Where("status = ?", param.Status)
	}
	if param.Currency != "" {
		query.Where("currency = ?", param.Currency)
	}

	query.Order("id asc").Find(&channelList)

	return channelList
}

// 更新通道状态
func (s *ChannelService) UpdateChannelStatus(param dto.SaveChannelStatus) error {
	return dal.Gorm.Model(&model.WChannel{}).
		Where("id = ?", param.Id).
		Updates(map[string]interface{}{
			"status": param.Status,
		}).Error
}
