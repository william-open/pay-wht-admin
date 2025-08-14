package service

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/model"
	"ruoyi-go/framework/dal"
)

type MerchantAccountService struct{}

// GetMerchantAccountList 商户账户列表
func (s *MerchantAccountService) GetMerchantAccountList(param dto.MerchantAccountRequest, isPaging bool) ([]dto.MerchantAccountListResponse, int) {
	var count int64
	accountList := make([]dto.MerchantAccountListResponse, 0)

	// 先计算总数（不分组）
	countQuery := dal.Gorm.Table("w_merchant_money AS a").
		Joins("JOIN w_merchant as b ON a.uid = b.m_id").
		Joins("JOIN w_currency_code as c ON a.currency = c.code")

	if param.MerchantName != "" {
		countQuery.Where("b.nickname LIKE ?", "%"+param.MerchantName+"%")
	}
	countQuery.Group("a.uid").Count(&count)

	// 主查询
	query := dal.Gorm.Table("w_merchant_money AS a").
		Select(`
            a.uid,
			a.id,
            b.nickname as merchant_title,
            SUM(a.money) as total_money,
            SUM(a.freeze_money) as total_freeze_money
        `).
		Joins("JOIN w_merchant as b ON a.uid = b.m_id").
		Group("a.uid") // 确保GROUP BY包含所有非聚合字段

	if param.MerchantName != "" {
		query.Where("b.nickname LIKE ?", "%"+param.MerchantName+"%")
	}

	if isPaging {
		query.Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}

	query.Order("total_money desc").Find(&accountList)

	return accountList, int(count)
}

// GetAccountByMerchantId 根据商户id查询商户账户
func (s *MerchantAccountService) GetAccountByMerchantId(merchantId int) dto.MerchantAccountDetailResponse {

	var account dto.MerchantAccountDetailResponse

	dal.Gorm.Model(model.WMerchantMoney{}).Where("uid = ?", merchantId).Last(&account)

	return account
}

// GetMerchantAccountCurrencyList 商户账户货币明显列表
func (s *MerchantAccountService) GetMerchantAccountCurrencyList(param dto.MerchantAccountCurrencyRequest, isPaging bool) ([]dto.MerchantAccountCurrencyListResponse, int) {
	var count int64
	currencyList := make([]dto.MerchantAccountCurrencyListResponse, 0)

	// 先计算总数（不分组）
	countQuery := dal.Gorm.Table("w_merchant_money AS a").
		Joins("JOIN w_currency_code as c ON a.currency = c.code")

	countQuery.Where("a.uid = ?", param.MId)
	if param.Currency != "" {
		countQuery.Where("a.currency like ?", "%"+param.Currency+"%")
	}
	countQuery.Count(&count)

	// 主查询
	query := dal.Gorm.Table("w_merchant_money AS a").
		Select(`
				a.*,b.country
			
        `).
		Joins("JOIN w_currency_code as b ON a.currency = b.code")

	query.Where("a.uid = ?", param.MId)
	if param.Currency != "" {
		query.Where("a.currency like ?", "%"+param.Currency+"%")
	}

	if isPaging {
		query.Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}

	query.Order("a.money desc").Find(&currencyList)

	return currencyList, int(count)
}
