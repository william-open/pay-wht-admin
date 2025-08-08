package businesscontroller

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/service"
	"ruoyi-go/framework/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AddressController struct{}

// List 币种钱包地址列表
func (*AddressController) List(ctx *gin.Context) {

	var param dto.AddressListRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	list, total := (&service.AddressService{}).GetAddressList(param, true)
	response.NewSuccess().SetPageData(list, total).Json(ctx)
}

// Detail 币种钱包地址详情
func (*AddressController) Detail(ctx *gin.Context) {

	addressId, _ := strconv.Atoi(ctx.Param("addressId"))

	address := (&service.AddressService{}).GetAddressByAddressId(addressId)

	response.NewSuccess().SetData("data", address).Json(ctx)
}
