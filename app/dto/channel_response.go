package dto

import "ruoyi-go/framework/datetime"

// 通道列表
type ChannelListResponse struct {
	Id          int               `json:"id"`
	Title       string            `json:"title"`
	Status      string            `json:"status"`
	Max         string            `json:"max"`
	Min         string            `json:"min"`
	DefaultRate string            `json:"defaultRate"`
	Coding      string            `json:"coding"`
	AddRate     string            `json:"addRate"`
	Type        int               `json:"type"`
	Charge      int               `json:"charge"`
	Currency    string            `json:"currency"`
	Remark      string            `json:"remark"`
	Country     string            `json:"country"`
	CreateBy    string            `json:"createBy"`
	CreateTime  datetime.Datetime `json:"createTime"`
	UpdateBy    string            `json:"updateBy"`
	UpdateTime  datetime.Datetime `json:"updateTime"`
}

// 通道详情
type ChannelDetailResponse struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Status      string `json:"status"`
	Max         string `json:"max"`
	Min         string `json:"min"`
	DefaultRate string `json:"default_rate"`
	Coding      string `json:"coding"`
	AddRate     string `json:"add_rate"`
	Type        int    `json:"type"`
	Charge      int    `json:"charge"`
	Currency    string `json:"currency"`
	Remark      string `json:"remark"`
}

// 下拉列表数据列表
type DropDownListResponse struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Coding string `json:"coding"`
}
