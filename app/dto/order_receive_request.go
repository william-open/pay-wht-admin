package dto

// 代收列表
type OrderReceiveListRequest struct {
	PageRequest
	Keyword    string `query:"keyword" form:"keyword"`
	ChannelId  int64  `query:"channelId" form:"channelId"`
	UpChanelId int64  `query:"upChanelId" form:"upChanelId"`
	Status     string `query:"status" form:"status"`
	Currency   string `query:"currency" form:"currency"`
	YearMonth  string `query:"yearMonth" form:"yearMonth"`
}

// 更新代收订单
type UpdateOrderReceiveRequest struct {
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
	PayAPI         string  `json:"payApi"`
	PayQueryAPI    string  `json:"payQueryApi"`
	PayoutAPI      string  `json:"payoutApi"`
	PayoutQueryAPI string  `json:"payoutQueryApi"`
	BalanceInquiry string  `json:"balanceInquiry"`
	SendingAddress string  `json:"sendingAddress"`
	Supplementary  string  `json:"supplementary"`
	Documentation  string  `json:"documentation"`
	NeedQuery      int8    `json:"needQuery"`
	CallbackDomain string  `json:"callbackDomain"`
	Remark         string  `json:"remark"`
	Currency       string  `json:"currency"`
	ChannelCode    string  `json:"channelCode"`
	Md5Key         string  `json:"md5Key"`
	RsaPrivateKey  string  `json:"rsaPrivateKey"`
	RsaPublicKey   string  `json:"rsaPublicKey"`
	IpWhiteList    string  `json:"ipWhiteList"`
}
