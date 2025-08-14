package dto

// 通道供应商列表
type PayUpstreamListResponse struct {
	Id             int     `json:"id"`
	Title          string  `json:"title"`          // 名称
	Type           int8    `json:"type"`           // 1:代收 2:代付 3:都有
	WayID          string  `json:"wayID"`          // 对应通道(JSON)
	Account        string  `json:"account"`        // 商户账号
	PayKey         string  `json:"payKey"`         // 密钥
	ReceivingKey   string  `son:"receivingKey"`    // 代付密钥
	SuccessRate    float64 `json:"successRate"`    // 成功率
	OrderQuantity  int     `json:"orderQuantity"`  // 总的订单数
	Rate           float64 `json:"rate"`           // 默认费率
	AppID          string  `json:"appID"`          // appid
	AppSecret      string  `json:"appSecret"`      // 安全密钥
	UpdateTime     *int    `json:"updateTime"`     // 上次更改时间
	ControlStatus  int8    `json:"controlStatus"`  // 风控状态:0否1是
	Sort           int     `json:"sort"`           // 排序
	PayingMoney    float64 `json:"payingMoney"`    // 当天交易金额
	MinMoney       float64 `json:"minMoney"`       // 单笔最小交易额
	MaxMoney       float64 `json:"maxMoney"`       // 单笔最大交易额
	Status         int8    `json:"status"`         // 状态 0:关闭;1:开启
	PayStatus      int8    `json:"payStatus"`      // 状态 0:关闭;1:开启;2:;3:系统错误
	OutStatus      int8    `json:"outStatus"`      // 状态 0:关闭;1:开启;2:;3:系统错误
	PayAPI         string  `json:"payAPI"`         // 代收下单API
	PayQueryAPI    string  `json:"payQueryAPI"`    // 代收查询地址
	PayoutAPI      string  `json:"payoutAPI"`      // 代付下单API
	PayoutQueryAPI string  `json:"payoutQueryAPI"` // 代付查询地址
	BalanceInquiry string  `json:"balanceInquiry"` // 余额查询地址
	SendingAddress string  `json:"sendingAddress"` // 下发地址
	Supplementary  string  `json:"supplementary"`  // 补单地址
	Documentation  string  `json:"documentation"`  // 文档地址
	NeedQuery      int8    `json:"needQuery"`      // 确定时是否需要查询
	IPWhiteList    string  `json:"iPWhiteList"`    // 回调IP白名单
	CallbackDomain string  `json:"callbackDomain"` // 回调访问的域名
	Currency       string  `json:"currency"`       // 国家货币符号
	Remark         string  `json:"remark"`         // 备注
}

// 通道供应商详情
type PayUpstreamDetailResponse struct {
	Id             int     `json:"id"`
	Title          string  `json:"title"`          // 名称
	Type           int8    `json:"type"`           // 1:代收 2:代付 3:都有
	WayID          string  `json:"wayID"`          // 对应通道(JSON)
	Account        string  `json:"account"`        // 商户账号
	PayKey         string  `json:"payKey"`         // 密钥
	ReceivingKey   string  `son:"receivingKey"`    // 代付密钥
	SuccessRate    float64 `json:"successRate"`    // 成功率
	OrderQuantity  int     `json:"orderQuantity"`  // 总的订单数
	Rate           float64 `json:"rate"`           // 默认费率
	AppID          string  `json:"appID"`          // appid
	AppSecret      string  `json:"appSecret"`      // 安全密钥
	UpdateTime     *int    `json:"updateTime"`     // 上次更改时间
	ControlStatus  int8    `json:"controlStatus"`  // 风控状态:0否1是
	Sort           int     `json:"sort"`           // 排序
	PayingMoney    float64 `json:"payingMoney"`    // 当天交易金额
	MinMoney       float64 `json:"minMoney"`       // 单笔最小交易额
	MaxMoney       float64 `json:"maxMoney"`       // 单笔最大交易额
	Status         int8    `json:"status"`         // 状态 0:关闭;1:开启
	PayStatus      int8    `json:"payStatus"`      // 状态 0:关闭;1:开启;2:;3:系统错误
	OutStatus      int8    `json:"outStatus"`      // 状态 0:关闭;1:开启;2:;3:系统错误
	PayAPI         string  `json:"payAPI"`         // 代收下单API
	PayQueryAPI    string  `json:"payQueryAPI"`    // 代收查询地址
	PayoutAPI      string  `json:"payoutAPI"`      // 代付下单API
	PayoutQueryAPI string  `json:"payoutQueryAPI"` // 代付查询地址
	BalanceInquiry string  `json:"balanceInquiry"` // 余额查询地址
	SendingAddress string  `json:"sendingAddress"` // 下发地址
	Supplementary  string  `json:"supplementary"`  // 补单地址
	Documentation  string  `json:"documentation"`  // 文档地址
	NeedQuery      int8    `json:"needQuery"`      // 确定时是否需要查询
	IPWhiteList    string  `json:"iPWhiteList"`    // 回调IP白名单
	CallbackDomain string  `json:"callbackDomain"` // 回调访问的域名
	Currency       string  `json:"currency"`       // 国家货币符号
	Remark         string  `json:"remark"`         // 备注
}

// UpstreamDropDownListResponse 通道供应商列表根据状态返回下拉列表
type UpstreamDropDownListResponse struct {
	Id    int    `json:"id"`
	Title string `json:"title"` // 名称
}
