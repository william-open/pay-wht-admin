package model

type SAuthDept struct {
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

func (SAuthDept) TableName() string {
	return "w_system_auth_dept"
}
