package service

import (
	"strings"
	"wht-admin/app/dto"
	"wht-admin/app/model"
	"wht-admin/common/password"
	"wht-admin/common/utils"
	"wht-admin/framework/dal"
)

type MerchantService struct{}

// CreateMerchant 新增商户
func (s *MerchantService) CreateMerchant(param dto.SaveMerchant) (int, error) {

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
		PayPwd:            param.PayPwd,
	}

	result := dal.Gorm.Model(model.WMerchant{}).Create(m)
	if result.Error != nil {
		return 0, result.Error
	}
	return m.MId, nil // 返回新增的 MId
}

// UpdateMerchant 更新商户
func (s *MerchantService) UpdateMerchant(param dto.SaveMerchant) error {

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

// GetMerchantList 商户列表
func (s *MerchantService) GetMerchantList(param dto.MerchantListRequest, isPaging bool) ([]dto.MerchantListResponse, int) {
	var count int64
	merchantList := make([]dto.MerchantListResponse, 0)

	query := dal.Gorm.Model(model.WMerchant{}).Order("w_merchant.m_id desc")

	query.Where("user_type = ?", 1)
	if param.MerchantName != "" {
		query.Where("username LIKE ?", "%"+param.MerchantName+"%")
	}

	if param.Status != "" {
		query.Where("status = ?", param.Status)
	}
	if isPaging {
		query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}
	query.Find(&merchantList)

	return merchantList, int(count)
}

// GetMerchantByMerchantId 根据商户id查询菜单
func (s *MerchantService) GetMerchantByMerchantId(merchantId int) dto.MerchantDetailResponse {

	var merchant dto.MerchantDetailResponse

	dal.Gorm.Model(model.WMerchant{}).Where("m_id = ?", merchantId).Last(&merchant)

	return merchant
}

// GetMerchantByMerchantName 根据商户账号名称查询商户
func (s *MerchantService) GetMerchantByMerchantName(merchantName string) dto.MerchantDetailResponse {

	var merchant dto.MerchantDetailResponse

	dal.Gorm.Model(model.WMerchant{}).Where("username = ?", merchantName).Last(&merchant)

	return merchant
}

// GetMerchantByAppId 根据appId查询商户
func (s *MerchantService) GetMerchantByAppId(appId string) int64 {
	// ✅ 唯一性校验
	var count int64
	dal.Gorm.Model(model.WMerchant{}).Where("app_id = ?", appId).Count(&count)

	return count
}

// GetDropDownList 商户下拉列表
func (s *MerchantService) GetDropDownList() []dto.MerchantDropDownListResponse {
	merchantList := make([]dto.MerchantDropDownListResponse, 0)

	query := dal.Gorm.Model(model.WMerchant{}).Where("status = ?", 1).Where("user_type = ?", 1).Order("w_merchant.m_id desc")

	query.Find(&merchantList)

	return merchantList
}

// UpdateMerchantPwd 更新商户密码
func (s *MerchantService) UpdateMerchantPwd(param dto.UpdateMerchantPwdRequest) error {
	var originalPwd = param.LoginPwd
	param.PayPwd = password.Generate(param.PayPwd)
	param.LoginPwd = password.Generate(param.LoginPwd)
	err := dal.Gorm.Model(model.WMerchant{}).Where("m_id = ?", param.MId).Updates(&model.WMerchant{
		Password: param.LoginPwd,
		PayPwd:   param.PayPwd,
	}).Error
	if err != nil {
		return err
	}
	var salt = utils.RandomString(5)
	var shopPwd = utils.MakeMd5(strings.Trim(originalPwd, " ") + salt)
	if err := (&ShopAdminService{}).UpdateShopAdminPwd(dto.UpdateShopAdmin{
		MId:      uint(param.MId),
		Password: shopPwd,
		Salt:     salt,
	}); err != nil {
		return err
	}
	return nil
}
