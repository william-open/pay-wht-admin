package validator

import (
	"errors"
	"ruoyi-go/app/dto"
)

// 添加通道验证
func CreateChannelValidator(param dto.CreateChannelRequest) error {

	if param.Title == "" {
		return errors.New("请输入通道名称")
	}

	if param.Coding == "" {
		return errors.New("请输入通道编码")
	}

	if param.Type < 1 {
		return errors.New("请选择通道类型")
	}

	if param.Currency == "" {
		return errors.New("请选择国家")
	}
	return nil
}

// 更新通道验证
func UpdateChannelValidator(param dto.UpdateChannelRequest) error {

	if param.Id <= 0 {
		return errors.New("参数错误")
	}

	if param.Title == "" {
		return errors.New("请输入通道名称")
	}

	if param.Coding == "" {
		return errors.New("请输入通道编码")
	}

	if param.Type < 1 {
		return errors.New("请选择通道类型")
	}

	if param.Currency == "" {
		return errors.New("请选择国家")
	}
	return nil

	return nil
}

// 修改通道状态验证
func ChangeChannelStatusValidator(param dto.UpdateChannelRequest) error {

	if param.Id <= 0 {
		return errors.New("参数错误")
	}

	return nil
}
