package validator

import (
	"errors"
	"wht-admin/app/dto"
)

// CreateMerchantWhitelistValidator 添加商户白名单验证
func CreateMerchantWhitelistValidator(param dto.CreateMerchantWhitelistRequest) error {

	if param.MID < 1 {
		return errors.New("系统错误")
	}

	if param.IPAddress == "" {
		return errors.New("请输入IP地址")
	}
	return nil
}

// RemoveMerchantWhitelistValidator 删除商户白名单验证
func RemoveMerchantWhitelistValidator(param dto.RemoveMerchantWhitelistRequest) error {

	if param.Id <= 0 {
		return errors.New("参数错误")
	}

	return nil
}
