package businesscontroller

import (
	"strconv"
	"wht-admin/app/dto"
	"wht-admin/app/service"
	"wht-admin/framework/response"

	"github.com/gin-gonic/gin"
)

type CollectionController struct{}

// List 归集列表
func (*CollectionController) List(ctx *gin.Context) {

	var param dto.CollectionListRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	list, total := (&service.CollectionService{}).GetCollectionList(param, true)
	response.NewSuccess().SetPageData(list, total).Json(ctx)
}

// Detail 归集详情
func (*CollectionController) Detail(ctx *gin.Context) {

	collectionId, _ := strconv.Atoi(ctx.Param("id"))

	collection := (&service.CollectionService{}).GetCollectionByCollectionId(collectionId)

	response.NewSuccess().SetData("data", collection).Json(ctx)
}
