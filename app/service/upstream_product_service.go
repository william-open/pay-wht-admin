package service

import (
	"log"
	"wht-admin/app/dto"
	"wht-admin/app/model"
	"wht-admin/config"
	"wht-admin/framework/dal"
)

type UpstreamProductService struct{}

// 新增
func (s *UpstreamProductService) CreatePayUpstreamProduct(param dto.SavePayUpstreamProduct) error {

	tx := dal.Gorm.Begin()

	payUpstreamChannel := model.WPayUpstreamProduct{
		Title:        param.Title,
		UpstreamId:   param.UpstreamId,
		Currency:     param.Currency,
		SysChannelId: param.SysChannelId,
		UpstreamCode: param.UpstreamCode,
		DefaultRate:  param.DefaultRate,
		AddRate:      param.AddRate,
		Weight:       param.Weight,
		OrderRange:   param.OrderRange,
		Remark:       param.Remark,
	}

	if err := tx.Model(model.WPayUpstreamProduct{}).Create(&payUpstreamChannel).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// 更新上游供应商通道
func (s *UpstreamProductService) UpdatePayUpstreamProduct(param dto.SavePayUpstreamProduct) error {

	tx := dal.Gorm.Begin()

	if err := tx.Model(model.WPayUpstreamProduct{}).Where("id = ?", param.Id).Updates(&model.WPayUpstreamProduct{
		Title:        param.Title,
		UpstreamId:   param.UpstreamId,
		Currency:     param.Currency,
		SysChannelId: param.SysChannelId,
		UpstreamCode: param.UpstreamCode,
		DefaultRate:  param.DefaultRate,
		AddRate:      param.AddRate,
		Weight:       param.Weight,
		OrderRange:   param.OrderRange,
		Remark:       param.Remark,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// 获取上游供应商通道列表
func (s *UpstreamProductService) GetPayUpstreamProductList(param dto.PayUpstreamProductListRequest, isPaging bool) ([]dto.PayUpstreamProductListResponse, int) {

	var count int64
	channels := make([]dto.PayUpstreamProductListResponse, 0)

	query := dal.Gorm.Table("w_pay_upstream_product AS a").
		Select("a.*,b.coding,b.type").
		Joins("LEFT JOIN w_pay_way AS b ON a.sys_channel_id = b.id")

	query.Where("a.upstream_id = ?", param.UpstreamId)
	if param.Keyword != "" {
		query.Where("a.title LIKE ? OR a.currency LIKE ? OR a.upstream_code LIKE ?", "%"+param.Title+"%", "%"+param.Currency+"%", "%"+param.UpstreamCode+"%")
	}

	if param.Status != "" {
		query.Where("a.status = ?", param.Status)
	}

	if isPaging {
		query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}

	query.Order("a.id desc").Find(&channels)

	return channels, int(count)
}

// GetPayUpstreamProductById 获取上游供应商通道详情
func (s *UpstreamProductService) GetPayUpstreamProductById(id int) dto.PayUpstreamProductDetailResponse {

	var upstreamProduct dto.PayUpstreamProductDetailResponse

	dal.Gorm.Model(model.WPayUpstreamProduct{}).Where("id = ?", id).Last(&upstreamProduct)

	return upstreamProduct
}

// GetTestPayUpstreamProductById 获取测试上游供应商通道详情
func (s *UpstreamProductService) GetTestPayUpstreamProductById(id int) dto.TestPayUpstreamProductDetailResponse {

	var resp dto.TestPayUpstreamProductDetailResponse

	dal.Gorm.Table("w_pay_upstream_product AS a").
		Joins("LEFT JOIN w_pay_way AS b ON a.sys_channel_id = b.id").
		Joins("LEFT JOIN w_pay_upstream AS c ON a.upstream_id = c.id").
		Select("a.currency,a.order_range AS product_order_range,a.title AS product_title ,c.title AS upstream_name,b.title AS channel_title,b.coding AS channel_code,a.upstream_code AS product_code").
		Where("a.id = ?", id).Find(&resp)

	// 测试下游商户信息
	var merchantCurrency dto.MerchantCurrencyDetailResponse
	var testMerchantId = config.Data.TestDownMerchant.MerchantID
	log.Printf("测试商户ID:%v", testMerchantId)
	dal.Gorm.Table("w_merchant AS a").
		Joins("LEFT JOIN w_merchant_money AS b ON a.m_id = b.uid").
		Select("a.m_id,a.nickname,b.currency,b.money,b.freeze_money").
		Where("a.m_id = ?", testMerchantId).
		Where("b.currency = ?", resp.Currency).
		Find(&merchantCurrency)
	resp.MerchantName = merchantCurrency.Nickname
	resp.MerchantBalance = merchantCurrency.Money

	return resp
}

// GetPayUpstreamProductByTitle 根据上游供应商通道名称查询通道名称
func (s *UpstreamProductService) GetPayUpstreamProductByTitle(title string) dto.PayUpstreamProductDetailResponse {

	var upstreamChannel dto.PayUpstreamProductDetailResponse

	dal.Gorm.Model(model.WPayUpstreamProduct{}).Where("title = ?", title).Last(&upstreamChannel)

	return upstreamChannel
}

// 根据上游商户号查询
func (s *UpstreamProductService) GetPayUpstreamProductByAccount(account string) dto.PayUpstreamProductDetailResponse {

	var upstreamChannel dto.PayUpstreamProductDetailResponse

	dal.Gorm.Model(model.WPayUpstreamProduct{}).Where("account = ?", account).Last(&upstreamChannel)

	return upstreamChannel
}

// 根据指定条件查询供应商通道
func (s *UpstreamProductService) GeUpstreamProductByCond(param dto.UpstreamProductRequest) []dto.UpstreamProductResponse {

	var upstreamChannel = make([]dto.UpstreamProductResponse, 0)

	dal.Gorm.Model(model.WPayUpstreamProduct{}).
		Where("status = ?", param.Status).
		Where("currency = ?", param.Currency).
		Where("sys_channel_id = ?", param.SysChannelId).
		Where("upstream_id = ?", param.UpstreamId).
		Find(&upstreamChannel)

	return upstreamChannel
}

// 根据指定条件查询供应商通道产品
func (s *UpstreamProductService) CheckUpstreamProductExist(param dto.ExistPayUpstreamProductRequest) dto.UpstreamProductResponse {

	var upstreamChannel dto.UpstreamProductResponse

	dal.Gorm.Model(model.WPayUpstreamProduct{}).
		Where("currency = ?", param.Currency).
		Where("sys_channel_id = ?", param.SysChannelId).
		Where("upstream_code = ?", param.UpstreamCode).
		Where("upstream_id = ?", param.UpstreamId).
		Last(&upstreamChannel)

	return upstreamChannel
}

// 更新上游供应商通道状态
func (s *UpstreamProductService) UpdateUpstreamProductStatus(param dto.SaveStatus) error {

	tx := dal.Gorm.Begin()

	if err := tx.Model(model.WPayUpstreamProduct{}).Where("id = ?", param.Id).Updates(&model.WPayUpstreamProduct{
		Status: param.Status,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// GetPayChannelTypeById 根据上游供应商通道ID查询通道类型
func (s *UpstreamProductService) GetPayChannelTypeById(id int64) dto.UpstreamProductTypeResponse {

	var resp dto.UpstreamProductTypeResponse

	dal.Gorm.Table("w_pay_upstream_product AS a").
		Joins("LEFT JOIN w_pay_way AS b ON a.sys_channel_id = b.id").
		Select("b.type,b.coding").
		Where("a.id = ?", id).
		Find(&resp)

	return resp
}
