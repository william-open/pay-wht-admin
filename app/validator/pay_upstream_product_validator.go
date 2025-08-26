package validator

import (
	"errors"
	"wht-admin/app/dto"
)

// 添加通道供应商产品验证
func CreatePayUpstreamProductValidator(param dto.CreatePayUpstreamProductRequest) error {

	if param.Title == "" {
		return errors.New("请输入供应商通道产品名称")
	}

	if param.UpstreamId < 0 {
		return errors.New("请输入供应商ID")
	}

	if param.Currency == "" {
		return errors.New("请选择供应商产品国家")
	}

	if param.SysChannelId < 1 {
		return errors.New("请选择系统通道编码")
	}

	if param.UpstreamCode == "" {
		return errors.New("请选择供应商通道产品编码")
	}

	if param.Weight < 1 {
		return errors.New("请设置通道权重")
	}

	if param.OrderRange == "" {
		return errors.New("请设置通道金额范围")
	}

	return nil
}

// 更新通道供应商产品验证
func UpdatePayUpstreamProductValidator(param dto.UpdatePayUpstreamProductRequest) error {

	if param.Id <= 0 {
		return errors.New("参数错误")
	}

	if param.Title == "" {
		return errors.New("请输入供应商通道产品名称")
	}

	if param.UpstreamId < 0 {
		return errors.New("请输入供应商ID")
	}

	if param.Currency == "" {
		return errors.New("请选择供应商产品国家")
	}

	if param.SysChannelId < 1 {
		return errors.New("请选择系统通道编码")
	}

	if param.UpstreamCode == "" {
		return errors.New("请选择供应商通道产品编码")
	}

	if param.Weight < 1 {
		return errors.New("请设置通道权重")
	}

	if param.OrderRange == "" {
		return errors.New("请设置通道金额范围")
	}

	return nil
}

// TestPayUpstreamProductValidator 测试上游供应商通道产品验证
func TestPayUpstreamProductValidator(param dto.TestCreatePayUpstreamProductRequest) error {

	if param.Id < 1 {
		return errors.New("请选择上游供应商通道产品")
	}

	if param.Amount == "" {
		return errors.New("请输入订单金额")
	}

	return nil
}
