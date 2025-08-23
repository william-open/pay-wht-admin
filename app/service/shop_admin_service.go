package service

import (
	"wht-admin/app/dto"
	"wht-admin/app/model"
	"wht-admin/framework/dal"
)

type ShopAdminService struct{}

// 新增用户
func (s *ShopAdminService) CreateShopAdmin(param dto.SaveShopAdmin) error {

	tx := dal.Gorm.Begin()

	user := model.SAuthAdmin{
		DeptId:        param.DeptId,
		PostId:        param.DeptId,
		AppId:         param.AppId,
		MId:           param.MId,
		Username:      param.Username,
		Nickname:      param.Nickname,
		Password:      param.Password,
		Avatar:        param.Avatar,
		Role:          param.Role,
		Salt:          param.Salt,
		Sort:          param.Sort,
		IsMultipoint:  param.IsMultipoint,
		IsDisable:     param.IsDisable,
		IsDelete:      param.IsDelete,
		LastLoginIp:   param.LastLoginIp,
		LastLoginTime: param.LastLoginTime,
		CreateTime:    param.CreateTime,
		UpdateTime:    param.UpdateTime,
		DeleteTime:    param.DeleteTime,
	}

	if err := tx.Model(model.SAuthAdmin{}).Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// 更新用户
func (s *ShopAdminService) ShopAdminService(param dto.SaveShopAdmin) error {

	tx := dal.Gorm.Begin()

	if err := tx.Model(model.SAuthAdmin{}).Where("id = ?", param.ID).Updates(&model.SAuthAdmin{
		DeptId:        param.DeptId,
		PostId:        param.DeptId,
		MId:           param.MId,
		Username:      param.Username,
		Nickname:      param.Nickname,
		Password:      param.Password,
		Avatar:        param.Avatar,
		Role:          param.Role,
		Salt:          param.Salt,
		Sort:          param.Sort,
		IsMultipoint:  param.IsMultipoint,
		IsDisable:     param.IsDisable,
		IsDelete:      param.IsDelete,
		LastLoginIp:   param.LastLoginIp,
		LastLoginTime: param.LastLoginTime,
		CreateTime:    param.CreateTime,
		UpdateTime:    param.UpdateTime,
		DeleteTime:    param.DeleteTime,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// 新增商户默认部门
func (s *ShopAdminService) CreateShopDept(param dto.SaveShopDept) error {

	tx := dal.Gorm.Begin()

	user := model.SAuthDept{
		Name: param.Name,
		Duty: param.Duty,
		MId:  param.MId,
	}

	if err := tx.Model(model.SAuthDept{}).Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
