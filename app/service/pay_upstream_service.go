package service

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/model"
	"ruoyi-go/framework/dal"
)

type PayUpstreamService struct{}

// 新增
func (s *PayUpstreamService) CreatePayUpstream(param dto.SavePayUpstream) error {

	tx := dal.Gorm.Begin()

	payUpstream := model.WPayUpstream{
		Title:          param.Title,
		Type:           param.Type,
		Status:         param.Status,
		WayID:          param.WayId,
		Account:        param.Account,
		PayKey:         param.PayKey,
		ReceivingKey:   param.ReceivingKey,
		SuccessRate:    param.SuccessRate,
		OrderQuantity:  param.OrderQuantity,
		Rate:           param.Rate,
		AppID:          param.AppID,
		AppSecret:      param.AppSecret,
		UpdateTime:     param.UpdateTime,
		ControlStatus:  param.ControlStatus,
		Sort:           param.Sort,
		PayingMoney:    param.PayingMoney,
		MinMoney:       param.MinMoney,
		MaxMoney:       param.MaxMoney,
		PayStatus:      param.PayStatus,
		OutStatus:      param.OutStatus,
		PayAPI:         param.PayAPI,
		PayQueryAPI:    param.PayQueryAPI,
		PayoutAPI:      param.PayoutAPI,
		PayoutQueryAPI: param.PayoutQueryAPI,
		BalanceInquiry: param.BalanceInquiry,
		SendingAddress: param.SendingAddress,
		Supplementary:  param.Supplementary,
		Documentation:  param.Documentation,
		NeedQuery:      param.NeedQuery,
		IPWhiteList:    param.IPWhiteList,
		CallbackDomain: param.CallbackDomain,
		Remark:         param.Remark,
	}

	if err := tx.Model(model.WPayUpstream{}).Create(&payUpstream).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// 更新上游供应商
func (s *PayUpstreamService) UpdatePayUpstream(param dto.SavePayUpstream) error {

	tx := dal.Gorm.Begin()

	if err := tx.Model(model.WPayUpstream{}).Where("id = ?", param.Id).Updates(&model.WPayUpstream{
		Title:          param.Title,
		Type:           param.Type,
		Status:         param.Status,
		WayID:          param.WayId,
		Account:        param.Account,
		PayKey:         param.PayKey,
		ReceivingKey:   param.ReceivingKey,
		SuccessRate:    param.SuccessRate,
		OrderQuantity:  param.OrderQuantity,
		Rate:           param.Rate,
		AppID:          param.AppID,
		AppSecret:      param.AppSecret,
		UpdateTime:     param.UpdateTime,
		ControlStatus:  param.ControlStatus,
		Sort:           param.Sort,
		PayingMoney:    param.PayingMoney,
		MinMoney:       param.MinMoney,
		MaxMoney:       param.MaxMoney,
		PayStatus:      param.PayStatus,
		OutStatus:      param.OutStatus,
		PayAPI:         param.PayAPI,
		PayQueryAPI:    param.PayQueryAPI,
		PayoutAPI:      param.PayoutAPI,
		PayoutQueryAPI: param.PayoutQueryAPI,
		BalanceInquiry: param.BalanceInquiry,
		SendingAddress: param.SendingAddress,
		Supplementary:  param.Supplementary,
		Documentation:  param.Documentation,
		NeedQuery:      param.NeedQuery,
		IPWhiteList:    param.IPWhiteList,
		CallbackDomain: param.CallbackDomain,
		Remark:         param.Remark,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// 获取上游供应商列表
func (s *PayUpstreamService) GetPayUpstreamList(param dto.PayUpstreamListRequest, isPaging bool) ([]dto.PayUpstreamListResponse, int) {

	var count int64
	channels := make([]dto.PayUpstreamListResponse, 0)

	query := dal.Gorm.Model(model.WPayUpstream{}).Order("id desc")

	if param.Keyword != "" {
		query.Where("title LIKE ? OR account LIKE ? OR appid LIKE ?", "%"+param.Title+"%", "%"+param.Account+"%", "%"+param.AppID+"%")
	}

	if param.Status != "" {
		query.Where("status = ?", param.Status)
	}

	if param.Type > 0 {
		query.Where("type = ?", param.Type)
	}

	if isPaging {
		query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}

	query.Find(&channels)

	return channels, int(count)
}

// 获取上游供应商详情
func (s *PayUpstreamService) GetPayUpstreamById(id int) dto.PayUpstreamDetailResponse {

	var country dto.PayUpstreamDetailResponse

	dal.Gorm.Model(model.WPayUpstream{}).Where("id = ?", id).Last(&country)

	return country
}

// 根据上游供应商名称查询上游供应商
func (s *PayUpstreamService) GetPayUpstreamByTitle(title string) dto.PayUpstreamDetailResponse {

	var channel dto.PayUpstreamDetailResponse

	dal.Gorm.Model(model.WPayUpstream{}).Where("title = ?", title).Last(&channel)

	return channel
}

// 根据上游商户号查询
func (s *PayUpstreamService) GetPayUpstreamByAccount(account string) dto.PayUpstreamDetailResponse {

	var upstream dto.PayUpstreamDetailResponse

	dal.Gorm.Model(model.WPayUpstream{}).Where("account = ?", account).Last(&upstream)

	return upstream
}

// 获取所有上游供应商列表
func (s *PayUpstreamService) GetAlUpstreamList(status int) []dto.UpstreamDropDownListResponse {

	upstreamList := make([]dto.UpstreamDropDownListResponse, 0)

	query := dal.Gorm.Model(model.WPayUpstream{}).Where("status = ?", status).Order("id asc")

	query.Find(&upstreamList)

	return upstreamList
}

// 查询上游供应商通道列表
func (s *PayUpstreamService) GetUpstreamChannelList(param dto.PayUpstreamListRequest, isPaging bool) ([]dto.PayUpstreamListResponse, int) {

	var count int64
	channels := make([]dto.PayUpstreamListResponse, 0)

	query := dal.Gorm.Model(model.WPayUpstream{}).Order("id desc")

	if param.Keyword != "" {
		query.Where("title LIKE ? OR account LIKE ? OR appid LIKE ?", "%"+param.Title+"%", "%"+param.Account+"%", "%"+param.AppID+"%")
	}

	if param.Status != "" {
		query.Where("status = ?", param.Status)
	}

	if param.Type > 0 {
		query.Where("type = ?", param.Type)
	}

	if isPaging {
		query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}

	query.Find(&channels)

	return channels, int(count)
}
