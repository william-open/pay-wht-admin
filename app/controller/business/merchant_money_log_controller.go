package businesscontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"strconv"
	"strings"
	"wht-admin/app/dto"
	"wht-admin/app/security"
	"wht-admin/app/service"
	"wht-admin/app/validator"
	"wht-admin/common/key"
	"wht-admin/common/password"
	"wht-admin/common/utils"
	"wht-admin/framework/response"
)

type MerchantMoneyLogController struct{}

// List 商户资金日志列表
func (*MerchantMoneyLogController) List(ctx *gin.Context) {

	var param dto.MoneyLogListRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	list, total := (&service.MoneyLogService{}).GetMoneyLogList(param, true)
	response.NewSuccess().SetPageData(list, total).Json(ctx)
}

// 商户资金日志详情
func (*MerchantMoneyLogController) Detail(ctx *gin.Context) {

	AgentId, _ := strconv.Atoi(ctx.Param("AgentId"))

	Agent := (&service.AgentService{}).GetAgentByAgentId(AgentId)

	response.NewSuccess().SetData("data", Agent).Json(ctx)
}

// Create 新增商户资金日志
func (*MerchantMoneyLogController) Create(ctx *gin.Context) {

	var param dto.CreateAgentRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.CreateAgentValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if Agent := (&service.AgentService{}).GetAgentByAgentName(param.Username); Agent.MId > 0 {
		response.NewError().SetMsg("新增商户资金日志" + param.Username + "失败，商户资金日志名称已存在").Json(ctx)
		return
	}
	u, _ := uuid.NewUUID()
	aesSecretKey, err := key.GenerateAESKey()
	if err != nil {
		response.NewError().SetMsg("新增商户资金日志失败，AES创建失败").Json(ctx)
		return
	}
	privatePEM, publicPEM, err := key.GenerateRSAKeys()
	if err != nil {
		response.NewError().SetMsg("新增商户资金日志失败,RSA创建失败").Json(ctx)
		return
	}
	emptyJson := datatypes.JSON([]byte("{}"))
	param.UpstreamId = string(emptyJson)
	param.Ways = string(emptyJson)
	AgentId, err := (&service.AgentService{}).CreateAgent(dto.SaveAgent{
		Username:          param.Username,
		Password:          password.Generate(param.Password),
		Nickname:          param.Nickname,
		CallbackSecretKey: param.CallbackSecretKey,
		NotifyUrl:         param.NotifyUrl,
		AesSecretKey:      aesSecretKey,
		PublicKey:         publicPEM,
		PrivateKey:        privatePEM,
		ApiKey:            utils.GenerateApiKey(),
		AppId:             GenerateUniqueAgentCode("80"),
		Status:            param.Status,
		CreateBy:          security.GetAuthUserName(ctx),
		PayType:           param.PayType,
		UserType:          param.UserType,
		Remark:            param.Remark,
		UpstreamId:        param.UpstreamId,
		Ways:              param.Ways,
	})
	if err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}
	// 创建商户资金日志后台管理员账号
	var salt = utils.RandomString(5)
	var shopPwd = utils.MakeMd5(strings.Trim(param.Password, " ") + salt)
	if err := (&service.ShopAdminService{}).CreateShopAdmin(dto.SaveShopAdmin{
		Username: param.Username,
		Nickname: param.Nickname,
		AppId:    u.String(),
		Role:     "1",
		MId:      uint(AgentId),
		Password: shopPwd,
		Salt:     salt,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}
	// 创建商户资金日志后台默认部门
	if err := (&service.ShopAdminService{}).CreateShopDept(dto.SaveShopDept{
		Name: "商户资金日志默认部门",
		Duty: param.Nickname,
		MId:  uint(AgentId),
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}
	response.NewSuccess().Json(ctx)
}

// Update 更新商户资金日志
func (*MerchantMoneyLogController) Update(ctx *gin.Context) {

	var param dto.UpdateAgentRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.UpdateAgentValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}
	Agent := (&service.AgentService{}).GetAgentByAgentName(param.Username)
	if Agent.MId > 0 && Agent.MId != param.MId {
		response.NewError().SetMsg("修改商户资金日志账号" + param.Username + "失败，商户资金日志账号已存在").Json(ctx)
		return
	}
	var editPassword = ""
	if param.Password == "" {
		editPassword = Agent.Password
	} else {
		editPassword = password.Generate(param.Password)
	}

	if err := (&service.AgentService{}).UpdateAgent(dto.SaveAgent{
		MId:               param.MId,
		Username:          param.Username,
		Password:          editPassword,
		Nickname:          param.Nickname,
		CallbackSecretKey: param.CallbackSecretKey,
		NotifyUrl:         param.NotifyUrl,
		Remark:            param.Remark,
		Status:            param.Status,
		UpdateBy:          security.GetAuthUserName(ctx),
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// Remove 删除商户资金日志
func (*MerchantMoneyLogController) Remove(ctx *gin.Context) {

	menuId, _ := strconv.Atoi(ctx.Param("menuId"))

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
