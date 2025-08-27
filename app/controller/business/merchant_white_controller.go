package businesscontroller

import (
	"github.com/gin-gonic/gin"
	"wht-admin/app/dto"
	"wht-admin/app/security"
	"wht-admin/app/service"
	"wht-admin/app/validator"
	"wht-admin/common/utils"
	"wht-admin/framework/response"
)

type MerchantWhitelistController struct{}

// List 商户白名单列表
func (*MerchantWhitelistController) List(ctx *gin.Context) {

	var param dto.MerchantWhitelistRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	list, total := (&service.MerchantWhitelistService{}).GetMerchantWhitelist(param, true)
	response.NewSuccess().SetPageData(list, total).Json(ctx)
}

// Create 新增商户白名单
func (*MerchantWhitelistController) Create(ctx *gin.Context) {

	var param dto.CreateMerchantWhitelistRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.CreateMerchantWhitelistValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	ipList := utils.HandleIpList(param.IPAddress)
	if len(ipList) == 0 {
		response.NewError().SetMsg("请输入IP地址，或者IP地址中有不合法的IP").Json(ctx)
		return
	}

	err := (&service.MerchantWhitelistService{}).CreateMerchantWhitelist(dto.SaveMerchantWhitelist{
		MID:        param.MID,
		CanReceive: param.CanReceive,
		CanPayout:  param.CanPayout,
		CanAdmin:   param.CanAdmin,
		CreateBy:   security.GetAuthUserName(ctx),
		Remark:     param.Remark,
	}, ipList)
	if err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// Remove 删除商户白名单
func (*MerchantWhitelistController) Remove(ctx *gin.Context) {

	ids, err := utils.StringToIntSlice(ctx.Param("ids"), ",")
	if err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := (&service.MerchantWhitelistService{}).DelMerchantIpAddress(ids); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}
