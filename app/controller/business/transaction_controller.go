package businesscontroller

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/service"
	"ruoyi-go/framework/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionController struct{}

// List 交易列表
func (*TransactionController) List(ctx *gin.Context) {

	var param dto.TransactionListRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	list, total := (&service.TransactionService{}).GetTransactionList(param, true)
	response.NewSuccess().SetPageData(list, total).Json(ctx)
}

// Detail 交易详情
func (*TransactionController) Detail(ctx *gin.Context) {

	transactionId, _ := strconv.Atoi(ctx.Param("id"))

	transaction := (&service.TransactionService{}).GetTransactionByTransactionId(transactionId)

	response.NewSuccess().SetData("data", transaction).Json(ctx)
}
