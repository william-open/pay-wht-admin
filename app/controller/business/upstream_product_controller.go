package businesscontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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

// 测试供应商通道产品详情
func (*UpstreamProductController) TestDetail(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	channel := (&service.UpstreamProductService{}).GetTestPayUpstreamProductById(id)
	log.Printf("数据：%+v", channel)

	response.NewSuccess().SetData("data", channel).Json(ctx)
}

// TestOrderCreate 测试上游供应商通道产品
func (*UpstreamProductController) TestOrderCreate(ctx *gin.Context) {

	var param dto.TestCreatePayUpstreamProductRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.TestPayUpstreamProductValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	// 判断是代收还是代付，走不同service下单
	var channelType dto.UpstreamProductTypeResponse
	channelType = (&service.UpstreamProductService{}).GetPayChannelTypeById(param.Id)
	log.Printf("测试上游供应商,通道类型: %+v", channelType)
	param.PayType = channelType.Coding
	if channelType.Type == 1 { //代收
		result, err := (&service.TestingChannelService{}).CreateCollectOrder(param)
		if err != nil {
			response.NewError().SetMsg("测试供应商通道产品代收下单失败，失败" + err.Error()).Json(ctx)
			return
		}
		log.Printf("测试上游供应商,返回结果: %+v", result)
		response.NewSuccess().SetData("data", result).Json(ctx)
	} else { //代付
		result, err := (&service.TestingChannelService{}).CreatePayoutOrder(param)
		if err != nil {
			response.NewError().SetMsg("测试供应商通道产品代付下单失败，失败" + err.Error()).Json(ctx)
			return
		}
		log.Printf("测试上游供应商,返回结果: %+v", result)
		response.NewSuccess().SetData("data", result).Json(ctx)
	}

	//response.NewSuccess().Json(ctx)
}
