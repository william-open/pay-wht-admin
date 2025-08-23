package validator

import (
	"errors"
	"wht-admin/app/dto"
)

// 添加代理商户验证
func CreateAgentMValidator(param dto.CreateAgentMRequest) error {

	if param.AId < 0 {
		return errors.New("请输入代理ID")
	}

	if param.MId < 0 {
		return errors.New("请输入商户ID")
	}

	if param.SysChannelId < 0 {
		return errors.New("请选择通道")
	}

	if param.DefaultRate < 0 {
		return errors.New("请输入抽佣比例")
	}

	if param.Currency == "" {
		return errors.New("请选择国家")
	}

	return nil
}

// 更新代理商户验证
func UpdateAgentMValidator(param dto.UpdateAgentMRequest) error {

	if param.AId < 0 {
		return errors.New("请输入代理ID")
	}

	if param.MId < 0 {
		return errors.New("请输入商户ID")
	}

	if param.SysChannelId < 0 {
		return errors.New("请选择通道")
	}

	if param.DefaultRate < 0 {
		return errors.New("请输入抽佣比例")
	}

	if param.Currency == "" {
		return errors.New("请选择国家")
	}
	return nil
}
