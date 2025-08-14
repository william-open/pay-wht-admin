package businesscontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"math/rand"
	"ruoyi-go/app/dto"
	"ruoyi-go/app/security"
	"ruoyi-go/app/service"
	"ruoyi-go/app/validator"
	"ruoyi-go/common/key"
	"ruoyi-go/common/password"
	"ruoyi-go/common/utils"
	"ruoyi-go/framework/response"
	"strconv"
	"strings"
	"time"
)

type AgentController struct{}

// List 代理列表
func (*AgentController) List(ctx *gin.Context) {

	var param dto.AgentListRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}
	// 代理商户
	param.UserType = 2
	list, total := (&service.AgentService{}).GetAgentList(param, true)
	response.NewSuccess().SetPageData(list, total).Json(ctx)
}

// 代理详情
func (*AgentController) Detail(ctx *gin.Context) {

	AgentId, _ := strconv.Atoi(ctx.Param("AgentId"))

	Agent := (&service.AgentService{}).GetAgentByAgentId(AgentId)

	response.NewSuccess().SetData("data", Agent).Json(ctx)
}

// Create 新增代理
func (*AgentController) Create(ctx *gin.Context) {

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
		response.NewError().SetMsg("新增代理" + param.Username + "失败，代理名称已存在").Json(ctx)
		return
	}
	u, _ := uuid.NewUUID()
	aesSecretKey, err := key.GenerateAESKey()
	if err != nil {
		response.NewError().SetMsg("新增代理失败，AES创建失败").Json(ctx)
		return
	}
	privatePEM, publicPEM, err := key.GenerateRSAKeys()
	if err != nil {
		response.NewError().SetMsg("新增代理失败,RSA创建失败").Json(ctx)
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
	// 创建代理后台管理员账号
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
	// 创建代理后台默认部门
	if err := (&service.ShopAdminService{}).CreateShopDept(dto.SaveShopDept{
		Name: "代理默认部门",
		Duty: param.Nickname,
		MId:  uint(AgentId),
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}
	response.NewSuccess().Json(ctx)
}

// Update 更新代理
func (*AgentController) Update(ctx *gin.Context) {

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
		response.NewError().SetMsg("修改代理账号" + param.Username + "失败，代理账号已存在").Json(ctx)
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

// Remove 删除代理
func (*AgentController) Remove(ctx *gin.Context) {

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

// 更新白名单 Whitelist
func (*AgentController) Whitelist(ctx *gin.Context) {

	var param dto.UpdateWhitelistRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.UpdateAgentWhitelistValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := (&service.AgentService{}).UpdateAgentWhitelist(dto.SaveAgentWhitelist{
		MId:        param.MId,
		ApiIp:      param.ApiIp,
		LoginApiIp: param.LoginApiIp,
		ApiDomain:  param.ApiDomain,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

func GenerateUniqueAgentCode(prefix string) string {
	rand.Seed(time.Now().UnixNano())

	for {
		suffix := rand.Intn(1000000) // 6位随机数
		if suffix%100000 == 0 {      // ❌ 避免结尾为000000
			continue
		}

		code := fmt.Sprintf("%s%06d", prefix, suffix)

		// ✅ 唯一性校验
		count := (&service.AgentService{}).GetAgentByAppId(code)
		if count == 0 {
			return code
		}
	}
}
