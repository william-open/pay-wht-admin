package businesscontroller

import (
	"github.com/gin-gonic/gin"
	"ruoyi-go/app/dto"
	"ruoyi-go/app/service"
	"ruoyi-go/app/validator"
	"ruoyi-go/framework/response"
	"strconv"
)

type ChannelController struct{}

// 通道列表
func (*ChannelController) List(ctx *gin.Context) {

	var param dto.ChannelListRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	channels, total := (&service.ChannelService{}).GetChannelList(param, true)

	response.NewSuccess().SetPageData(channels, total).Json(ctx)
}

// 通道详情
func (*ChannelController) Detail(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	channel := (&service.ChannelService{}).GetChannelById(id)

	response.NewSuccess().SetData("data", channel).Json(ctx)
}

// 新增通道
func (*ChannelController) Create(ctx *gin.Context) {

	var param dto.CreateChannelRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.CreateChannelValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if country := (&service.ChannelService{}).GetChannelByTitle(param.Title); country.Id > 0 {
		response.NewError().SetMsg("新增通道" + param.Title + "失败，通道名称码已存在").Json(ctx)
		return
	}

	if country := (&service.ChannelService{}).GetChannelByCoding(param.Coding); country.Id > 0 {
		response.NewError().SetMsg("新增通道" + param.Coding + "失败，通道编码已存在").Json(ctx)
		return
	}

	if err := (&service.ChannelService{}).CreateChannel(dto.SaveChannel{
		Title:       param.Title,
		Coding:      param.Coding,
		Max:         param.Max,
		Min:         param.Min,
		DefaultRate: param.DefaultRate,
		AddRate:     param.AddRate,
		Status:      param.Status,
		Type:        param.Type,
		Currency:    param.Currency,
		Remark:      param.Remark,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 更新通道
func (*ChannelController) Update(ctx *gin.Context) {

	var param dto.UpdateChannelRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.UpdateChannelValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if channel := (&service.ChannelService{}).GetChannelByTitle(param.Title); channel.Id > 0 && channel.Id != param.Id {
		response.NewError().SetMsg("修改通道" + param.Title + "失败，通道名称已存在").Json(ctx)
		return
	}

	if channel := (&service.ChannelService{}).GetChannelByCoding(param.Coding); channel.Id > 0 && channel.Id != param.Id {
		response.NewError().SetMsg("修改通道" + param.Coding + "失败，通道编码已存在").Json(ctx)
		return
	}

	if err := (&service.ChannelService{}).UpdateChannel(dto.SaveChannel{
		Id:          param.Id,
		Title:       param.Title,
		Coding:      param.Coding,
		Status:      param.Status,
		Max:         param.Max,
		Min:         param.Min,
		DefaultRate: param.DefaultRate,
		AddRate:     param.AddRate,
		Type:        param.Type,
		Currency:    param.Currency,
		Remark:      param.Remark,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 修改通道状态
func (*ChannelController) ChangeStatus(ctx *gin.Context) {

	var param dto.UpdateChannelRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.ChangeChannelStatusValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}
	if err := (&service.ChannelService{}).UpdateChannel(dto.SaveChannel{
		Id:     param.Id,
		Status: param.Status,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 根据状态查询下拉列表数据
func (*ChannelController) GetListByStatus(ctx *gin.Context) {
	var param dto.QueryChannelByStatusRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	channelList := (&service.ChannelService{}).GetAllDropDownList(param)

	response.NewSuccess().SetData("data", channelList).Json(ctx)
}
