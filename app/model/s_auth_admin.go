package model

type SAuthAdmin struct {
	ID              uint   `gorm:"primarykey;comment:'主键'"`
	DeptId          uint   `gorm:"not null;default:0;comment:'部门ID'"`
	MId             uint   `gorm:"comment:'商户ID'"`
	PostId          uint   `gorm:"not null;default:0;comment:'岗位ID'"`
	Username        string `gorm:"not null;default:'';comment:'用户账号''"`
	Nickname        string `gorm:"not null;default:'';comment:'用户昵称'"`
	Password        string `gorm:"not null;default:'';comment:'用户密码'"`
	Avatar          string `gorm:"not null;default:'';comment:'用户头像'"`
	AppId           string `gorm:"not null;default:'';comment:'应用ID'"`
	Role            string `gorm:"not null;default:'';comment:'角色主键'"`
	Salt            string `gorm:"not null;default:'';comment:'加密盐巴'"`
	Sort            uint16 `gorm:"not null;default:0;comment:'排序编号'"`
	IsMultipoint    uint8  `gorm:"not null;default:0;comment:'多端登录: 0=否, 1=是''"`
	IsDisable       uint8  `gorm:"not null;default:0;comment:'是否禁用: 0=否, 1=是'"`
	IsDelete        uint8  `gorm:"not null;default:0;comment:'是否删除: 0=否, 1=是'"`
	LastLoginIp     string `gorm:"not null;default:'';comment:'最后登录IP'"`
	LastLoginTime   int64  `gorm:"not null;default:0;comment:'最后登录时间'"`
	CreateTime      int64  `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime      int64  `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
	DeleteTime      int64  `gorm:"not null;default:0;comment:'删除时间'"`
	PayPassword     string `gorm:"not null;default:'';comment:'用户支付密码'"`
	PaySalt         string `gorm:"not null;default:'';comment:'用户支付盐值'"`
	ApiKey          string `gorm:"not null;default:'';comment:'APIKey'"`
	GoogleSecret    string `gorm:"comment:'谷歌密钥'"`
	IsGoogleEnabled uint8  `gorm:"not null;default:0;comment:'是否开启谷歌验证码: 0=否, 1=是'"`
}

func (SAuthAdmin) TableName() string {
	return "w_system_auth_admin"
}
