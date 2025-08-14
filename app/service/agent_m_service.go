package service

import (
	"fmt"
	"log"
	"ruoyi-go/app/dto"
	"ruoyi-go/app/model"
	"ruoyi-go/framework/dal"
	"time"
)

type AgentMService struct{}

// CreateAgentM 绑定代理商户
func (s *AgentMService) CreateAgentM(param dto.SaveAgentM) (int64, error) {
	// 开启事务
	tx := dal.Gorm.Begin()
	if tx.Error != nil {
		return 0, fmt.Errorf("事务启动失败: %w", tx.Error)
	}

	// 使用defer+recover处理panic情况
	var success bool
	defer func() {
		if !success {
			if r := recover(); r != nil {
				tx.Rollback()
				panic(r) // 重新抛出panic
			}
		}
	}()

	// 创建代理商户记录
	m := &model.WAgentMerchant{
		MID:          param.MId,
		AID:          param.AId,
		Currency:     param.Currency,
		SysChannelID: param.SysChannelId,
		DefaultRate:  param.DefaultRate,
		SingleFee:    param.SingleFee,
		Status:       param.Status,
		CreateBy:     param.CreateBy,
		Remark:       param.Remark,
	}

	if err := tx.Model(model.WAgentMerchant{}).Create(m).Error; err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("创建代理商户记录失败: %w", err)
	}

	// 更新商户表的上级ID
	updateData := map[string]interface{}{
		"pid":         param.AId,
		"update_by":   param.UpdateBy,
		"update_time": time.Now(), // 确保更新时间准确
	}

	if err := tx.Model(model.WMerchant{}).
		Where("m_id = ?", param.MId).
		Updates(updateData).Error; err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("更新商户上级ID失败: %w", err)
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return 0, fmt.Errorf("事务提交失败: %w", err)
	}

	success = true
	return m.ID, nil
}

// UpdateAgentM 更新代理商户
func (s *AgentMService) UpdateAgentM(param dto.SaveAgentM) error {

	return dal.Gorm.Model(model.WAgentMerchant{}).Where("id = ?", param.Id).Updates(&model.WAgentMerchant{
		MID:          param.MId,
		AID:          param.AId,
		Currency:     param.Currency,
		SysChannelID: param.SysChannelId,
		DefaultRate:  param.DefaultRate,
		SingleFee:    param.SingleFee,
		Status:       param.Status,
		UpdateBy:     param.UpdateBy,
		Remark:       param.Remark,
	}).Error
}

// GetAgentMList 代理商户列表
func (s *AgentMService) GetAgentMList(param dto.AgentMListRequest, isPaging bool) ([]dto.AgentMListResponse, int) {
	var count int64
	AgentMList := make([]dto.AgentMListResponse, 0)

	query := dal.Gorm.Table("w_agent_merchant as a").
		Select("a.id,a.sys_channel_id,a.a_id,a.default_rate,a.single_fee,a.create_time,a.create_by,b.nickname as merchant_title,d.nickname as agent_title,c.title as channel_title,c.coding").
		Joins("LEFT JOIN w_merchant as b ON a.m_id = b.m_id").
		Joins("LEFT JOIN w_pay_way as c ON c.id = a.sys_channel_id").
		Joins("LEFT JOIN w_merchant as d ON d.m_id = a.a_id").
		Where("a.a_id = ?", param.AId)

	if param.MerchantName != "" {
		query.Where("b.nickname LIKE ?", "%"+param.MerchantName+"%")
	}

	if param.ChannelName != "" {
		query.Where("c.title = ?", param.ChannelName)
	}

	// 先计算总数
	if err := query.Count(&count).Error; err != nil {
		log.Printf("Count error: %v", err)
		return AgentMList, 0
	}

	// 再分页
	if isPaging {
		query.Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}

	if err := query.Order("a.id desc").Scan(&AgentMList).Error; err != nil {
		log.Printf("Scan error: %v", err)
		return AgentMList, 0
	}

	return AgentMList, int(count)
}

// GetAgentMById 根据id查询代理
func (s *AgentMService) GetAgentMById(id int) dto.AgentMDetailResponse {

	var AgentM dto.AgentMDetailResponse

	dal.Gorm.Model(model.WAgentMerchant{}).Where("m_id = ?", id).Last(&AgentM)

	return AgentM
}

// ExistByWhere 根据条件查询数据
func (s *AgentMService) ExistByWhere(aId, mId, sysChannelId int64, currency string) dto.AgentMDetailResponse {

	var AgentM dto.AgentMDetailResponse

	dal.Gorm.Model(model.WAgentMerchant{}).Where("a_id = ? and m_id = ? and sys_channel_id = ? and currency = ?", aId, mId, sysChannelId, currency).Last(&AgentM)

	return AgentM
}

// DeleteAgentM 删除代理商户
func (s *AgentMService) DeleteAgentM(id int) error {
	return dal.Gorm.Model(model.WAgentMerchant{}).Where("id = ?", id).Delete(&model.WAgentMerchant{}).Error
}
