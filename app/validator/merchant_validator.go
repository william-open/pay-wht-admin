package validator

import (
	"errors"
	"regexp"
	"ruoyi-go/app/dto"
	"strings"
)

// 添加商户验证
func CreateMerchantValidator(param dto.CreateMerchantRequest) error {

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

// 更新商户验证
func UpdateMerchantValidator(param dto.UpdateMerchantRequest) error {

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

func isValidEmail(email string) bool {
	// 简单的邮箱正则，可以根据需要调整更严格
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
