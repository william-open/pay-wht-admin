package service

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/model"
	"ruoyi-go/framework/dal"
)

type CurrencyService struct{}

// CreateCurrency 新增币种
func (s *CurrencyService) CreateCurrency(param dto.SaveCurrency) error {

	return dal.Gorm.Model(model.WCurrency{}).Create(&model.WCurrency{
		PId:             param.PId,
		Currency:        param.Currency,
		Symbol:          param.Symbol,
		Logo:            param.Logo,
		ContractAddress: param.ContractAddress,
		CurrencyType:    param.CurrencyType,
		Protocol:        param.Protocol,
		Decimals:        param.Decimals,
		Status:          param.Status,
		CreateBy:        param.CreateBy,
		Remark:          param.Remark,
	}).Error
}

// UpdateCurrency 更新币种
func (s *CurrencyService) UpdateCurrency(param dto.SaveCurrency) error {

	return dal.Gorm.Model(model.WCurrency{}).Where("currency_id = ?", param.CurrencyId).Updates(&model.WCurrency{
		PId:             param.PId,
		Currency:        param.Currency,
		Symbol:          param.Symbol,
		Logo:            param.Logo,
		ContractAddress: param.ContractAddress,
		CurrencyType:    param.CurrencyType,
		Protocol:        param.Protocol,
		Decimals:        param.Decimals,
		Status:          param.Status,
		UpdateBy:        param.UpdateBy,
		Remark:          param.Remark,
	}).Error
}

// GetCurrencyList 币种列表
func (s *CurrencyService) GetCurrencyList(param dto.CurrencyListRequest, isPaging bool) ([]dto.CurrencyListResponse, int) {
	var count int64
	currencyList := make([]dto.CurrencyListResponse, 0)

	query := dal.Gorm.Model(model.WCurrency{}).Order("w_currency.currency_id desc")

	if param.CurrencyName != "" {
		query.Where("currency LIKE ?", "%"+param.CurrencyName+"%")
	}

	if param.Status != "" {
		query.Where("status = ?", param.Status)
	}
	if isPaging {
		query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}
	query.Find(&currencyList)

	return currencyList, int(count)
}

// GetCurrencyByCurrencyId 根据币种id查询币种
func (s *CurrencyService) GetCurrencyByCurrencyId(currencyId int) dto.CurrencyDetailResponse {

	var currency dto.CurrencyDetailResponse

	dal.Gorm.Model(model.WCurrency{}).Where("currency_id = ?", currencyId).Last(&currency)

	return currency
}

// GetCurrencyByCurrencyName 根据币种名称查询币种
func (s *CurrencyService) GetCurrencyByCurrencyName(currencyName string) dto.CurrencyDetailResponse {

	var currency dto.CurrencyDetailResponse

	dal.Gorm.Model(model.WCurrency{}).Where("currency = ?", currencyName).Last(&currency)

	return currency
}

// DeleteCurrency 删除币种
func (s *CurrencyService) DeleteCurrency(currencyId int) error {
	return dal.Gorm.Model(model.WCurrency{}).Where("currency_id = ?", currencyId).Delete(&model.WCurrency{}).Error
}
