package dto

// 保存通道供应商
type SavePayUpstream struct {
	Id             int     `json:"id"`
	Title          string  `json:"title"`
	Type           int8    `json:"type"`
	Status         int8    `json:"status"`
	WayId          string  `json:"wayId"`
	Account        string  `json:"account"`
	PayKey         string  `json:"payKey"`
	ReceivingKey   string  `json:"receivingKey"`
	SuccessRate    float64 `json:"successRate"`
	OrderQuantity  int     `json:"orderQuantity"`
	Rate           float64 `json:"rate"`
	AppID          string  `json:"appID"`
	AppSecret      string  `json:"appSecret"`
	UpdateTime     *int    `json:"updateTime"`
	ControlStatus  int8    `json:"controlStatus"`
	Sort           int     `json:"sort"`
	PayingMoney    float64 `json:"payingMoney"`
	MinMoney       float64 `json:"minMoney"`
	MaxMoney       float64 `json:"maxMoney"`
	PayStatus      int8    `json:"payStatus"`
	OutStatus      int8    `json:"outStatus"`
	PayAPI         string  `json:"payAPI"`
	PayQueryAPI    string  `json:"payQueryAPI"`
	PayoutAPI      string  `json:"payoutAPI"`
	PayoutQueryAPI string  `json:"payoutQueryAPI"`
	BalanceInquiry string  `json:"balanceInquiry"`
	SendingAddress string  `json:"sendingAddress"`
	Supplementary  string  `json:"supplementary"`
	Documentation  string  `json:"documentation"`
	NeedQuery      int8    `json:"needQuery"`
	IPWhiteList    string  `json:"iPWhiteList"`
	CallbackDomain string  `json:"callbackDomain"`
	Remark         string  `json:"remark"`
	ChannelCode    string  `json:"channelCode"`
}

// 通道供应商列表
type PayUpstreamListRequest struct {
	PageRequest
	Keyword        string  `json:"keyword"`
	Title          string  `json:"title"`
	Type           int     `json:"type"`
	Status         string  `json:"status"`
	WayId          string  `json:"wayId"`
	Account        string  `json:"account"`
	PayKey         string  `json:"payKey"`
	ReceivingKey   int     `json:"receivingKey"`
	SuccessRate    float64 `json:"successRate"`
	OrderQuantity  int     `json:"orderQuantity"`
	Rate           float64 `json:"rate"`
	AppID          string  `json:"appID"`
	AppSecret      string  `json:"appSecret"`
	UpdateTime     *int    `json:"updateTime"`
	ControlStatus  int8    `json:"controlStatus"`
	Sort           int     `json:"sort"`
	PayingMoney    float64 `json:"payingMoney"`
	MinMoney       float64 `json:"minMoney"`
	MaxMoney       float64 `json:"maxMoney"`
	PayStatus      int8    `json:"payStatus"`
	OutStatus      int8    `json:"outStatus"`
	PayAPI         *string `json:"payAPI"`
	PayQueryAPI    *string `json:"payQueryAPI"`
	PayoutAPI      *string `json:"payoutAPI"`
	PayoutQueryAPI *string `json:"payoutQueryAPI"`
	BalanceInquiry *string `json:"balanceInquiry"`
	SendingAddress *string `json:"sendingAddress"`
	Supplementary  *string `json:"supplementary"`
	Documentation  *string `json:"documentation"`
	NeedQuery      int8    `json:"needQuery"`
	IPWhiteList    *string `json:"iPWhiteList"`
	CallbackDomain *string `json:"callbackDomain"`
	Remark         *string `json:"remark"`
}

// 新增通道供应商
type CreatePayUpstreamRequest struct {
	Title          string  `json:"title"`
	Type           int8    `json:"type"`
	Status         int8    `json:"status"`
	WayId          string  `json:"wayId"`
	Account        string  `json:"account"`
	PayKey         string  `json:"payKey"`
	ReceivingKey   string  `json:"receivingKey"`
	SuccessRate    float64 `json:"successRate"`
	OrderQuantity  int     `json:"orderQuantity"`
	Rate           float64 `json:"rate"`
	AppID          string  `json:"appID"`
	AppSecret      string  `json:"appSecret"`
	UpdateTime     *int    `json:"updateTime"`
	ControlStatus  int8    `json:"controlStatus"`
	Sort           int     `json:"sort"`
	PayingMoney    float64 `json:"payingMoney"`
	MinMoney       float64 `json:"minMoney"`
	MaxMoney       float64 `json:"maxMoney"`
	PayStatus      int8    `json:"payStatus"`
	OutStatus      int8    `json:"outStatus"`
	PayAPI         string  `json:"payAPI"`
	PayQueryAPI    string  `json:"payQueryAPI"`
	PayoutAPI      string  `json:"payoutAPI"`
	PayoutQueryAPI string  `json:"payoutQueryAPI"`
	BalanceInquiry string  `json:"balanceInquiry"`
	SendingAddress string  `json:"sendingAddress"`
	Supplementary  string  `json:"supplementary"`
	Documentation  string  `json:"documentation"`
	NeedQuery      int8    `json:"needQuery"`
	IPWhiteList    string  `json:"iPWhiteList"`
	CallbackDomain string  `json:"callbackDomain"`
	Remark         string  `json:"remark"`
}

// 更新通道供应商
type UpdatePayUpstreamRequest struct {
	Id             int     `json:"id"`
	Title          string  `json:"title"`
	Type           int8    `json:"type"`
	Status         int8    `json:"status"`
	WayId          string  `json:"wayId"`
	Account        string  `json:"account"`
	PayKey         string  `json:"payKey"`
	ReceivingKey   string  `json:"receivingKey"`
	SuccessRate    float64 `json:"successRate"`
	OrderQuantity  int     `json:"orderQuantity"`
	Rate           float64 `json:"rate"`
	AppID          string  `json:"appID"`
	AppSecret      string  `json:"appSecret"`
	UpdateTime     *int    `json:"updateTime"`
	ControlStatus  int8    `json:"controlStatus"`
	Sort           int     `json:"sort"`
	PayingMoney    float64 `json:"payingMoney"`
	MinMoney       float64 `json:"minMoney"`
	MaxMoney       float64 `json:"maxMoney"`
	PayStatus      int8    `json:"payStatus"`
	OutStatus      int8    `json:"outStatus"`
	PayAPI         string  `json:"payAPI"`
	PayQueryAPI    string  `json:"payQueryAPI"`
	PayoutAPI      string  `json:"payoutAPI"`
	PayoutQueryAPI string  `json:"payoutQueryAPI"`
	BalanceInquiry string  `json:"balanceInquiry"`
	SendingAddress string  `json:"sendingAddress"`
	Supplementary  string  `json:"supplementary"`
	Documentation  string  `json:"documentation"`
	NeedQuery      int8    `json:"needQuery"`
	IPWhiteList    string  `json:"iPWhiteList"`
	CallbackDomain string  `json:"callbackDomain"`
	Remark         string  `json:"remark"`
}

// 根据状态查询供应商
type QueryUpstreamByStatusRequest struct {
	Status string `query:"status" form:"status"`
}

// 查询供应商通道列表
type QueryUpstreamChannelRequest struct {
	Status     string `query:"status" form:"status"`
	UpstreamId string `query:"upstreamId" form:"upstreamId"`
	Currency   string `query:"currency" form:"currency"`
}
