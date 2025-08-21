package dto

import "time"

// 代收列表
type OrderReceiveListResponse struct {
	OrderID        uint64    `json:"orderId"`        // 全局唯一订单ID
	MID            uint64    `json:"mId"`            // 商户ID
	SupplierID     int64     `json:"supplierId"`     // 上游供应商ID
	MOrderId       string    `json:"mOrderId"`       // 商户订单号
	Amount         float64   `json:"amount"`         // 订单金额
	Fees           float64   `json:"fees"`           // 手续费
	PayAmount      float64   `json:"payAmount"`      // 实际支付金额
	RealMoney      float64   `json:"realMoney"`      // 实际到账金额
	FreezeAmount   float64   `json:"freezeAmount"`   // 冻结金额
	Currency       string    `json:"currency"`       // 货币代码
	NotifyURL      string    `json:"notifyUrl"`      // 异步回调通知URL
	ReturnURL      string    `json:"returnUrl"`      // 同步回调URL
	MDomain        string    `json:"mDomain"`        // 下单域名
	MIP            string    `json:"mIp"`            // 下单IP
	Title          string    `json:"title"`          // 订单标题
	MTitle         string    `json:"mTitle"`         // 商户名称
	ChannelCode    string    `json:"channelCode"`    // 通道编码
	ChannelTitle   string    `json:"channelTitle"`   // 通道名称
	UpChannelCode  string    `json:"upChannelCode"`  // 上游通道编码
	UpChannelTitle string    `json:"upChannelTitle"` // 上游通道名称
	MRate          string    `json:"mRate"`          // 商户费率
	UpRate         string    `json:"upRate"`         // 上游商户费率
	UpFixedFee     string    `json:"upFixedFee"`     // 上游通道固定费用
	MFixedFee      string    `json:"mFixedFee"`      // 商户通道固定费用
	Country        string    `json:"country"`        // 国家
	AccountNo      string    `json:"accountNo"`      // 付款人账号
	AccountName    string    `json:"accountName"`    // 付款人姓名
	PayEmail       string    `json:"payEmail"`       // 付款人邮箱
	PayPhone       string    `json:"payPhone"`       // 付款人手机号码
	BankCode       string    `json:"bankCode"`       // 付款人银行编码
	BankName       string    `json:"bankName"`       // 付款人银行名
	Status         int8      `json:"status"`         // 0:待支付,1:成功,2:失败,3:退款
	NotifyStatus   int8      `json:"notifyStatus"`   // 回调通知状态:0表示未回调，1表示回调成功，2回调失败
	UpOrderID      *uint64   `json:"upOrderId"`      // 上游交易订单ID
	ChannelID      int64     `json:"channelId"`      // 系统支付渠道ID
	UpChannelID    int64     `json:"upChannelId"`    // 上游通道ID
	NotifyTime     time.Time `json:"notifyTime"`     // 回调通知时间
	CreateTime     time.Time `json:"createTime"`     // 创建时间
	UpdateTime     time.Time `json:"updateTime"`     // 更新时间
}

// 代收详情
type OrderReceiveDetailResponse struct {
	OrderID      uint64    `json:"orderId"`      // 全局唯一订单ID
	MID          uint64    `json:"mId"`          // 商户ID
	SupplierID   int64     `json:"supplierId"`   // 上游供应商ID
	Amount       float64   `json:"amount"`       // 订单金额
	Fees         float64   `json:"fees"`         // 手续费
	PayAmount    float64   `json:"payAmount"`    // 实际支付金额
	RealMoney    float64   `json:"realMoney"`    // 实际到账金额
	FreezeAmount float64   `json:"freezeAmount"` // 冻结金额
	Currency     string    `json:"currency"`     // 货币代码
	NotifyURL    string    `json:"notifyUrl"`    // 异步回调通知URL
	ReturnURL    string    `json:"returnUrl"`    // 同步回调URL
	MDomain      string    `json:"mDomain"`      // 下单域名
	MIP          string    `json:"mIp"`          // 下单IP
	Title        string    `json:"title"`        // 订单标题
	AccountNo    string    `json:"accountNo"`    // 付款人账号
	AccountName  string    `json:"accountName"`  // 付款人姓名
	PayEmail     string    `json:"payEmail"`     // 付款人邮箱
	PayPhone     string    `json:"payPhone"`     // 付款人手机号码
	BankCode     string    `json:"bankCode"`     // 付款人银行编码
	BankName     string    `json:"bankName"`     // 付款人银行名
	Status       int8      `json:"status"`       // 0:待支付,1:成功,2:失败,3:退款
	UpOrderID    *uint64   `json:"upOrderId"`    // 上游交易订单ID
	ChannelID    int64     `json:"channelId"`    // 系统支付渠道ID
	UpChannelID  int64     `json:"upChannelId"`  // 上游通道ID
	NotifyTime   time.Time `json:"notifyTime"`   // 回调通知时间
	CreateTime   time.Time `json:"createTime"`   // 创建时间
	UpdateTime   time.Time `json:"updateTime"`   // 更新时间
}
