package validator

import (
	"errors"
	"wht-admin/app/dto"
)

// 添加国家验证
func CreateCountryValidator(param dto.CreateCountryRequest) error {

	if param.Code == "" {
		return errors.New("请输入币种代码")
	}

	if param.Symbol == "" {
		return errors.New("请输入货币符号")
	}

	if param.Country == "" {
		return errors.New("请输入国家名称")
	}

	if param.NameEn == "" {
		return errors.New("请输入英文名称")
	}

	if param.NameZh == "" {
		return errors.New("请输入中文名称")
	}
	return nil
}

// 更新国家验证
func UpdateCountryValidator(param dto.UpdateCountryRequest) error {

	if param.Id <= 0 {
		return errors.New("参数错误")
	}

	if param.Code == "" {
		return errors.New("请输入币种代码")
	}

	if param.Symbol == "" {
		return errors.New("请输入货币符号")
	}

	if param.Country == "" {
		return errors.New("请输入国家名称")
	}

	if param.NameEn == "" {
		return errors.New("请输入英文名称")
	}

	if param.NameZh == "" {
		return errors.New("请输入中文名称")
	}

	return nil
}

// 修改国家状态验证
func ChangeCountryStatusValidator(param dto.UpdateCountryRequest) error {

	if param.Id <= 0 {
		return errors.New("参数错误")
	}

	if param.IsOpen == "" {
		return errors.New("请选择状态")
	}

	return nil
}
