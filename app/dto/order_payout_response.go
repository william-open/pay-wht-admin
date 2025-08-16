package dto

import "time"

// 代付列表
type OrderPayoutListResponse struct {
	OrderID      uint64    `json:"order_id"`      // 全局唯一订单ID
	MID          uint64    `json:"m_id"`          // 商户ID
	SupplierID   int64     `json:"supplier_id"`   // 上游供应商ID
	MOrderID     string    `json:"m_order_id"`    // 商户订单号
	Amount       float64   `json:"amount"`        // 订单金额
	Fees         float64   `json:"fees"`          // 手续费
	PayAmount    float64   `json:"pay_amount"`    // 实际支付金额
	RealMoney    float64   `json:"real_money"`    // 实际到账金额
	FreezeAmount float64   `json:"freeze_amount"` // 冻结金额
	Currency     string    `json:"currency"`      // 货币代码
	NotifyURL    string    `json:"notify_url"`    // 异步回调通知URL
	ReturnURL    string    `json:"return_url"`    // 同步回调URL
	MDomain      string    `json:"m_domain"`      // 下单域名
	MIP          string    `json:"m_ip"`          // 下单IP
	Title        string    `json:"title"`         // 订单标题
	AccountNo    string    `json:"account_no"`    // 付款人账号
	AccountName  string    `json:"account_name"`  // 付款人姓名
	PayEmail     string    `json:"pay_email"`     // 付款人邮箱
	PayPhone     string    `json:"pay_phone"`     // 付款人手机号码
	BankCode     string    `json:"bank_code"`     // 付款人银行编码
	BankName     string    `json:"bank_name"`     // 付款人银行名
	Status       int8      `json:"status"`        // 0:待支付,1:成功,2:失败,3:退款
	UpOrderID    *uint64   `json:"up_order_id"`   // 上游交易订单ID
	ChannelID    int64     `json:"channel_id"`    // 系统支付渠道ID
	UpChannelID  int64     `json:"up_channel_id"` // 上游通道ID
	NotifyTime   time.Time `json:"notify_time"`   // 回调通知时间
	CreateTime   time.Time `json:"create_time"`   // 创建时间
	UpdateTime   time.Time `json:"update_time"`   // 更新时间

}

// 代付详情
type OrderPayoutDetailResponse struct {
	OrderID      uint64    `json:"order_id"`      // 全局唯一订单ID
	MID          uint64    `json:"m_id"`          // 商户ID
	SupplierID   int64     `json:"supplier_id"`   // 上游供应商ID
	MOrderID     string    `json:"m_order_id"`    // 商户订单号
	Amount       float64   `json:"amount"`        // 订单金额
	Fees         float64   `json:"fees"`          // 手续费
	PayAmount    float64   `json:"pay_amount"`    // 实际支付金额
	RealMoney    float64   `json:"real_money"`    // 实际到账金额
	FreezeAmount float64   `json:"freeze_amount"` // 冻结金额
	Currency     string    `json:"currency"`      // 货币代码
	NotifyURL    string    `json:"notify_url"`    // 异步回调通知URL
	ReturnURL    string    `json:"return_url"`    // 同步回调URL
	MDomain      string    `json:"m_domain"`      // 下单域名
	MIP          string    `json:"m_ip"`          // 下单IP
	Title        string    `json:"title"`         // 订单标题
	AccountNo    string    `json:"account_no"`    // 付款人账号
	AccountName  string    `json:"account_name"`  // 付款人姓名
	PayEmail     string    `json:"pay_email"`     // 付款人邮箱
	PayPhone     string    `json:"pay_phone"`     // 付款人手机号码
	BankCode     string    `json:"bank_code"`     // 付款人银行编码
	BankName     string    `json:"bank_name"`     // 付款人银行名
	Status       int8      `json:"status"`        // 0:待支付,1:成功,2:失败,3:退款
	UpOrderID    *uint64   `json:"up_order_id"`   // 上游交易订单ID
	ChannelID    int64     `json:"channel_id"`    // 系统支付渠道ID
	UpChannelID  int64     `json:"up_channel_id"` // 上游通道ID
	NotifyTime   time.Time `json:"notify_time"`   // 回调通知时间
	CreateTime   time.Time `json:"create_time"`   // 创建时间
	UpdateTime   time.Time `json:"update_time"`   // 更新时间
}
