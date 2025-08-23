package validator

import (
	"errors"
	"wht-admin/app/dto"
)

// CreateCollectionAddressValidator 添加归集钱包地址验证
func CreateCollectionAddressValidator(param dto.CreateCollectionAddressRequest) error {

	if param.Address == "" {
		return errors.New("请输入钱包地址")
	}

	return nil
}

// UpdateCollectionAddressValidator 更新归集钱包地址验证
func UpdateCollectionAddressValidator(param dto.UpdateCollectionAddressRequest) error {

	if param.Id <= 0 {
		return errors.New("参数错误")
	}

	if param.Address == "" {
		return errors.New("请输入钱包地址")
	}
	return nil
}
