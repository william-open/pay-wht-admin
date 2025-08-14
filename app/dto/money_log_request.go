package dto

// SaveMoneyLog 保存资金日志
type SaveMoneyLog struct {
	ID          uint    `json:"id"`
	UID         int     `json:"uId"`
	Money       float64 `json:"money"`
	OrderNo     string  `json:"orderNo"`
	Type        int8    `json:"type"`
	Operator    string  `json:"operator"`
	Currency    string  `json:"currency"`
	Description string  `json:"description"`
	OldBalance  float64 `json:"oldBalance"`
	Balance     float64 `json:"balance"`
	CreateTime  string  `json:"createTime"`
}

// MoneyLogListRequest 资金日志列表
type MoneyLogListRequest struct {
	PageRequest
	Nickname string `query:"nickname" form:"nickname"`
	Currency string `query:"currency" form:"currency"`
	UserType int8   `query:"userType" form:"userType"`
}

// CreateMoneyLogRequest 新增资金日志
type CreateMoneyLogRequest struct {
	UID         int     `json:"uId"`
	Money       float64 `json:"money"`
	OrderNo     string  `json:"orderNo"`
	Type        int8    `json:"type"`
	Operator    string  `json:"operator"`
	Currency    string  `json:"currency"`
	Description string  `json:"description"`
	OldBalance  float64 `json:"oldBalance"`
	Balance     float64 `json:"balance"`
}

// UpdateMoneyLogRequest 更新资金日志
type UpdateMoneyLogRequest struct {
	UID         int     `json:"uId" binding:"required"`
	Money       float64 `json:"money"`
	OrderNo     string  `json:"orderNo"`
	Type        int8    `json:"type"`
	Operator    string  `json:"operator"`
	Currency    string  `json:"currency"`
	Description string  `json:"description"`
	OldBalance  float64 `json:"oldBalance"`
	Balance     float64 `json:"balance"`
}
