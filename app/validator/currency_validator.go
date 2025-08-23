package validator

import (
	"errors"
	"wht-admin/app/dto"
)

// CreateCurrencyValidator 添加币种验证
func CreateCurrencyValidator(param dto.CreateCurrencyRequest) error {

	if param.Currency == "" {
		return errors.New("请输入币种名称")
	}

	return nil
}

// UpdateCurrencyValidator 更新币种验证
func UpdateCurrencyValidator(param dto.UpdateCurrencyRequest) error {

	if param.CurrencyId <= 0 {
		return errors.New("参数错误")
	}

	if param.Currency == "" {
		return errors.New("请输入币种名称")
	}

	return nil
}
