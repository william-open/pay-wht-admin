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

type CollectionAddressController struct{}

// List 归集钱包列表
func (*CollectionAddressController) List(ctx *gin.Context) {

	var param dto.AddressListRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	list, total := (&service.AddressService{}).GetAddressList(param, true)
	response.NewSuccess().SetPageData(list, total).Json(ctx)
}

// Detail 归集钱包详情
func (*CollectionAddressController) Detail(ctx *gin.Context) {

	addressId, _ := strconv.Atoi(ctx.Param("addressId"))

	address := (&service.AddressService{}).GetAddressByAddressId(addressId)

	response.NewSuccess().SetData("data", address).Json(ctx)
}

// Create 新增归集钱包地址
func (*CollectionAddressController) Create(ctx *gin.Context) {

	var param dto.CreateCollectionAddressRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.CreateCollectionAddressValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if data := (&service.CollectionAddressService{}).GetCollectionAddressByAddress(param.Address); data.Id > 0 {
		response.NewError().SetMsg("新增归集钱包地址" + param.Address + "失败，归集钱包地址已存在").Json(ctx)
		return
	}

	if err := (&service.CollectionAddressService{}).CreateCollectionAddress(dto.SaveCollectionAddress{
		MId:          param.MId,
		ChainSymbol:  param.ChainSymbol,
		Symbol:       param.Symbol,
		Protocol:     param.Protocol,
		CurrencyType: param.CurrencyType,
		Address:      param.Address,
		CreateBy:     security.GetAuthUserName(ctx),
		Remark:       param.Remark,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// Update 更新归集钱包地址
func (*CollectionAddressController) Update(ctx *gin.Context) {

	var param dto.UpdateCollectionAddressRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.UpdateCollectionAddressValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}
	merchant := (&service.CollectionAddressService{}).GetCollectionAddressByAddress(param.Address)
	if merchant.MId > 0 && merchant.MId != param.MId {
		response.NewError().SetMsg("修改归集钱包地址" + param.Address + "失败，钱包地址已存在").Json(ctx)
		return
	}

	if err := (&service.CollectionAddressService{}).UpdateCollectionAddress(dto.SaveCollectionAddress{
		MId:          param.MId,
		ChainSymbol:  param.ChainSymbol,
		Symbol:       param.Symbol,
		Protocol:     param.Protocol,
		CurrencyType: param.CurrencyType,
		Address:      param.Address,
		Remark:       param.Remark,
		Status:       param.Status,
		UpdateBy:     security.GetAuthUserName(ctx),
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// Remove 删除归集钱包地址
func (*CollectionAddressController) Remove(ctx *gin.Context) {

	menuId, _ := strconv.Atoi(ctx.Param("id"))

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
