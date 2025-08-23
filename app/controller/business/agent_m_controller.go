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

type AgentMController struct{}

// List 代理商户列表
func (*AgentMController) List(ctx *gin.Context) {

	var param dto.AgentMListRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}
	list, total := (&service.AgentMService{}).GetAgentMList(param, true)
	response.NewSuccess().SetPageData(list, total).Json(ctx)
}

// 代理商户详情
func (*AgentMController) Detail(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("Id"))

	detail := (&service.AgentMService{}).GetAgentMById(id)

	response.NewSuccess().SetData("data", detail).Json(ctx)
}

// Create 新增代理商户
func (*AgentMController) Create(ctx *gin.Context) {

	var param dto.CreateAgentMRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.CreateAgentMValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if exist := (&service.AgentMService{}).ExistByWhere(param.AId, param.MId, param.SysChannelId, param.Currency); exist.Id > 0 {
		response.NewError().SetMsg("新增代理商户失败，代理商户名称已存在").Json(ctx)
		return
	}
	_, err := (&service.AgentMService{}).CreateAgentM(dto.SaveAgentM{
		AId:          param.AId,
		MId:          param.MId,
		Currency:     param.Currency,
		SysChannelId: param.SysChannelId,
		DefaultRate:  param.DefaultRate,
		SingleFee:    param.SingleFee,
		CreateBy:     security.GetAuthUserName(ctx),
		Remark:       param.Remark,
	})
	if err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}
	response.NewSuccess().Json(ctx)
}

// Update 更新代理商户
func (*AgentMController) Update(ctx *gin.Context) {

	var param dto.UpdateAgentMRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.UpdateAgentMValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}
	exist := (&service.AgentMService{}).ExistByWhere(param.AId, param.MId, param.SysChannelId, param.Currency)
	if exist.Id > 0 && exist.Id != param.Id {
		response.NewError().SetMsg("修改代理商户账号失败，代理商户账号已存在").Json(ctx)
		return
	}

	if err := (&service.AgentMService{}).UpdateAgentM(dto.SaveAgentM{
		AId:          param.AId,
		MId:          param.MId,
		Currency:     param.Currency,
		SysChannelId: param.SysChannelId,
		DefaultRate:  param.DefaultRate,
		SingleFee:    param.SingleFee,
		Status:       param.Status,
		UpdateBy:     security.GetAuthUserName(ctx),
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// Remove 删除代理商户
func (*AgentMController) Remove(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("agentMId"))

	if detail := (&service.AgentMService{}).GetAgentMById(id); detail.Id > 0 {
		response.NewError().SetMsg("信息不存在，不允许删除").Json(ctx)
		return
	}

	if err := (&service.AgentMService{}).DeleteAgentM(id); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}
