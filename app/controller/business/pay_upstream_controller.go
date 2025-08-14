package businesscontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"ruoyi-go/app/dto"
	"ruoyi-go/app/service"
	"ruoyi-go/app/validator"
	"ruoyi-go/framework/response"
	"strconv"
)

type PayUpstreamController struct{}

// 通道供应商列表
func (*PayUpstreamController) List(ctx *gin.Context) {

	var param dto.PayUpstreamListRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	channels, total := (&service.PayUpstreamService{}).GetPayUpstreamList(param, true)

	response.NewSuccess().SetPageData(channels, total).Json(ctx)
}

// 通道供应商详情
func (*PayUpstreamController) Detail(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	channel := (&service.PayUpstreamService{}).GetPayUpstreamById(id)

	response.NewSuccess().SetData("data", channel).Json(ctx)
}

// 新增通道供应商
func (*PayUpstreamController) Create(ctx *gin.Context) {

	var param dto.CreatePayUpstreamRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.CreatePayUpstreamValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if upstream := (&service.PayUpstreamService{}).GetPayUpstreamByTitle(param.Title); upstream.Id > 0 {
		response.NewError().SetMsg("新增通道供应商" + param.Title + "失败，上游名称已存在").Json(ctx)
		return
	}

	if upstream := (&service.PayUpstreamService{}).GetPayUpstreamByAccount(param.Account); upstream.Id > 0 {
		response.NewError().SetMsg("新增通道供应商" + param.Account + "失败，上游商户号已存在").Json(ctx)
		return
	}

	wayIdJson := datatypes.JSON([]byte("{}"))
	param.WayId = string(wayIdJson)
	if err := (&service.PayUpstreamService{}).CreatePayUpstream(dto.SavePayUpstream{
		Title:          param.Title,
		Type:           param.Type,
		WayId:          param.WayId,
		Account:        param.Account,
		PayKey:         param.PayKey,
		ReceivingKey:   param.ReceivingKey,
		SuccessRate:    param.SuccessRate,
		Rate:           param.Rate,
		AppID:          param.AppID,
		AppSecret:      param.AppSecret,
		ControlStatus:  param.ControlStatus,
		Status:         param.Status,
		PayAPI:         param.PayAPI,
		PayQueryAPI:    param.PayQueryAPI,
		PayoutAPI:      param.PayoutAPI,
		PayoutQueryAPI: param.PayoutQueryAPI,
		BalanceInquiry: param.BalanceInquiry,
		SendingAddress: param.SendingAddress,
		Supplementary:  param.Supplementary,
		Documentation:  param.Documentation,
		NeedQuery:      param.NeedQuery,
		IPWhiteList:    param.IPWhiteList,
		CallbackDomain: param.CallbackDomain,
		Remark:         param.Remark,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 更新通道供应商
func (*PayUpstreamController) Update(ctx *gin.Context) {

	var param dto.UpdatePayUpstreamRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.UpdatePayUpstreamValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if upstream := (&service.PayUpstreamService{}).GetPayUpstreamByTitle(param.Title); upstream.Id > 0 {
		response.NewError().SetMsg("新增通道供应商" + param.Title + "失败，上游名称已存在").Json(ctx)
		return
	}

	if upstream := (&service.PayUpstreamService{}).GetPayUpstreamByAccount(param.Account); upstream.Id > 0 {
		response.NewError().SetMsg("新增通道供应商" + param.Account + "失败，上游商户号已存在").Json(ctx)
		return
	}

	if err := (&service.PayUpstreamService{}).UpdatePayUpstream(dto.SavePayUpstream{
		Id:             param.Id,
		Title:          param.Title,
		Type:           param.Type,
		WayId:          param.WayId,
		Account:        param.Account,
		PayKey:         param.PayKey,
		ReceivingKey:   param.ReceivingKey,
		SuccessRate:    param.SuccessRate,
		Rate:           param.Rate,
		AppID:          param.AppID,
		AppSecret:      param.AppSecret,
		ControlStatus:  param.ControlStatus,
		Status:         param.Status,
		PayAPI:         param.PayAPI,
		PayQueryAPI:    param.PayQueryAPI,
		PayoutAPI:      param.PayoutAPI,
		PayoutQueryAPI: param.PayoutQueryAPI,
		BalanceInquiry: param.BalanceInquiry,
		SendingAddress: param.SendingAddress,
		Supplementary:  param.Supplementary,
		Documentation:  param.Documentation,
		NeedQuery:      param.NeedQuery,
		IPWhiteList:    param.IPWhiteList,
		CallbackDomain: param.CallbackDomain,
		Remark:         param.Remark,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 修改通道供应商状态
func (*PayUpstreamController) ChangeStatus(ctx *gin.Context) {

	var param dto.UpdatePayUpstreamRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.ChangePayUpstreamStatusValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}
	if err := (&service.PayUpstreamService{}).UpdatePayUpstream(dto.SavePayUpstream{
		Id:     param.Id,
		Status: param.Status,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 根据状态查询所有国家
func (*PayUpstreamController) GetListByStatus(ctx *gin.Context) {

	var param dto.QueryUpstreamByStatusRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	status, err := strconv.Atoi(param.Status)
	if err != nil {
		fmt.Println("类型转化错误:", err)
		return
	}

	upstreamList := (&service.PayUpstreamService{}).GetAlUpstreamList(status)

	response.NewSuccess().SetData("data", upstreamList).Json(ctx)
}

// 查询供应商通道列表
func (*PayUpstreamController) GetUpChannelList(ctx *gin.Context) {

	var param dto.UpstreamProductRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	upstreamChannel := (&service.UpstreamProductService{}).GeUpstreamProductByCond(param)

	response.NewSuccess().SetData("data", upstreamChannel).Json(ctx)
}
