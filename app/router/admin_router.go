package router

import (
	"wht-admin/app/controller"
	businesscontroller "wht-admin/app/controller/business"
	monitorcontroller "wht-admin/app/controller/monitor"
	systemcontroller "wht-admin/app/controller/system"
	"wht-admin/app/middleware"
	"wht-admin/common/types/constant"

	"github.com/gin-gonic/gin"
)

// 后台路由组
func RegisterAdminGroupApi(api *gin.RouterGroup) {

	api.Use(middleware.Cors())                                                                  // 跨域中间件
	api.GET("/captchaImage", (&controller.AuthController{}).CaptchaImage)                       // 获取验证码
	api.POST("/register", (&controller.AuthController{}).Register)                              // 注册
	api.POST("/login", middleware.LogininforMiddleware(), (&controller.AuthController{}).Login) // 登录
	api.POST("/logout", (&controller.AuthController{}).Logout)                                  // 退出登录

	// 启用鉴权中间件，下面的路由都需要鉴权中间件验证通过后才可访问
	api.Use(middleware.AuthMiddleware())

	api.GET("/getInfo", (&controller.AuthController{}).GetInfo)       // 获取用户信息
	api.GET("/getRouters", (&controller.AuthController{}).GetRouters) // 获取路由信息

	api.GET("/system/user/profile", (&systemcontroller.UserController{}).GetProfile)                      // 个人信息
	api.PUT("/system/user/profile", (&systemcontroller.UserController{}).UpdateProfile)                   // 修改用户
	api.PUT("/system/user/profile/updatePwd", (&systemcontroller.UserController{}).UserProfileUpdatePwd)  // 重置密码
	api.POST("/system/user/profile/avatar", (&systemcontroller.UserController{}).UserProfileUpdateAvatar) // 更新头像

	api.GET("/system/user/deptTree", middleware.HasPerm("system:user:list"), (&systemcontroller.UserController{}).DeptTree)          // 获取部门树列表
	api.GET("/system/user/list", middleware.HasPerm("system:user:list"), (&systemcontroller.UserController{}).List)                  // 获取用户列表
	api.GET("/system/user/", middleware.HasPerm("system:user:query"), (&systemcontroller.UserController{}).Detail)                   // 根据用户编号获取详细信息
	api.GET("/system/user/:userId", middleware.HasPerm("system:user:query"), (&systemcontroller.UserController{}).Detail)            // 根据用户编号获取详细信息
	api.GET("/system/user/authRole/:userId", middleware.HasPerm("system:user:query"), (&systemcontroller.UserController{}).AuthRole) // 根据用户编号获取详细信息

	api.POST("/system/user", middleware.HasPerm("system:user:add"), middleware.OperLogMiddleware("新增用户", constant.REQUEST_BUSINESS_TYPE_INSERT), (&systemcontroller.UserController{}).Create)
	api.PUT("/system/user", middleware.HasPerm("system:user:edit"), middleware.OperLogMiddleware("更新用户", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.UserController{}).Update)
	api.DELETE("/system/user/:userIds", middleware.HasPerm("system:user:remove"), middleware.OperLogMiddleware("删除用户", constant.REQUEST_BUSINESS_TYPE_DELETE), (&systemcontroller.UserController{}).Remove)
	api.PUT("/system/user/changeStatus", middleware.HasPerm("system:user:edit"), middleware.OperLogMiddleware("修改用户状态", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.UserController{}).ChangeStatus)
	api.PUT("/system/user/resetPwd", middleware.HasPerm("system:user:edit"), middleware.OperLogMiddleware("修改用户密码", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.UserController{}).ResetPwd)
	api.PUT("/system/user/authRole", middleware.HasPerm("system:user:edit"), middleware.OperLogMiddleware("用户授权角色", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.UserController{}).AddAuthRole)
	api.POST("/system/user/export", middleware.HasPerm("system:user:export"), middleware.OperLogMiddleware("导出用户", constant.REQUEST_BUSINESS_TYPE_EXPORT), (&systemcontroller.UserController{}).Export)
	api.POST("/system/user/importData", middleware.HasPerm("system:user:import"), middleware.OperLogMiddleware("导入用户", constant.REQUEST_BUSINESS_TYPE_IMPORT), (&systemcontroller.UserController{}).ImportData)
	api.POST("/system/user/importTemplate", middleware.OperLogMiddleware("导入用户模板", constant.REQUEST_BUSINESS_TYPE_IMPORT), (&systemcontroller.UserController{}).ImportTemplate)

	api.GET("/system/role/list", middleware.HasPerm("system:role:list"), (&systemcontroller.RoleController{}).List)                                            // 获取角色列表
	api.GET("/system/role/:roleId", middleware.HasPerm("system:role:query"), (&systemcontroller.RoleController{}).Detail)                                      // 获取角色详情
	api.GET("/system/role/deptTree/:roleId", middleware.HasPerm("system:role:query"), (&systemcontroller.RoleController{}).DeptTree)                           // 获取部门树
	api.GET("/system/role/authUser/allocatedList", middleware.HasPerm("system:role:list"), (&systemcontroller.RoleController{}).RoleAuthUserAllocatedList)     // 查询已分配用户角色列表
	api.GET("/system/role/authUser/unallocatedList", middleware.HasPerm("system:role:list"), (&systemcontroller.RoleController{}).RoleAuthUserUnallocatedList) // 查询未分配用户角色列表

	api.POST("/system/role", middleware.HasPerm("system:role:add"), middleware.OperLogMiddleware("新增角色", constant.REQUEST_BUSINESS_TYPE_INSERT), (&systemcontroller.RoleController{}).Create)
	api.PUT("/system/role", middleware.HasPerm("system:role:edit"), middleware.OperLogMiddleware("更新角色", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.RoleController{}).Update)
	api.DELETE("/system/role/:roleIds", middleware.HasPerm("system:role:remove"), middleware.OperLogMiddleware("删除角色", constant.REQUEST_BUSINESS_TYPE_DELETE), (&systemcontroller.RoleController{}).Remove)
	api.PUT("/system/role/changeStatus", middleware.HasPerm("system:role:edit"), middleware.OperLogMiddleware("修改角色状态", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.RoleController{}).ChangeStatus)
	api.PUT("/system/role/dataScope", middleware.HasPerm("system:role:edit"), middleware.OperLogMiddleware("分配数据权限", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.RoleController{}).DataScope)
	api.PUT("/system/role/authUser/selectAll", middleware.HasPerm("system:role:edit"), middleware.OperLogMiddleware("批量选择用户授权", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.RoleController{}).RoleAuthUserSelectAll)
	api.PUT("/system/role/authUser/cancel", middleware.HasPerm("system:role:edit"), middleware.OperLogMiddleware("取消授权用户", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.RoleController{}).RoleAuthUserCancel)
	api.PUT("/system/role/authUser/cancelAll", middleware.HasPerm("system:role:edit"), middleware.OperLogMiddleware("批量取消授权用户", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.RoleController{}).RoleAuthUserCancelAll)
	api.POST("/system/role/export", middleware.HasPerm("system:role:export"), middleware.OperLogMiddleware("导出角色", constant.REQUEST_BUSINESS_TYPE_EXPORT), (&systemcontroller.RoleController{}).Export)

	api.GET("/system/menu/list", middleware.HasPerm("system:menu:list"), (&systemcontroller.MenuController{}).List)       // 获取菜单列表
	api.GET("/system/menu/treeselect", (&systemcontroller.MenuController{}).Treeselect)                                   // 获取菜单下拉树列表
	api.GET("/system/menu/roleMenuTreeselect/:roleId", (&systemcontroller.MenuController{}).RoleMenuTreeselect)           // 加载对应角色菜单列表树
	api.GET("/system/menu/:menuId", middleware.HasPerm("system:menu:query"), (&systemcontroller.MenuController{}).Detail) // 获取菜单详情

	api.POST("/system/menu", middleware.HasPerm("system:menu:add"), middleware.OperLogMiddleware("新增菜单", constant.REQUEST_BUSINESS_TYPE_INSERT), (&systemcontroller.MenuController{}).Create)
	api.PUT("/system/menu", middleware.HasPerm("system:menu:edit"), middleware.OperLogMiddleware("更新菜单", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.MenuController{}).Update)
	api.DELETE("/system/menu/:menuId", middleware.HasPerm("system:menu:remove"), middleware.OperLogMiddleware("删除菜单", constant.REQUEST_BUSINESS_TYPE_DELETE), (&systemcontroller.MenuController{}).Remove)

	api.GET("/system/dept/list", middleware.HasPerm("system:dept:list"), (&systemcontroller.DeptController{}).List)                        // 获取部门列表
	api.GET("/system/dept/list/exclude/:deptId", middleware.HasPerm("system:dept:list"), (&systemcontroller.DeptController{}).ListExclude) // 查询部门列表（排除节点）
	api.GET("/system/dept/:deptId", middleware.HasPerm("system:dept:query"), (&systemcontroller.DeptController{}).Detail)                  // 获取部门详情

	api.POST("/system/dept", middleware.HasPerm("system:dept:add"), middleware.OperLogMiddleware("新增部门", constant.REQUEST_BUSINESS_TYPE_INSERT), (&systemcontroller.DeptController{}).Create)
	api.PUT("/system/dept", middleware.HasPerm("system:dept:edit"), middleware.OperLogMiddleware("更新部门", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.DeptController{}).Update)
	api.DELETE("/system/dept/:deptId", middleware.HasPerm("system:dept:remove"), middleware.OperLogMiddleware("删除部门", constant.REQUEST_BUSINESS_TYPE_DELETE), (&systemcontroller.DeptController{}).Remove)

	api.GET("/system/post/list", middleware.HasPerm("system:post:list"), (&systemcontroller.PostController{}).List)       // 获取岗位列表
	api.GET("/system/post/:postId", middleware.HasPerm("system:post:query"), (&systemcontroller.PostController{}).Detail) // 获取岗位详情

	api.POST("/system/post", middleware.HasPerm("system:post:add"), middleware.OperLogMiddleware("新增岗位", constant.REQUEST_BUSINESS_TYPE_INSERT), (&systemcontroller.PostController{}).Create)
	api.PUT("/system/post", middleware.HasPerm("system:post:edit"), middleware.OperLogMiddleware("更新岗位", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.PostController{}).Update)
	api.DELETE("/system/post/:postIds", middleware.HasPerm("system:post:remove"), middleware.OperLogMiddleware("删除岗位", constant.REQUEST_BUSINESS_TYPE_DELETE), (&systemcontroller.PostController{}).Remove)
	api.POST("/system/post/export", middleware.HasPerm("system:post:export"), middleware.OperLogMiddleware("导出岗位", constant.REQUEST_BUSINESS_TYPE_EXPORT), (&systemcontroller.PostController{}).Export)

	api.GET("/system/dict/list", middleware.HasPerm("system:dict:list"), (&systemcontroller.DictTypeController{}).List)            // 获取字典类型列表
	api.GET("/system/dict/type/:dictId", middleware.HasPerm("system:dict:query"), (&systemcontroller.DictTypeController{}).Detail) // 获取字典类型详情
	api.GET("/system/dict/type/optionselect", (&systemcontroller.DictTypeController{}).Optionselect)                               // 获取字典选择框列表

	api.POST("/system/dict/type", middleware.HasPerm("system:dict:add"), middleware.OperLogMiddleware("新增字典类型", constant.REQUEST_BUSINESS_TYPE_INSERT), (&systemcontroller.DictTypeController{}).Create)
	api.PUT("/system/dict/type", middleware.HasPerm("system:dict:edit"), middleware.OperLogMiddleware("更新字典类型", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.DictTypeController{}).Update)
	api.DELETE("/system/dict/type/:dictIds", middleware.HasPerm("system:dict:remove"), middleware.OperLogMiddleware("删除字典类型", constant.REQUEST_BUSINESS_TYPE_DELETE), (&systemcontroller.DictTypeController{}).Remove)
	api.POST("/system/dict/type/export", middleware.HasPerm("system:dict:export"), middleware.OperLogMiddleware("导出字典类型", constant.REQUEST_BUSINESS_TYPE_EXPORT), (&systemcontroller.DictTypeController{}).Export)
	api.DELETE("/system/dict/type/refreshCache", middleware.HasPerm("system:dict:remove"), middleware.OperLogMiddleware("刷新字典类型缓存", constant.REQUEST_BUSINESS_TYPE_DELETE), (&systemcontroller.DictTypeController{}).RefreshCache)

	api.GET("/system/dict/data/list", middleware.HasPerm("system:dict:list"), (&systemcontroller.DictDataController{}).List)         // 获取字典数据列表
	api.GET("/system/dict/data/:dictCode", middleware.HasPerm("system:dict:query"), (&systemcontroller.DictDataController{}).Detail) // 获取字典数据详情
	api.GET("/system/dict/data/type/:dictType", (&systemcontroller.DictDataController{}).Type)                                       // 根据字典类型查询字典数据

	api.POST("/system/dict/data", middleware.HasPerm("system:dict:add"), middleware.OperLogMiddleware("新增字典数据", constant.REQUEST_BUSINESS_TYPE_INSERT), (&systemcontroller.DictDataController{}).Create)
	api.PUT("/system/dict/data", middleware.HasPerm("system:dict:edit"), middleware.OperLogMiddleware("更新字典数据", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.DictDataController{}).Update)
	api.DELETE("/system/dict/data/:dictCodes", middleware.HasPerm("system:dict:remove"), middleware.OperLogMiddleware("删除字典数据", constant.REQUEST_BUSINESS_TYPE_DELETE), (&systemcontroller.DictDataController{}).Remove)
	api.POST("/system/dict/data/export", middleware.HasPerm("system:dict:export"), middleware.OperLogMiddleware("导出字典数据", constant.REQUEST_BUSINESS_TYPE_EXPORT), (&systemcontroller.DictDataController{}).Export)

	api.GET("/system/config/list", middleware.HasPerm("system:config:list"), (&systemcontroller.ConfigController{}).List)         // 获取参数配置列表
	api.GET("/system/config/:configId", middleware.HasPerm("system:config:query"), (&systemcontroller.ConfigController{}).Detail) // 获取参数配置详情
	api.GET("/system/config/configKey/:configKey", (&systemcontroller.ConfigController{}).ConfigKey)                              // 根据参数键名查询参数值

	api.POST("/system/config", middleware.HasPerm("system:config:add"), middleware.OperLogMiddleware("新增参数配置", constant.REQUEST_BUSINESS_TYPE_INSERT), (&systemcontroller.ConfigController{}).Create)
	api.PUT("/system/config", middleware.HasPerm("system:config:edit"), middleware.OperLogMiddleware("更新参数配置", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.ConfigController{}).Update)
	api.DELETE("/system/config/:configIds", middleware.HasPerm("system:config:remove"), middleware.OperLogMiddleware("删除参数配置", constant.REQUEST_BUSINESS_TYPE_DELETE), (&systemcontroller.ConfigController{}).Remove)
	api.POST("/system/config/export", middleware.HasPerm("system:config:export"), middleware.OperLogMiddleware("导出参数配置", constant.REQUEST_BUSINESS_TYPE_EXPORT), (&systemcontroller.ConfigController{}).Export)
	api.DELETE("/system/config/refreshCache", middleware.HasPerm("system:config:remove"), middleware.OperLogMiddleware("刷新参数配置缓存", constant.REQUEST_BUSINESS_TYPE_DELETE), (&systemcontroller.ConfigController{}).RefreshCache)

	api.GET("/monitor/logininfor/list", middleware.HasPerm("monitor:operlog:list"), (&monitorcontroller.LogininforController{}).List) // 获取登录日志列表

	api.DELETE("/monitor/logininfor/:infoIds", middleware.HasPerm("monitor:logininfor:remove"), middleware.OperLogMiddleware("删除登录日志", constant.REQUEST_BUSINESS_TYPE_DELETE), (&monitorcontroller.LogininforController{}).Remove)
	api.DELETE("/monitor/logininfor/clean", middleware.HasPerm("monitor:logininfor:remove"), middleware.OperLogMiddleware("清空登录日志", constant.REQUEST_BUSINESS_TYPE_DELETE), (&monitorcontroller.LogininforController{}).Clean)
	api.GET("/monitor/logininfor/unlock/:userName", middleware.HasPerm("monitor:logininfor:unlock"), middleware.OperLogMiddleware("账户解锁", constant.REQUEST_BUSINESS_TYPE_DELETE), (&monitorcontroller.LogininforController{}).Unlock)
	api.POST("/monitor/logininfor/export", middleware.HasPerm("monitor:logininfor:export"), middleware.OperLogMiddleware("导出登录日志", constant.REQUEST_BUSINESS_TYPE_EXPORT), (&monitorcontroller.LogininforController{}).Export)

	api.GET("/monitor/operlog/list", middleware.HasPerm("monitor:logininfor:list"), (&monitorcontroller.OperlogController{}).List) // 获取操作日志列表

	api.DELETE("/monitor/operlog/:operIds", middleware.HasPerm("monitor:operlog:remove"), middleware.OperLogMiddleware("删除操作日志", constant.REQUEST_BUSINESS_TYPE_DELETE), (&monitorcontroller.OperlogController{}).Remove)
	api.DELETE("/monitor/operlog/clean", middleware.HasPerm("monitor:operlog:remove"), middleware.OperLogMiddleware("清空操作日志", constant.REQUEST_BUSINESS_TYPE_DELETE), (&monitorcontroller.OperlogController{}).Clean)
	api.POST("/monitor/operlog/export", middleware.HasPerm("monitor:operlog:export"), middleware.OperLogMiddleware("导出操作日志", constant.REQUEST_BUSINESS_TYPE_EXPORT), (&monitorcontroller.OperlogController{}).Export)

	// 商户功能路由
	api.GET("/business/merchant/dropDownList", middleware.HasPerm("business:merchant:dropDownList"), (&businesscontroller.MerchantController{}).DropDownList) // 获取下拉列表商户
	api.GET("/business/merchant/list", middleware.HasPerm("business:merchant:list"), (&businesscontroller.MerchantController{}).List)                         // 获取商户列表
	api.GET("/business/merchant/:merchantId", middleware.HasPerm("business:merchant:query"), (&businesscontroller.MerchantController{}).Detail)               // 获取商户详情
	api.POST("/business/merchant", middleware.HasPerm("business:merchant:add"), middleware.OperLogMiddleware("新增商户", constant.REQUEST_BUSINESS_TYPE_INSERT), (&businesscontroller.MerchantController{}).Create)
	api.PUT("/business/merchant", middleware.HasPerm("business:merchant:edit"), middleware.OperLogMiddleware("更新商户", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&businesscontroller.MerchantController{}).Update)
	api.GET("/merchant/whitelist/list", middleware.HasPerm("merchant:whitelist:list"), (&businesscontroller.MerchantWhitelistController{}).List) // 获取商户白名单列表
	api.DELETE("/merchant/whitelist/:ids", middleware.HasPerm("merchant:whitelist:remove"), middleware.OperLogMiddleware("删除商户白名单", constant.REQUEST_BUSINESS_TYPE_DELETE), (&businesscontroller.MerchantWhitelistController{}).Remove)
	api.POST("/merchant/whitelist", middleware.HasPerm("merchant:whitelist:add"), middleware.OperLogMiddleware("新增商户白名单", constant.REQUEST_BUSINESS_TYPE_INSERT), (&businesscontroller.MerchantWhitelistController{}).Create)
	//api.DELETE("/business/merchant/:merchantId", middleware.HasPerm("system:menu:remove"), middleware.OperLogMiddleware("删除菜单", constant.REQUEST_BUSINESS_TYPE_DELETE), (&systemcontroller.MenuController{}).Remove)

	// 商户账户功能路由
	api.GET("/business/merchant_account/list", middleware.HasPerm("business:merchant_account:list"), (&businesscontroller.MerchantAccountController{}).List)                           // 获取商户账户列表
	api.GET("/business/merchant_account_currency/list", middleware.HasPerm("business:merchant_account_currency:list"), (&businesscontroller.MerchantAccountController{}).CurrencyList) // 获取商户货币账户列表
	api.GET("/business/merchant_account/:merchantId", middleware.HasPerm("business:merchant_account:query"), (&businesscontroller.MerchantAccountController{}).Detail)                 // 获取商户账户详情

	// 代理功能路由
	api.GET("/business/agent/list", middleware.HasPerm("business:agent:list"), (&businesscontroller.AgentController{}).List)        // 获取代理列表
	api.GET("/business/agent/:agentId", middleware.HasPerm("business:agent:query"), (&businesscontroller.AgentController{}).Detail) // 获取代理详情
	api.POST("/business/agent", middleware.HasPerm("business:agent:add"), middleware.OperLogMiddleware("新增代理", constant.REQUEST_BUSINESS_TYPE_INSERT), (&businesscontroller.AgentController{}).Create)
	api.PUT("/business/agent", middleware.HasPerm("business:agent:edit"), middleware.OperLogMiddleware("更新代理", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&businesscontroller.AgentController{}).Update)
	api.PUT("/business/agent/whitelist", middleware.HasPerm("business:agent:whitelist"), middleware.OperLogMiddleware("更新白名单", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&businesscontroller.AgentController{}).Whitelist)
	api.PUT("/business/agent/changeStatus", middleware.HasPerm("business:agent:changeStatus"), middleware.OperLogMiddleware("修改代理状态", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&businesscontroller.AgentController{}).ChangeStatus)

	//api.DELETE("/business/merchant/:merchantId", middleware.HasPerm("system:menu:remove"), middleware.OperLogMiddleware("删除菜单", constant.REQUEST_BUSINESS_TYPE_DELETE), (&systemcontroller.MenuController{}).Remove)

	// 代理商户功能路由
	api.GET("/agent/merchant/list", middleware.HasPerm("agent:merchant:list"), (&businesscontroller.AgentMController{}).List)         // 获取代理商户列表
	api.GET("/agent/merchant/:agentMId", middleware.HasPerm("agent:merchant:query"), (&businesscontroller.AgentMController{}).Detail) // 获取代理商户详情
	api.POST("/agent/merchant", middleware.HasPerm("agent:merchant:add"), middleware.OperLogMiddleware("新增代理商户", constant.REQUEST_BUSINESS_TYPE_INSERT), (&businesscontroller.AgentMController{}).Create)
	api.PUT("/agent/merchant", middleware.HasPerm("agent:merchant:edit"), middleware.OperLogMiddleware("更新代理商户", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&businesscontroller.AgentMController{}).Update)
	api.DELETE("/agent/merchant/:agentMId", middleware.HasPerm("agent:merchant:remove"), middleware.OperLogMiddleware("删除代理商户", constant.REQUEST_BUSINESS_TYPE_DELETE), (&businesscontroller.AgentMController{}).Remove)

	// 币种功能路由
	api.GET("/business/currency/list", middleware.HasPerm("business:currency:list"), (&businesscontroller.CurrencyController{}).List)           // 获取币种列表
	api.GET("/business/currency/:currencyId", middleware.HasPerm("business:currency:query"), (&businesscontroller.CurrencyController{}).Detail) // 获取币种详情
	api.POST("/business/currency", middleware.HasPerm("business:currency:add"), middleware.OperLogMiddleware("新增币种", constant.REQUEST_BUSINESS_TYPE_INSERT), (&businesscontroller.CurrencyController{}).Create)
	api.PUT("/business/currency", middleware.HasPerm("business:currency:edit"), middleware.OperLogMiddleware("更新币种", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&businesscontroller.CurrencyController{}).Update)
	api.DELETE("/business/currency/:currencyId", middleware.HasPerm("business:currency:remove"), middleware.OperLogMiddleware("删除币种", constant.REQUEST_BUSINESS_TYPE_DELETE), (&businesscontroller.CurrencyController{}).Remove)

	// 币种钱包地址功能路由
	api.GET("/business/address/list", middleware.HasPerm("business:address:list"), (&businesscontroller.AddressController{}).List)          // 获取币种钱包地址列表
	api.GET("/business/address/:addressId", middleware.HasPerm("business:address:query"), (&businesscontroller.AddressController{}).Detail) // 获取币种钱包地址详情

	// 交易功能路由
	api.GET("/business/transaction/list", middleware.HasPerm("business:transaction:list"), (&businesscontroller.TransactionController{}).List)   // 获取交易列表
	api.GET("/business/transaction/:id", middleware.HasPerm("business:transaction:query"), (&businesscontroller.TransactionController{}).Detail) // 获取交易详情

	// 归集功能路由
	api.GET("/business/collection/list", middleware.HasPerm("business:collection:list"), (&businesscontroller.CollectionController{}).List)   // 获取归集列表
	api.GET("/business/collection/:id", middleware.HasPerm("business:collection:query"), (&businesscontroller.CollectionController{}).Detail) // 获取归集详情

	// 归集钱包地址功能路由
	api.GET("/business/collection_address/list", middleware.HasPerm("business:collection_address:list"), (&businesscontroller.CollectionAddressController{}).List)   // 获取归集钱包地址列表
	api.GET("/business/collection_address/:id", middleware.HasPerm("business:collection_address:query"), (&businesscontroller.CollectionAddressController{}).Detail) // 获取归集钱包地址详情
	api.POST("/business/collection_address", middleware.HasPerm("business:collection_address:add"), middleware.OperLogMiddleware("新增币种", constant.REQUEST_BUSINESS_TYPE_INSERT), (&businesscontroller.CollectionAddressController{}).Create)
	api.PUT("/business/collection_address", middleware.HasPerm("business:collection_address:edit"), middleware.OperLogMiddleware("更新币种", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&businesscontroller.CollectionAddressController{}).Update)
	api.DELETE("/business/collection_address/:id", middleware.HasPerm("business:collection_address:remove"), middleware.OperLogMiddleware("删除币种", constant.REQUEST_BUSINESS_TYPE_DELETE), (&businesscontroller.CollectionAddressController{}).Remove)

	// 国家功能路由
	api.GET("/business/country/list", middleware.HasPerm("business:country:list"), (&businesscontroller.CountryController{}).List)                                        // 获取国家列表
	api.GET("/business/country/getAllListByStatus", middleware.HasPerm("business:country:getAllListByStatus"), (&businesscontroller.CountryController{}).GetListByStatus) // 通过状态获取国家
	api.GET("/business/country/:id", middleware.HasPerm("business:country:query"), (&businesscontroller.CountryController{}).Detail)                                      // 获取国家详情
	api.POST("/business/country", middleware.HasPerm("business:country:add"), middleware.OperLogMiddleware("新增国家", constant.REQUEST_BUSINESS_TYPE_INSERT), (&businesscontroller.CountryController{}).Create)
	api.PUT("/business/country", middleware.HasPerm("business:country:edit"), middleware.OperLogMiddleware("更新国家", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&businesscontroller.CountryController{}).Update)
	api.PUT("/business/country/changeStatus", middleware.HasPerm("business:country:edit"), middleware.OperLogMiddleware("修改国家状态", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&businesscontroller.CountryController{}).ChangeStatus)

	// 通道功能路由
	api.GET("/business/channel/getAllListByStatus", middleware.HasPerm("business:channel:getAllListByStatus"), (&businesscontroller.ChannelController{}).GetListByStatus) // 通过状态获取系统通道
	api.GET("/business/channel/list", middleware.HasPerm("business:channel:list"), (&businesscontroller.ChannelController{}).List)                                        // 获取通道列表
	api.GET("/business/channel/:id", middleware.HasPerm("business:channel:query"), (&businesscontroller.ChannelController{}).Detail)                                      // 获取通道详情
	api.POST("/business/channel", middleware.HasPerm("business:channel:add"), middleware.OperLogMiddleware("新增通道", constant.REQUEST_BUSINESS_TYPE_INSERT), (&businesscontroller.ChannelController{}).Create)
	api.PUT("/business/channel", middleware.HasPerm("business:channel:edit"), middleware.OperLogMiddleware("更新通道", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&businesscontroller.ChannelController{}).Update)
	api.PUT("/business/channel/changeStatus", middleware.HasPerm("business:channel:changeStatus"), middleware.OperLogMiddleware("修改通道状态", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&businesscontroller.ChannelController{}).ChangeStatus)

	// 通道供应商功能路由
	api.GET("/business/upstream/channelList", middleware.HasPerm("business:upstream:channelList"), (&businesscontroller.PayUpstreamController{}).GetUpChannelList)              // 查询上游供应商通道
	api.GET("/business/upstream/getAllListByStatus", middleware.HasPerm("business:upstream:getAllListByStatus"), (&businesscontroller.PayUpstreamController{}).GetListByStatus) // 通过状态获取供应商
	api.GET("/business/upstream/list", middleware.HasPerm("business:upstream:list"), (&businesscontroller.PayUpstreamController{}).List)                                        // 获取通道供应商列表
	api.GET("/business/upstream/:id", middleware.HasPerm("business:upstream:query"), (&businesscontroller.PayUpstreamController{}).Detail)                                      // 获取通道供应商详情
	api.POST("/business/upstream", middleware.HasPerm("business:upstream:add"), middleware.OperLogMiddleware("新增通道供应商", constant.REQUEST_BUSINESS_TYPE_INSERT), (&businesscontroller.PayUpstreamController{}).Create)
	api.PUT("/business/upstream", middleware.HasPerm("business:upstream:edit"), middleware.OperLogMiddleware("更新通道供应商", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&businesscontroller.PayUpstreamController{}).Update)
	api.PUT("/business/upstream/changeStatus", middleware.HasPerm("business:upstream:edit"), middleware.OperLogMiddleware("修改通道供应商状态", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&businesscontroller.PayUpstreamController{}).ChangeStatus)

	// 商户通道功能路由
	api.GET("/merchant/channel/list", middleware.HasPerm("merchant:channel:list"), (&businesscontroller.MerchantChannelController{}).List)   // 获取商户通道列表
	api.GET("/merchant/channel/:id", middleware.HasPerm("merchant:channel:query"), (&businesscontroller.MerchantChannelController{}).Detail) // 获取商户通道详情
	api.POST("/merchant/channel", middleware.HasPerm("merchant:channel:add"), middleware.OperLogMiddleware("新增商户通道", constant.REQUEST_BUSINESS_TYPE_INSERT), (&businesscontroller.MerchantChannelController{}).Create)
	api.PUT("/merchant/channel", middleware.HasPerm("merchant:channel:edit"), middleware.OperLogMiddleware("更新商户通道", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&businesscontroller.MerchantChannelController{}).Update)
	api.PUT("/merchant/channel/changeStatus", middleware.HasPerm("merchant:channel:edit"), middleware.OperLogMiddleware("修改商户通道状态", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&businesscontroller.MerchantChannelController{}).ChangeStatus)

	// 上游供应商产品路由
	api.GET("/upstream/product/list", middleware.HasPerm("upstream:product:list"), (&businesscontroller.UpstreamProductController{}).List)                 // 获取上游供应商产品列表
	api.GET("/upstream/product/:id", middleware.HasPerm("upstream:product:query"), (&businesscontroller.UpstreamProductController{}).Detail)               // 获取上游供应商产品详情
	api.GET("/upstream/test_product/:id", middleware.HasPerm("upstream:product:test_query"), (&businesscontroller.UpstreamProductController{}).TestDetail) // 获取测试上游供应商产品详情信息
	api.POST("/upstream/product", middleware.HasPerm("upstream:product:add"), middleware.OperLogMiddleware("新增商户通道", constant.REQUEST_BUSINESS_TYPE_INSERT), (&businesscontroller.UpstreamProductController{}).Create)
	api.POST("/upstream/test_product/order", middleware.HasPerm("upstream:product:test_order"), middleware.OperLogMiddleware("测试上游供应商通道", constant.REQUEST_BUSINESS_TYPE_INSERT), (&businesscontroller.UpstreamProductController{}).TestOrderCreate)
	api.PUT("/upstream/product", middleware.HasPerm("upstream:product:edit"), middleware.OperLogMiddleware("更新商户通道", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&businesscontroller.UpstreamProductController{}).Update)
	api.PUT("/upstream/product/changeStatus", middleware.HasPerm("upstream:product:edit"), middleware.OperLogMiddleware("修改商户通道产品状态", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&businesscontroller.UpstreamProductController{}).ChangeStatus)

	// 代理资金日志功能路由
	api.GET("/agent/money_log/list", middleware.HasPerm("agent:money_log:list"), (&businesscontroller.AgentMoneyLogController{}).List)   // 获取代理资金日志列表
	api.GET("/agent/money_log/:id", middleware.HasPerm("agent:money_log:query"), (&businesscontroller.AgentMoneyLogController{}).Detail) // 获取代理资金日志详情

	// 商户资金日志功能路由
	api.GET("/merchant/money_log/list", middleware.HasPerm("merchant:money_log:list"), (&businesscontroller.AgentMoneyLogController{}).List)   // 获取代理资金日志列表
	api.GET("/merchant/money_log/:id", middleware.HasPerm("merchant:money_log:query"), (&businesscontroller.AgentMoneyLogController{}).Detail) // 获取代理资金日志详情

	// 代收功能路由
	api.GET("/business/order_receive/list", middleware.HasPerm("business:order_receive:list"), (&businesscontroller.OrderReceiveController{}).List)                   // 获取代收列表
	api.GET("/business/order_receive/:orderId/:yearMonth", middleware.HasPerm("business:order_receive:query"), (&businesscontroller.OrderReceiveController{}).Detail) // 获取代收详情

	// 代付功能路由
	api.GET("/business/order_payout/list", middleware.HasPerm("business:order_payout:list"), (&businesscontroller.OrderPayoutController{}).List)                   // 获取代付列表
	api.GET("/business/order_payout/:orderId/:yearMonth", middleware.HasPerm("business:order_payout:query"), (&businesscontroller.OrderPayoutController{}).Detail) // 获取代付详情

}
