package businesscontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"wht-admin/app/dto"
	"wht-admin/app/service"
	"wht-admin/app/validator"
	"wht-admin/framework/response"
)

type UpstreamProductController struct{}

// 供应商通道产品列表
func (*UpstreamProductController) List(ctx *gin.Context) {

	var param dto.PayUpstreamProductListRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	channels, total := (&service.UpstreamProductService{}).GetPayUpstreamProductList(param, true)

	response.NewSuccess().SetPageData(channels, total).Json(ctx)
}

// 供应商通道产品详情
func (*UpstreamProductController) Detail(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	channel := (&service.UpstreamProductService{}).GetPayUpstreamProductById(id)

	response.NewSuccess().SetData("data", channel).Json(ctx)
}

// 新增供应商通道产品
func (*UpstreamProductController) Create(ctx *gin.Context) {

	var param dto.CreatePayUpstreamProductRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.CreatePayUpstreamProductValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if product := (&service.UpstreamProductService{}).GetPayUpstreamProductByTitle(param.Title); product.Id > 0 {
		response.NewError().SetMsg("新增供应商通道产品" + param.Title + "失败，上游名称已存在").Json(ctx)
		return
	}
	var paramExist = dto.ExistPayUpstreamProductRequest{
		UpstreamId:   param.UpstreamId,
		SysChannelId: param.SysChannelId,
		UpstreamCode: param.UpstreamCode,
		Currency:     param.Currency,
	}
	if productExist := (&service.UpstreamProductService{}).CheckUpstreamProductExist(paramExist); productExist.Id > 0 {
		response.NewError().SetMsg("新增供应商通道产品失败，上游供应商通道已存在").Json(ctx)
		return
	}

	if err := (&service.UpstreamProductService{}).CreatePayUpstreamProduct(dto.SavePayUpstreamProduct{
		Title:        param.Title,
		UpstreamId:   param.UpstreamId,
		Currency:     param.Currency,
		SysChannelId: param.SysChannelId,
		UpstreamCode: param.UpstreamCode,
		SuccessRate:  param.SuccessRate,
		DefaultRate:  param.DefaultRate,
		AddRate:      param.AddRate,
		Weight:       param.Weight,
		Status:       param.Status,
		OrderRange:   param.OrderRange,
		Remark:       param.Remark,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 更新供应商通道产品
func (*UpstreamProductController) Update(ctx *gin.Context) {

	var param dto.UpdatePayUpstreamProductRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.UpdatePayUpstreamProductValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if product := (&service.UpstreamProductService{}).GetPayUpstreamProductByTitle(param.Title); product.Id > 0 {
		response.NewError().SetMsg("新增供应商通道产品" + param.Title + "失败，上游名称已存在").Json(ctx)
		return
	}

	var paramExist = dto.ExistPayUpstreamProductRequest{
		UpstreamId:   param.UpstreamId,
		SysChannelId: param.SysChannelId,
		UpstreamCode: param.UpstreamCode,
		Currency:     param.Currency,
	}
	if productExist := (&service.UpstreamProductService{}).CheckUpstreamProductExist(paramExist); productExist.Id > 0 {
		response.NewError().SetMsg("新增供应商通道产品失败，上游供应商通道已存在").Json(ctx)
		return
	}

	if err := (&service.UpstreamProductService{}).UpdatePayUpstreamProduct(dto.SavePayUpstreamProduct{
		Id:           param.Id,
		Title:        param.Title,
		UpstreamId:   param.UpstreamId,
		Currency:     param.Currency,
		SysChannelId: param.SysChannelId,
		UpstreamCode: param.UpstreamCode,
		SuccessRate:  param.SuccessRate,
		DefaultRate:  param.DefaultRate,
		AddRate:      param.AddRate,
		Weight:       param.Weight,
		Status:       param.Status,
		OrderRange:   param.OrderRange,
		Remark:       param.Remark,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 修改供应商通道产品状态
func (*UpstreamProductController) ChangeStatus(ctx *gin.Context) {

	var param dto.UpdateStatusRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.ChangeCommonStatusValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}
	if err := (&service.UpstreamProductService{}).UpdateUpstreamProductStatus(dto.SaveStatus{
		Id:     param.Id,
		Status: param.Status,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 根据状态查询所有国家
func (*UpstreamProductController) GetListByStatus(ctx *gin.Context) {

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
