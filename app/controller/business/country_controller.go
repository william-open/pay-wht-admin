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

type CountryController struct{}

// 国家列表
func (*CountryController) List(ctx *gin.Context) {

	var param dto.CountryListRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	countrys, total := (&service.CountryService{}).GetCountryList(param, true)

	response.NewSuccess().SetPageData(countrys, total).Json(ctx)
}

// 国家详情
func (*CountryController) Detail(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	country := (&service.CountryService{}).GetCountryById(id)

	response.NewSuccess().SetData("data", country).Json(ctx)
}

// 新增国家
func (*CountryController) Create(ctx *gin.Context) {

	var param dto.CreateCountryRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.CreateCountryValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if country := (&service.CountryService{}).GetCountryByCode(param.Code); country.Id > 0 {
		response.NewError().SetMsg("新增国家" + param.Code + "失败，币种代码已存在").Json(ctx)
		return
	}

	if country := (&service.CountryService{}).GetCountryBySymbol(param.Symbol); country.Id > 0 {
		response.NewError().SetMsg("新增国家" + param.Symbol + "失败，货币符号已存在").Json(ctx)
		return
	}

	isOpen := 0
	if param.IsOpen == "" && param.IsOpen != string(0) {
		isOpen = 1
	}

	if err := (&service.CountryService{}).CreateCountry(dto.SaveCountry{
		Code:    param.Code,
		NameZh:  param.NameZh,
		NameEn:  param.NameEn,
		IsOpen:  &isOpen,
		Symbol:  param.Symbol,
		Country: param.Country,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 更新国家
func (*CountryController) Update(ctx *gin.Context) {

	var param dto.UpdateCountryRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.UpdateCountryValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if country := (&service.CountryService{}).GetCountryByCode(param.Code); country.Id > 0 && country.Id != param.Id {
		response.NewError().SetMsg("修改国家" + param.Code + "失败，货币代码已存在").Json(ctx)
		return
	}

	if country := (&service.CountryService{}).GetCountryBySymbol(param.Symbol); country.Id > 0 && country.Id != param.Id {
		response.NewError().SetMsg("修改国家" + param.Symbol + "失败，货币符号已存在").Json(ctx)
		return
	}

	isOpen := 0
	if param.IsOpen == "" && param.IsOpen != string(0) {
		isOpen = 1
	}

	if err := (&service.CountryService{}).UpdateCountry(dto.SaveCountry{
		Id:      param.Id,
		Code:    param.Code,
		NameZh:  param.NameZh,
		NameEn:  param.NameEn,
		IsOpen:  &isOpen,
		Symbol:  param.Symbol,
		Country: param.Country,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 修改国家状态
func (*CountryController) ChangeStatus(ctx *gin.Context) {

	var param dto.UpdateCountryRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.ChangeCountryStatusValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}
	isOpen, err := strconv.Atoi(param.IsOpen)
	if err != nil {
		fmt.Println("转换失败:", err)
		return
	}
	if err := (&service.CountryService{}).UpdateCountry(dto.SaveCountry{
		Id:     param.Id,
		IsOpen: &isOpen,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 根据状态查询所有国家
func (*CountryController) GetListByStatus(ctx *gin.Context) {

	var param dto.QueryCountryByStatusRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	status, err := strconv.Atoi(param.Status)
	if err != nil {
		fmt.Println("类型转化错误:", err)
		return
	}

	countrys := (&service.CountryService{}).GetAllCountryList(status)

	response.NewSuccess().SetData("data", countrys).Json(ctx)
}
