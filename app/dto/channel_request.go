package dto

// 保存通道
type SaveChannel struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Status      int    `json:"status"`
	Max         string `json:"max"`
	Min         string `json:"min"`
	DefaultRate int    `json:"defaultRate"`
	Coding      string `json:"coding"`
	AddRate     int    `json:"addRate"`
	Type        int    `json:"type"`
	Charge      int    `json:"charge"`
	Currency    string `json:"currency"`
	Remark      string `json:"remark"`
}

// 通道列表
type ChannelListRequest struct {
	PageRequest
	Keyword  string `query:"keyword" form:"keyword"`
	Status   string `json:"status" form:"status"`
	Type     int    `json:"type" form:"type"`
	Currency string `json:"currency" form:"currency"`
}

// 新增通道
type CreateChannelRequest struct {
	Title       string `json:"title"`
	Status      int    `json:"status"`
	Max         string `json:"max"`
	Min         string `json:"min"`
	DefaultRate int    `json:"default_rate"`
	Coding      string `json:"coding"`
	AddRate     int    `json:"add_rate"`
	Type        int    `json:"type"`
	Charge      int    `json:"charge"`
	Currency    string `json:"currency"`
	Remark      string `json:"remark"`
}

// 更新通道
type UpdateChannelRequest struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Status      int    `json:"status"`
	Max         string `json:"max"`
	Min         string `json:"min"`
	DefaultRate int    `json:"default_rate"`
	Coding      string `json:"coding"`
	AddRate     int    `json:"add_rate"`
	Type        int    `json:"type"`
	Charge      int    `json:"charge"`
	Currency    string `json:"currency"`
	Remark      string `json:"remark"`
}

// 根据状态查询通道
type QueryChannelByStatusRequest struct {
	Status   string `query:"status" form:"status"`
	Currency string `query:"currency" form:"currency"`
}
