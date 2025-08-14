package service

import (
	"log"
	"ruoyi-go/app/dto"
	"ruoyi-go/app/model"
	"ruoyi-go/framework/dal"
)

type MoneyLogService struct{}

// CreateMoneyLog 新增资金日志
func (s *MoneyLogService) CreateMoneyLog(param dto.SaveMoneyLog) (uint, error) {

	m := &model.WMoneyLog{
		UID:         param.UID,
		Money:       param.Money,
		OldBalance:  param.OldBalance,
		Balance:     param.Balance,
		Currency:    param.Currency,
		Type:        param.Type,
		Operator:    param.Operator,
		OrderNo:     param.OrderNo,
		Description: param.Description,
	}

	result := dal.Gorm.Model(model.WMerchant{}).Create(m)
	if result.Error != nil {
		return 0, result.Error
	}
	return m.ID, nil // 返回新增的 MId
}

// UpdateMoneyLog 更新资金日志
func (s *MoneyLogService) UpdateMoneyLog(param dto.SaveMoneyLog) error {

	return dal.Gorm.Model(model.WMoneyLog{}).Where("id = ?", param.ID).Updates(&model.WMoneyLog{
		UID:         param.UID,
		Money:       param.Money,
		OldBalance:  param.OldBalance,
		Balance:     param.Balance,
		Currency:    param.Currency,
		Type:        param.Type,
		Operator:    param.Operator,
		OrderNo:     param.OrderNo,
		Description: param.Description,
	}).Error
}

// GetMoneyLogList 资金日志列表
func (s *MoneyLogService) GetMoneyLogList(param dto.MoneyLogListRequest, isPaging bool) ([]dto.MoneyLogResponse, int) {
	var count int64
	moneyLogList := make([]dto.MoneyLogResponse, 0)

	query := dal.Gorm.Table("w_money_log AS a").
		Select("a.*,m.nickname as merchant_title,c.country").
		Joins("LEFT JOIN w_merchant as m ON m.m_id = a.uid").
		Joins("LEFT JOIN w_currency_code as c ON a.currency = c.code").
		Where("m.user_type = ?", param.UserType)

	if param.Nickname != "" {
		query.Where("m.nickname LIKE ?", "%"+param.Nickname+"%")
	}

	// 先计算总数
	if err := query.Count(&count).Error; err != nil {
		log.Printf("Count error: %v", err)
		return moneyLogList, 0
	}

	// 再分页
	if isPaging {
		query.Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}

	if err := query.Order("a.id desc").Scan(&moneyLogList).Error; err != nil {
		log.Printf("Scan error: %v", err)
		return moneyLogList, 0
	}

	return moneyLogList, int(count)
}
