package businesscontroller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wht-admin/app/dto"
	"wht-admin/app/security"
	"wht-admin/app/service"
	"wht-admin/app/validator"
	"wht-admin/framework/response"
)

type MerchantChannelController struct{}

// List 商户通道列表
func (*MerchantChannelController) List(ctx *gin.Context) {

	var param dto.MerchantChannelListRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	list, total := (&service.MerchantChannelService{}).GetMerchantChannelList(param, true)
	response.NewSuccess().SetPageData(list, total).Json(ctx)
}

// Detail 商户通道详情
func (*MerchantChannelController) Detail(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	merchant := (&service.MerchantChannelService{}).GetDetailByMerchantChannelId(id)

	response.NewSuccess().SetData("data", merchant).Json(ctx)
}

// Create 新增商户通道
func (*MerchantChannelController) Create(ctx *gin.Context) {

	var param dto.CreateMerchantChannelRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.CreateMerchantChannelValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	for _, channel := range param.UpstreamProducts {

		if detail := (&service.MerchantChannelService{}).ExistMerchantChannel(param.MId, param.SysChannelID, channel, param.Currency); detail.ID > 0 {
			continue
		}
		_, err := (&service.MerchantChannelService{}).CreateMerchantChannel(dto.SaveMerchantChannel{
			MId:          param.MId,
			SysChannelID: param.SysChannelID,
			UpChannelID:  channel,
			Weight:       param.Weight,
			Currency:     param.Currency,
			DefaultRate:  param.DefaultRate,
			SingleFee:    param.SingleFee,
			OrderRange:   param.OrderRange,
			Status:       param.Status,
			CreateBy:     security.GetAuthUserName(ctx),
			Remark:       param.Remark,
		})
		if err != nil {
			response.NewError().SetMsg(err.Error()).Json(ctx)
			return
		}
	}

	response.NewSuccess().Json(ctx)
}

// Update 更新商户通道
func (*MerchantChannelController) Update(ctx *gin.Context) {

	var param dto.UpdateMerchantChannelRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.UpdateMerchantChannelValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := (&service.MerchantChannelService{}).UpdateMerchantChannel(dto.UpdateMerchantChannel{
		DefaultRate: param.DefaultRate,
		SingleFee:   param.SingleFee,
		Weight:      param.Weight,
		OrderRange:  param.OrderRange,
		ID:          param.ID,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// Remove 删除商户通道
func (*MerchantChannelController) Remove(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := (&service.MerchantChannelService{}).RemoveMerchantChannel(id); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// ChangeStatus 修改商户通道状态
func (*MerchantChannelController) ChangeStatus(ctx *gin.Context) {

	var param dto.UpdateMerchantChannelStatusRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.ChangeMerchantChannelStatusValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}
	if err := (&service.MerchantChannelService{}).UpdateMerchantChannelStatus(dto.UpdateMerchantChannelStatus{
		ID:     param.ID,
		Status: param.Status,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}
