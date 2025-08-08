package businesscontroller

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"ruoyi-go/app/dto"
	"ruoyi-go/app/security"
	"ruoyi-go/app/service"
	"ruoyi-go/app/validator"
	"ruoyi-go/common/password"
	"ruoyi-go/common/utils"
	"ruoyi-go/framework/response"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"ruoyi-go/common/key"
)

type MerchantController struct{}

// List 商户列表
func (*MerchantController) List(ctx *gin.Context) {

	var param dto.MerchantListRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	list, total := (&service.MerchantService{}).GetMerchantList(param, true)
	response.NewSuccess().SetPageData(list, total).Json(ctx)
}

// 商户详情
func (*MerchantController) Detail(ctx *gin.Context) {

	merchantId, _ := strconv.Atoi(ctx.Param("merchantId"))

	merchant := (&service.MerchantService{}).GetMerchantByMerchantId(merchantId)

	response.NewSuccess().SetData("data", merchant).Json(ctx)
}

// Create 新增商户
func (*MerchantController) Create(ctx *gin.Context) {

	var param dto.CreateMerchantRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.CreateMerchantValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if merchant := (&service.MerchantService{}).GetMerchantByMerchantName(param.Username); merchant.MId > 0 {
		response.NewError().SetMsg("新增商户" + param.Username + "失败，商户名称已存在").Json(ctx)
		return
	}
	u, _ := uuid.NewUUID()
	aesSecretKey, err := key.GenerateAESKey()
	if err != nil {
		response.NewError().SetMsg("新增商户失败，AES创建失败").Json(ctx)
		return
	}
	privatePEM, publicPEM, err := key.GenerateRSAKeys()
	if err != nil {
		response.NewError().SetMsg("新增商户失败,RSA创建失败").Json(ctx)
		return
	}

	merchantId, err := (&service.MerchantService{}).CreateMerchant(dto.SaveMerchant{
		Username:          param.Username,
		Password:          password.Generate(param.Password),
		Nickname:          param.Nickname,
		CallbackSecretKey: param.CallbackSecretKey,
		NotifyUrl:         param.NotifyUrl,
		AesSecretKey:      aesSecretKey,
		PublicKey:         publicPEM,
		PrivateKey:        privatePEM,
		ApiKey:            utils.GenerateApiKey(),
		AppId:             GenerateUniqueMerchantCode("80"),
		Status:            param.Status,
		CreateBy:          security.GetAuthUserName(ctx),
		Remark:            param.Remark,
	})
	if err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}
	// 创建商户后台管理员账号
	var salt = utils.RandomString(5)
	var shopPwd = utils.MakeMd5(strings.Trim(param.Password, " ") + salt)
	if err := (&service.ShopAdminService{}).CreateShopAdmin(dto.SaveShopAdmin{
		Username: param.Username,
		Nickname: param.Nickname,
		AppId:    u.String(),
		Role:     "1",
		MId:      uint(merchantId),
		Password: shopPwd,
		Salt:     salt,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}
	// 创建商户后台默认部门
	if err := (&service.ShopAdminService{}).CreateShopDept(dto.SaveShopDept{
		Name: "商户默认部门",
		Duty: param.Nickname,
		MId:  uint(merchantId),
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}
	response.NewSuccess().Json(ctx)
}

// Update 更新商户
func (*MerchantController) Update(ctx *gin.Context) {

	var param dto.UpdateMerchantRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.UpdateMerchantValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}
	merchant := (&service.MerchantService{}).GetMerchantByMerchantName(param.Username)
	if merchant.MId > 0 && merchant.MId != param.MId {
		response.NewError().SetMsg("修改商户账号" + param.Username + "失败，商户账号已存在").Json(ctx)
		return
	}
	var editPassword = ""
	if param.Password == "" {
		editPassword = merchant.Password
	} else {
		editPassword = password.Generate(param.Password)
	}

	if err := (&service.MerchantService{}).UpdateMerchant(dto.SaveMerchant{
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

// Remove 删除商户
func (*MerchantController) Remove(ctx *gin.Context) {

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

func GenerateUniqueMerchantCode(prefix string) string {
	rand.Seed(time.Now().UnixNano())

	for {
		suffix := rand.Intn(1000000) // 6位随机数
		if suffix%100000 == 0 {      // ❌ 避免结尾为000000
			continue
		}

		code := fmt.Sprintf("%s%06d", prefix, suffix)

		// ✅ 唯一性校验
		count := (&service.MerchantService{}).GetMerchantByAppId(code)
		if count == 0 {
			return code
		}
	}
}
