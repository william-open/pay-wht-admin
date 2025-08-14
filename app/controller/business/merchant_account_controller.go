package businesscontroller

import (
	"github.com/gin-gonic/gin"
	"ruoyi-go/app/dto"
	"ruoyi-go/app/service"
	"ruoyi-go/framework/response"
	"strconv"
)

type MerchantAccountController struct{}

// List 商户账户列表
func (*MerchantAccountController) List(ctx *gin.Context) {

	var param dto.MerchantAccountRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	list, total := (&service.MerchantAccountService{}).GetMerchantAccountList(param, true)
	response.NewSuccess().SetPageData(list, total).Json(ctx)
}

// Detail 商户账户详情
func (*MerchantAccountController) Detail(ctx *gin.Context) {

	merchantId, _ := strconv.Atoi(ctx.Param("merchantId"))

	merchant := (&service.MerchantAccountService{}).GetAccountByMerchantId(merchantId)

	response.NewSuccess().SetData("data", merchant).Json(ctx)
}

// CurrencyList 商户货币明显列表
func (*MerchantAccountController) CurrencyList(ctx *gin.Context) {

	var param dto.MerchantAccountCurrencyRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	list, total := (&service.MerchantAccountService{}).GetMerchantAccountCurrencyList(param, true)
	response.NewSuccess().SetPageData(list, total).Json(ctx)
}
