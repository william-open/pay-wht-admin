package validator

import (
	"errors"
	"wht-admin/app/dto"
)

// 添加商户通道验证
func CreateMerchantChannelValidator(param dto.CreateMerchantChannelRequest) error {

	if param.SysChannelID < 1 {
		return errors.New("请选择系统通道编码")
	}

	if param.UpChannelID < 1 {
		return errors.New("请选择上游通道编码")
	}

	if param.Currency == "" {
		return errors.New("请选择货币")
	}

	if param.MId < 1 {
		return errors.New("请选择商户")
	}

	if len(param.UpstreamProducts) == 0 {
		return errors.New("请选择上游通道产品")
	}
	return nil
}

// 更新商户通道验证
func UpdateMerchantChannelValidator(param dto.UpdateMerchantChannelRequest) error {

	if param.ID <= 0 {
		return errors.New("参数错误")
	}

	if param.SysChannelID < 1 {
		return errors.New("请选择系统通道编码")
	}

	if param.UpChannelID < 1 {
		return errors.New("请选择上游通道编码")
	}

	if param.Currency == "" {
		return errors.New("请选择货币")
	}

	if param.MId < 1 {
		return errors.New("请选择商户")
	}

	if len(param.UpstreamProducts) == 0 {
		return errors.New("请选择上游通道产品")
	}

	return nil
}

// 修改商户通道状态验证
func ChangeMerchantChannelStatusValidator(param dto.UpdateMerchantChannelRequest) error {

	if param.ID <= 0 {
		return errors.New("参数错误")
	}

	return nil
}
