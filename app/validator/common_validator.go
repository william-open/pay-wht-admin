package validator

import (
	"errors"
	"regexp"
	"wht-admin/app/dto"
)

// 修改状态验证
func ChangeCommonStatusValidator(param dto.UpdateStatusRequest) error {

	if param.Id <= 0 {
		return errors.New("参数错误")
	}

	return nil
}

func isValidEmail(email string) bool {
	// 简单的邮箱正则，可以根据需要调整更严格
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
