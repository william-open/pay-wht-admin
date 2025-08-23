package validator

import (
	"errors"
	"strings"
	"wht-admin/app/dto"
)

// 添加代理验证
func CreateAgentValidator(param dto.CreateAgentRequest) error {

	if param.Username == "" {
		return errors.New("请输入账号")
	}

	if !isValidEmail(param.Username) {
		return errors.New("账号必须是邮箱地址")
	}

	if param.NotifyUrl != "" && !strings.HasPrefix(param.NotifyUrl, "http") {
		return errors.New("回调地址" + param.NotifyUrl + "失败，地址必须以http(s)://开头")
	}

	return nil
}

// 更新代理验证
func UpdateAgentValidator(param dto.UpdateAgentRequest) error {

	if param.MId <= 0 {
		return errors.New("参数错误")
	}

	if param.Username == "" {
		return errors.New("请输入账号")
	}

	if !strings.HasPrefix(param.NotifyUrl, "http") {
		return errors.New("回调地址" + param.NotifyUrl + "失败，地址必须以http(s)://开头")
	}

	return nil
}

// 更新代理白名单验证
func UpdateAgentWhitelistValidator(param dto.UpdateWhitelistRequest) error {

	if param.MId <= 0 {
		return errors.New("参数错误")
	}

	if param.ApiIp == "" {
		return errors.New("请输入服务器下单IP")
	}

	if param.LoginApiIp == "" {
		return errors.New("请输入代理登录IP")
	}

	return nil
}
