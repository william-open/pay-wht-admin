package businesscontroller

import (
	"github.com/gin-gonic/gin"
	"ruoyi-go/app/dto"
	"ruoyi-go/app/security"
	"ruoyi-go/app/service"
	"ruoyi-go/app/validator"
	"ruoyi-go/common/password"
	"ruoyi-go/framework/response"
	"strconv"
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

	merchant := (&service.MerchantChannelService{}).GetMerchantByMerchantChannelId(id)

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

	var param dto.UpdateMerchantRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.UpdateMerchantValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}
	merchant := (&service.MerchantService{}).GetMerchantByMerchantName(param.Username)
	if merchant.MId > 0 && merchant.MId != param.MId {
		response.NewError().SetMsg("修改商户账号" + param.Username + "失败，商户账号已存在").Json(ctx)
		return
	}
	var editPassword = ""
	if param.Password == "" {
		editPassword = merchant.Password
	} else {
		editPassword = password.Generate(param.Password)
	}

	if err := (&service.MerchantService{}).UpdateMerchant(dto.SaveMerchant{
		MId:      param.MId,
		Username: param.Username,
		Password: editPassword,
		Nickname: param.Nickname,
		//CallbackSecretKey: param.CallbackSecretKey,
		//NotifyUrl:         param.NotifyUrl,
		Remark:   param.Remark,
		Status:   param.Status,
		UpdateBy: security.GetAuthUserName(ctx),
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// Remove 删除商户通道
func (*MerchantChannelController) Remove(ctx *gin.Context) {

	menuId, _ := strconv.Atoi(ctx.Param("menuId"))

	if (&service.MenuService{}).MenuHasChildren(menuId) {
		response.NewError().SetMsg("存在子菜单，不允许删除").Json(ctx)
		return
	}

	if (&service.MenuService{}).MenuExistRole(menuId) {
		response.NewError().SetMsg("菜单已分配，不允许删除").Json(ctx)
		return
	}

	if err := (&service.MenuService{}).DeleteMenu(menuId); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 修改商户通道状态
func (*MerchantChannelController) ChangeStatus(ctx *gin.Context) {

	var param dto.UpdateMerchantChannelRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.ChangeMerchantChannelStatusValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}
	if err := (&service.MerchantChannelService{}).UpdateMerchantChannel(dto.SaveMerchantChannel{
		ID:     param.ID,
		Status: param.Status,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}
