package validator

import (
	"errors"
	"wht-admin/app/dto"
)

// 添加商户验证
func CreateMerchantValidator(param dto.CreateMerchantRequest) error {

	if param.Username == "" {
		return errors.New("请输入账号")
	}

	if !isValidEmail(param.Username) {
		return errors.New("账号必须是邮箱地址")
	}

	//if param.NotifyUrl != "" && !strings.HasPrefix(param.NotifyUrl, "http") {
	//	return errors.New("回调地址" + param.NotifyUrl + "失败，地址必须以http(s)://开头")
	//}

	return nil
}

// 更新商户验证
func UpdateMerchantValidator(param dto.UpdateMerchantRequest) error {

	if param.MId <= 0 {
		return errors.New("参数错误")
	}

	if param.Username == "" {
		return errors.New("请输入账号")
	}

	//if !strings.HasPrefix(param.NotifyUrl, "http") {
	//	return errors.New("回调地址" + param.NotifyUrl + "失败，地址必须以http(s)://开头")
	//}

	return nil
}

// 更新商户白名单验证
func UpdateMerchantWhitelistValidator(param dto.UpdateWhitelistRequest) error {

	if param.MId <= 0 {
		return errors.New("参数错误")
	}

	if param.ApiIp == "" {
		return errors.New("请输入服务器下单IP")
	}

	if param.LoginApiIp == "" {
		return errors.New("请输入商户登录IP")
	}

	return nil
}

// UpdateMerchantPwdValidator 更新商户密码验证
func UpdateMerchantPwdValidator(param dto.UpdateMerchantPwdRequest) error {

	if param.MId <= 0 {
		return errors.New("参数错误")
	}

	if param.LoginPwd != "" {
		if len(param.LoginPwd) < 6 {
			return errors.New("登录密码至少6位")
		}
	}

	if param.PayPwd != "" {
		if len(param.PayPwd) < 6 {
			return errors.New("支付密码至少6位")
		}
	}

	return nil
}

// ResetMerchantGoogleSecretValidator 重置谷歌验证码
func ResetMerchantGoogleSecretValidator(param dto.ResetMerchantGoogleSecretRequest) error {

	if param.MId <= 0 {
		return errors.New("参数错误")
	}

	return nil
}
