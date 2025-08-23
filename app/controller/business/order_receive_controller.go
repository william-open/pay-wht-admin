package businesscontroller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"wht-admin/app/dto"
	"wht-admin/app/service"
	"wht-admin/framework/response"
)

type OrderReceiveController struct{}

// 代收列表
func (*OrderReceiveController) List(ctx *gin.Context) {

	var param dto.OrderReceiveListRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	channels, total := (&service.OrderReceiveService{}).GetOrderReceiveList(param, true)

	response.NewSuccess().SetPageData(channels, total).Json(ctx)
}

// 代收详情
func (*OrderReceiveController) Detail(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("orderId"))

	yearMonth := ctx.Param("yearMonth")
	if yearMonth == "" {
		// 如果年月为空，尝试从ID中推断或者使用默认逻辑
		// 这里假设可以根据ID范围确定分表，或者需要其他逻辑
		// 作为示例，我们使用当前月份
		yearMonth = time.Now().Format("200601")
	}
	detail, _ := (&service.OrderReceiveService{}).GetOrderReceiveById(uint(id), yearMonth)

	response.NewSuccess().SetData("data", detail).Json(ctx)
}
