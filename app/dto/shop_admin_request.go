package dto

// 保存用户
type SaveShopAdmin struct {
	ID            uint   `gorm:"primarykey;comment:'主键'"`
	DeptId        uint   `gorm:"not null;default:0;comment:'部门ID'"`
	PostId        uint   `gorm:"not null;default:0;comment:'岗位ID'"`
	MId           uint   `gorm:"comment:'商户ID'"`
	Username      string `gorm:"not null;default:'';comment:'用户账号''"`
	Nickname      string `gorm:"not null;default:'';comment:'用户昵称'"`
	Password      string `gorm:"not null;default:'';comment:'用户密码'"`
	AppId         string `gorm:"not null;default:'';comment:'应用ID'"`
	Avatar        string `gorm:"not null;default:'';comment:'用户头像'"`
	Role          string `gorm:"not null;default:'';comment:'角色主键'"`
	Salt          string `gorm:"not null;default:'';comment:'加密盐巴'"`
	Sort          uint16 `gorm:"not null;default:0;comment:'排序编号'"`
	IsMultipoint  uint8  `gorm:"not null;default:0;comment:'多端登录: 0=否, 1=是''"`
	IsDisable     uint8  `gorm:"not null;default:0;comment:'是否禁用: 0=否, 1=是'"`
	IsDelete      uint8  `gorm:"not null;default:0;comment:'是否删除: 0=否, 1=是'"`
	LastLoginIp   string `gorm:"not null;default:'';comment:'最后登录IP'"`
	LastLoginTime int64  `gorm:"not null;default:0;comment:'最后登录时间'"`
	CreateTime    int64  `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime    int64  `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
	DeleteTime    int64  `gorm:"not null;default:0;comment:'删除时间'"`
	PayPassword   string `gorm:"not null;default:'';comment:'用户支付密码'"`
	PaySalt       string `gorm:"not null;default:'';comment:'用户支付盐值'"`
	ApiKey        string `gorm:"not null;default:'';comment:'APIKEY'"`
}

// 更新用户信息
type UpdateShopAdmin struct {
	Password    string `json:"password"`
	Salt        string `json:"salt"`
	PayPassword string `json:"payPassword"`
	PaySalt     string `json:"paySalt"`
	MId         uint   `json:"mId"`
}

// 新增用户
type CreateShopAdminRequest struct {
	DeptId        uint   `gorm:"not null;default:0;comment:'部门ID'"`
	PostId        uint   `gorm:"not null;default:0;comment:'岗位ID'"`
	MId           uint   `gorm:"comment:'商户ID'"`
	Username      string `gorm:"not null;default:'';comment:'用户账号''"`
	Nickname      string `gorm:"not null;default:'';comment:'用户昵称'"`
	Password      string `gorm:"not null;default:'';comment:'用户密码'"`
	Avatar        string `gorm:"not null;default:'';comment:'用户头像'"`
	AppId         string `gorm:"not null;default:'';comment:'应用ID'"`
	Role          string `gorm:"not null;default:'';comment:'角色主键'"`
	Salt          string `gorm:"not null;default:'';comment:'加密盐巴'"`
	Sort          uint16 `gorm:"not null;default:0;comment:'排序编号'"`
	IsMultipoint  uint8  `gorm:"not null;default:0;comment:'多端登录: 0=否, 1=是''"`
	IsDisable     uint8  `gorm:"not null;default:0;comment:'是否禁用: 0=否, 1=是'"`
	IsDelete      uint8  `gorm:"not null;default:0;comment:'是否删除: 0=否, 1=是'"`
	LastLoginIp   string `gorm:"not null;default:'';comment:'最后登录IP'"`
	LastLoginTime int64  `gorm:"not null;default:0;comment:'最后登录时间'"`
	CreateTime    int64  `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime    int64  `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
	DeleteTime    int64  `gorm:"not null;default:0;comment:'删除时间'"`
}

// 更新用户
type UpdateShopAdminRequest struct {
	ID            uint   `gorm:"primarykey;comment:'主键'"`
	DeptId        uint   `gorm:"not null;default:0;comment:'部门ID'"`
	PostId        uint   `gorm:"not null;default:0;comment:'岗位ID'"`
	MId           uint   `gorm:"comment:'商户ID'"`
	Username      string `gorm:"not null;default:'';comment:'用户账号''"`
	AppId         string `gorm:"not null;default:'';comment:'应用ID'"`
	Nickname      string `gorm:"not null;default:'';comment:'用户昵称'"`
	Password      string `gorm:"not null;default:'';comment:'用户密码'"`
	Avatar        string `gorm:"not null;default:'';comment:'用户头像'"`
	Role          string `gorm:"not null;default:'';comment:'角色主键'"`
	Salt          string `gorm:"not null;default:'';comment:'加密盐巴'"`
	Sort          uint16 `gorm:"not null;default:0;comment:'排序编号'"`
	IsMultipoint  uint8  `gorm:"not null;default:0;comment:'多端登录: 0=否, 1=是''"`
	IsDisable     uint8  `gorm:"not null;default:0;comment:'是否禁用: 0=否, 1=是'"`
	IsDelete      uint8  `gorm:"not null;default:0;comment:'是否删除: 0=否, 1=是'"`
	LastLoginIp   string `gorm:"not null;default:'';comment:'最后登录IP'"`
	LastLoginTime int64  `gorm:"not null;default:0;comment:'最后登录时间'"`
	CreateTime    int64  `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime    int64  `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
	DeleteTime    int64  `gorm:"not null;default:0;comment:'删除时间'"`
}

// 更新个人密码
type ShopAdminProfileUpdatePwdRequest struct {
	OldPassword string `query:"oldPassword" form:"oldPassword"`
	NewPassword string `query:"newPassword" form:"newPassword"`
}

// 保存商户部门
type SaveShopDept struct {
	ID         uint   `gorm:"primarykey;comment:'主键'"`
	Pid        uint   `gorm:"not null;default:0;comment:'上级主键'"`
	MId        uint   `gorm:"not null;default:0;comment:'商户ID'"`
	Name       string `gorm:"not null;default:'';comment:'部门名称''"`
	Duty       string `gorm:"not null;default:'';comment:'负责人名'"`
	Mobile     string `gorm:"not null;default:'';comment:'联系电话'"`
	Sort       uint16 `gorm:"not null;default:0;comment:'排序编号'"`
	IsStop     uint8  `gorm:"not null;default:0;comment:'是否停用: 0=否, 1=是'"`
	IsDelete   uint8  `gorm:"not null;default:0;comment:'是否删除: 0=否, 1=是'"`
	CreateTime int64  `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime int64  `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
	DeleteTime int64  `gorm:"not null;default:0;comment:'删除时间'"`
}
