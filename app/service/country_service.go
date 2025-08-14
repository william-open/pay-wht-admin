package service

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/model"
	"ruoyi-go/framework/dal"
)

type CountryService struct{}

// 新增国家
func (s *CountryService) CreateCountry(param dto.SaveCountry) error {

	tx := dal.Gorm.Begin()

	country := model.WCountry{
		Code:    param.Code,
		NameEn:  param.NameEn,
		NameZh:  param.NameZh,
		Symbol:  param.Symbol,
		Country: param.Country,
		IsOpen:  param.IsOpen,
	}

	if err := tx.Model(model.WCountry{}).Create(&country).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// 更新国家
func (s *CountryService) UpdateCountry(param dto.SaveCountry) error {

	tx := dal.Gorm.Begin()

	if err := tx.Model(model.WCountry{}).Where("id = ?", param.Id).Updates(&model.WCountry{
		Code:    param.Code,
		NameEn:  param.NameEn,
		NameZh:  param.NameZh,
		Symbol:  param.Symbol,
		Country: param.Country,
		IsOpen:  param.IsOpen,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// 获取国家列表
func (s *CountryService) GetCountryList(param dto.CountryListRequest, isPaging bool) ([]dto.CountryListResponse, int) {

	var count int64
	countrys := make([]dto.CountryListResponse, 0)

	query := dal.Gorm.Model(model.WCountry{}).Order("id desc")

	if param.Country != "" {
		query.Where("country LIKE ?", "%"+param.Country+"%")
	}

	if param.Code != "" {
		query.Where("code LIKE ?", "%"+param.Code+"%")
	}

	if param.IsOpen != "" {
		query.Where("is_open = ?", param.IsOpen)
	}

	if isPaging {
		query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}

	query.Find(&countrys)

	return countrys, int(count)
}

// 获取国家详情
func (s *CountryService) GetCountryById(id int) dto.CountryDetailResponse {

	var country dto.CountryDetailResponse

	dal.Gorm.Model(model.WCountry{}).Where("id = ?", id).Last(&country)

	return country
}

// 根据币种代码查询国家
func (s *CountryService) GetCountryByCode(code string) dto.CountryDetailResponse {

	var country dto.CountryDetailResponse

	dal.Gorm.Model(model.WCountry{}).Where("code = ?", code).Last(&country)

	return country
}

// 根据货币符号查询国家
func (s *CountryService) GetCountryBySymbol(symbol string) dto.CountryDetailResponse {

	var country dto.CountryDetailResponse

	dal.Gorm.Model(model.WCountry{}).Where("symbol = ?", symbol).Last(&country)

	return country
}

// 获取所有国家列表
func (s *CountryService) GetAllCountryList(status int) []dto.CountryListResponse {

	countrys := make([]dto.CountryListResponse, 0)

	query := dal.Gorm.Model(model.WCountry{}).Where("is_open = ?", status).Order("id desc")

	query.Find(&countrys)

	return countrys
}
