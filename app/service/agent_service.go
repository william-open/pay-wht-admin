package service

import (
	"log"
	"ruoyi-go/app/dto"
	"ruoyi-go/app/model"
	"ruoyi-go/framework/dal"
)

type AgentService struct{}

// CreateMerchant 新增代理
func (s *AgentService) CreateAgent(param dto.SaveAgent) (int, error) {

	m := &model.WMerchant{
		Username:          param.Username,
		Password:          param.Password,
		Nickname:          param.Nickname,
		CallbackSecretKey: param.CallbackSecretKey,
		NotifyUrl:         param.NotifyUrl,
		AesSecretKey:      param.AesSecretKey,
		PublicKey:         param.PublicKey,
		PrivateKey:        param.PrivateKey,
		ApiKey:            param.ApiKey,
		AppId:             param.AppId,
		Status:            param.Status,
		CreateBy:          param.CreateBy,
		Remark:            param.Remark,
		UpstreamId:        param.UpstreamId,
		Ways:              param.Ways,
		UserType:          param.UserType,
		PayType:           param.PayType,
	}

	result := dal.Gorm.Model(model.WMerchant{}).Create(m)
	if result.Error != nil {
		return 0, result.Error
	}
	return m.MId, nil // 返回新增的 MId
}

// UpdateAgent 更新代理
func (s *AgentService) UpdateAgent(param dto.SaveAgent) error {

	return dal.Gorm.Model(model.WMerchant{}).Where("m_id = ?", param.MId).Updates(&model.WMerchant{
		Username:          param.Username,
		Password:          param.Password,
		Nickname:          param.Nickname,
		CallbackSecretKey: param.CallbackSecretKey,
		NotifyUrl:         param.NotifyUrl,
		AesSecretKey:      param.AesSecretKey,
		PublicKey:         param.PublicKey,
		PrivateKey:        param.PrivateKey,
		AppId:             param.AppId,
		Status:            param.Status,
		UpdateBy:          param.UpdateBy,
		Remark:            param.Remark,
		UserType:          param.UserType,
		PayType:           param.PayType,
	}).Error
}

// GetAgentList 代理列表
func (s *AgentService) GetAgentList(param dto.AgentListRequest, isPaging bool) ([]dto.AgentListResponse, int) {
	var count int64
	agentList := make([]dto.AgentListResponse, 0)

	query := dal.Gorm.Model(model.WMerchant{}).
		Select("w_merchant.*, COUNT(sub.m_id) as sub_count").
		Joins("LEFT JOIN w_merchant as sub ON sub.pid = w_merchant.m_id").
		Where("w_merchant.user_type = ?", 2). // 代理类型固定为2
		Group("w_merchant.m_id")

	if param.AgentName != "" {
		query.Where("w_merchant.username LIKE ?", "%"+param.AgentName+"%")
	}

	if param.Status != "" {
		query.Where("w_merchant.status = ?", param.Status)
	}

	// 先计算总数
	if err := query.Count(&count).Error; err != nil {
		log.Printf("Count error: %v", err)
		return agentList, 0
	}

	// 再分页
	if isPaging {
		query.Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}

	if err := query.Order("w_merchant.m_id desc").Scan(&agentList).Error; err != nil {
		log.Printf("Scan error: %v", err)
		return agentList, 0
	}

	return agentList, int(count)
}

// GetAgentByAgentId 根据代理id查询代理
func (s *AgentService) GetAgentByAgentId(agentId int) dto.AgentDetailResponse {

	var agent dto.AgentDetailResponse

	dal.Gorm.Model(model.WMerchant{}).Where("m_id = ?", agentId).Last(&agent)

	return agent
}

// GetAgentByAgentName 根据代理账号名称查询代理
func (s *AgentService) GetAgentByAgentName(agentName string) dto.AgentDetailResponse {

	var agent dto.AgentDetailResponse

	dal.Gorm.Model(model.WMerchant{}).Where("username = ?", agentName).Last(&agent)

	return agent
}

// GetAgentByAppId 根据appId查询代理
func (s *AgentService) GetAgentByAppId(appId string) int64 {
	// ✅ 唯一性校验
	var count int64
	dal.Gorm.Model(model.WMerchant{}).Where("app_id = ?", appId).Count(&count)

	return count
}

// UpdateAgentWhitelist 更新代理白名单
func (s *AgentService) UpdateAgentWhitelist(param dto.SaveAgentWhitelist) error {

	return dal.Gorm.Model(model.WMerchant{}).Where("m_id = ?", param.MId).Updates(&model.WMerchant{
		ApiIp:      param.ApiIp,
		LoginApiIp: param.LoginApiIp,
	}).Error
}
