package model

// WPayUpstream 支付上游表
type WPayUpstream struct {
	ID             int     `gorm:"primaryKey;column:id;autoIncrement" json:"id"`                                    // 主键ID
	Title          string  `gorm:"column:title;size:40;not null" json:"title"`                                      // 名称
	Type           int8    `gorm:"column:type;not null" json:"type"`                                                // 1:代收 2:代付 3:都有
	WayID          string  `gorm:"column:way_id;type:json;not null" json:"wayId"`                                   // 对应通道(JSON)
	Account        string  `gorm:"column:account;size:100;not null" json:"account"`                                 // 商户账号
	PayKey         string  `gorm:"column:pay_key;type:text;not null" json:"payKey"`                                 // 密钥
	ReceivingKey   string  `gorm:"column:receiving_key;type:text;not null" json:"receivingKey"`                     // 代付密钥
	SuccessRate    float64 `gorm:"column:success_rate;type:double(5,2);not null;default:100.00" json:"successRate"` // 成功率
	OrderQuantity  int     `gorm:"column:order_quantity;not null;default:0" json:"orderQuantity"`                   // 总的订单数
	Rate           float64 `gorm:"column:rate;type:double(4,2);not null" json:"rate"`                               // 默认费率
	AppID          string  `gorm:"column:app_id;size:20" json:"appId"`                                              // appid
	AppSecret      string  `gorm:"column:app_secret;type:text;not null" json:"appSecret"`                           // 安全密钥
	UpdateTime     *int    `gorm:"column:update_time" json:"updateTime"`                                            // 上次更改时间
	ControlStatus  int8    `gorm:"column:control_status;not null;default:0" json:"controlStatus"`                   // 风控状态:0否1是
	Sort           int     `gorm:"column:sort;not null;default:0" json:"sort"`                                      // 排序
	PayingMoney    float64 `gorm:"column:paying_money;type:decimal(16,2);not null;default:0.00" json:"payingMoney"` // 当天交易金额
	MinMoney       float64 `gorm:"column:min_money;type:decimal(15,2);not null;default:0.00" json:"minMoney"`       // 单笔最小交易额
	MaxMoney       float64 `gorm:"column:max_money;type:decimal(15,2);not null;default:0.00" json:"maxMoney"`       // 单笔最大交易额
	Status         int8    `gorm:"column:status;not null;default:0" json:"status"`                                  // 状态 0:关闭;1:开启
	PayStatus      int8    `gorm:"column:pay_status;not null;default:1" json:"payStatus"`                           // 状态 0:关闭;1:开启;2:;3:系统错误
	OutStatus      int8    `gorm:"column:out_status;not null;default:1" json:"outStatus"`                           // 状态 0:关闭;1:开启;2:;3:系统错误
	PayAPI         string  `gorm:"column:pay_api;size:100" json:"payApi"`                                           // 代收下单API
	PayQueryAPI    string  `gorm:"column:pay_query_api;size:100" json:"payQueryApi"`                                // 代收查询地址
	PayoutAPI      string  `gorm:"column:payout_api;size:100" json:"payoutApi"`                                     // 代付下单API
	PayoutQueryAPI string  `gorm:"column:payout_query_api;size:100" json:"payoutQueryApi"`                          // 代付查询地址
	BalanceInquiry string  `gorm:"column:balance_inquiry;size:100" json:"balanceInquiry"`                           // 余额查询地址
	SendingAddress string  `gorm:"column:sending_address;size:100" json:"sendingAddress"`                           // 下发地址
	Supplementary  string  `gorm:"column:supplementary;size:100" json:"supplementary"`                              // 补单地址
	Documentation  string  `gorm:"column:documentation;size:100" json:"documentation"`                              // 文档地址
	NeedQuery      int8    `gorm:"column:need_query;not null;default:1" json:"needQuery"`                           // 确定时是否需要查询
	IPWhiteList    string  `gorm:"column:ip_white_list;size:200" json:"ipWhiteList"`                                // 回调IP白名单
	CallbackDomain string  `gorm:"column:callback_domain;size:50" json:"callbackDomain"`                            // 回调访问的域名
	Remark         string  `gorm:"column:remark;size:50" json:"remark"`                                             // 备注
	Currency       string  `gorm:"column:currency;size:50" json:"currency"`                                         // 国家货币符号
	ChannelCode    string  `gorm:"column:channel_code;size:50" json:"channelCode"`                                  // 通道对接编码
	Md5Key         string  `gorm:"column:md5_key;size:50" json:"md5Key"`                                            // md5密钥
	RsaPrivateKey  string  `gorm:"column:rsa_private_key;size:500" json:"rsaPrivateKey"`                            // RSA私钥
	RsaPublicKey   string  `gorm:"column:rsa_public_key;size:500" json:"rsaPublicKey"`                              // RSA公钥
}

// TableName 自定义表名
func (WPayUpstream) TableName() string {
	return "w_pay_upstream"
}
