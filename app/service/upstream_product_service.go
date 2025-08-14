package service

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/model"
	"ruoyi-go/framework/dal"
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

// 获取上游供应商通道详情
func (s *UpstreamProductService) GetPayUpstreamProductById(id int) dto.PayUpstreamProductDetailResponse {

	var upstreamProduct dto.PayUpstreamProductDetailResponse

	dal.Gorm.Model(model.WPayUpstreamProduct{}).Where("id = ?", id).Last(&upstreamProduct)

	return upstreamProduct
}

// 根据上游供应商通道名称查询通道名称
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
