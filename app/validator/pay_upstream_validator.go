package validator

import (
	"errors"
	"ruoyi-go/app/dto"
)

// 添加通道供应商验证
func CreatePayUpstreamValidator(param dto.CreatePayUpstreamRequest) error {

	if param.Title == "" {
		return errors.New("请输入供应商名称")
	}

	if param.Account == "" {
		return errors.New("请输入商户账号")
	}

	if param.Type < 1 {
		return errors.New("请选择通道类型")
	}

	//if param.WayId == "" {
	//	return errors.New("请选择通道")
	//}

	if param.AppID == "" {
		return errors.New("请选择AppId")
	}

	if param.PayAPI == "" {
		return errors.New("请选择代收下单API")
	}

	if param.PayoutAPI == "" {
		return errors.New("请选择代付下单API")
	}

	if param.PayQueryAPI == "" {
		return errors.New("请选择代收查询地址")
	}

	if param.PayoutQueryAPI == "" {
		return errors.New("请选择代付查询地址")
	}

	return nil
}

// 更新通道供应商验证
func UpdatePayUpstreamValidator(param dto.UpdatePayUpstreamRequest) error {

	if param.Id <= 0 {
		return errors.New("参数错误")
	}

	if param.Title == "" {
		return errors.New("请输入供应商名称")
	}

	if param.Account == "" {
		return errors.New("请输入商户账号")
	}

	if param.Type < 1 {
		return errors.New("请选择通道类型")
	}

	//if param.WayId == "" {
	//	return errors.New("请选择通道")
	//}

	if param.AppID == "" {
		return errors.New("请选择AppId")
	}

	if param.PayKey == "" {
		return errors.New("请选择代收下单API")
	}

	if param.PayoutAPI == "" {
		return errors.New("请选择代付下单API")
	}

	if param.PayQueryAPI == "" {
		return errors.New("请选择代收查询地址")
	}

	if param.PayoutQueryAPI == "" {
		return errors.New("请选择代付查询地址")
	}

	return nil
}

// 修改通道供应商状态验证
func ChangePayUpstreamStatusValidator(param dto.UpdatePayUpstreamRequest) error {

	if param.Id <= 0 {
		return errors.New("参数错误")
	}

	return nil
}
