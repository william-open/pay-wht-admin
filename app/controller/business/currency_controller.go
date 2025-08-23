package businesscontroller

import (
	"strconv"
	"wht-admin/app/dto"
	"wht-admin/app/security"
	"wht-admin/app/service"
	"wht-admin/app/validator"
	"wht-admin/framework/response"

	"github.com/gin-gonic/gin"
)

type CurrencyController struct{}

// List 币种列表
func (*CurrencyController) List(ctx *gin.Context) {

	var param dto.CurrencyListRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	list, total := (&service.CurrencyService{}).GetCurrencyList(param, true)
	response.NewSuccess().SetPageData(list, total).Json(ctx)
}

// Detail 币种详情
func (*CurrencyController) Detail(ctx *gin.Context) {

	currencyId, _ := strconv.Atoi(ctx.Param("currencyId"))

	currency := (&service.CurrencyService{}).GetCurrencyByCurrencyId(currencyId)

	response.NewSuccess().SetData("data", currency).Json(ctx)
}

// Create 新增币种
func (*CurrencyController) Create(ctx *gin.Context) {

	var param dto.CreateCurrencyRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.CreateCurrencyValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if currency := (&service.CurrencyService{}).GetCurrencyByCurrencyName(param.Currency); currency.PId > 0 {
		response.NewError().SetMsg("新增币种" + param.Currency + "失败，币种名称已存在").Json(ctx)
		return
	}

	if err := (&service.CurrencyService{}).CreateCurrency(dto.SaveCurrency{
		PId:             param.PId,
		Currency:        param.Currency,
		Symbol:          param.Symbol,
		Logo:            param.Logo,
		ContractAddress: param.ContractAddress,
		CurrencyType:    param.CurrencyType,
		ChainSymbol:     param.ChainSymbol,
		Protocol:        param.Protocol,
		Decimals:        param.Decimals,
		Status:          param.Status,
		CreateBy:        security.GetAuthUserName(ctx),
		Remark:          param.Remark,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// Update 更新币种
func (*CurrencyController) Update(ctx *gin.Context) {

	var param dto.UpdateCurrencyRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.UpdateCurrencyValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}
	currency := (&service.CurrencyService{}).GetCurrencyByCurrencyName(param.Currency)
	if currency.PId > 0 && currency.PId != param.PId {
		response.NewError().SetMsg("修改币种账号" + param.Currency + "失败，币种账号已存在").Json(ctx)
		return
	}

	if err := (&service.CurrencyService{}).UpdateCurrency(dto.SaveCurrency{
		CurrencyId:      param.CurrencyId,
		PId:             param.PId,
		Currency:        param.Currency,
		Symbol:          param.Symbol,
		Logo:            param.Logo,
		ContractAddress: param.ContractAddress,
		CurrencyType:    param.CurrencyType,
		Protocol:        param.Protocol,
		Decimals:        param.Decimals,
		Remark:          param.Remark,
		Status:          param.Status,
		UpdateBy:        security.GetAuthUserName(ctx),
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// Remove 删除币种
func (*CurrencyController) Remove(ctx *gin.Context) {

	currencyId, _ := strconv.Atoi(ctx.Param("currencyId"))

	if err := (&service.CurrencyService{}).DeleteCurrency(currencyId); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}
