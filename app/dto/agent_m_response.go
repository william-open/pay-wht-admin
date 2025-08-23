package dto

import (
	"wht-admin/framework/datetime"
)

// AgentMListResponse 代理商户列表
type AgentMListResponse struct {
	Id            int64             `json:"id"`
	MId           int64             `json:"mId"`
	AId           int64             `json:"aId"`
	Currency      string            `json:"currency"`
	SysChannelId  int64             `json:"sysChannelId"`
	DefaultRate   string            `json:"defaultRate"`
	SingleFee     string            `json:"singleFee"`
	Status        string            `json:"status"`
	MerchantTitle string            `json:"merchantTitle"`
	AgentTitle    string            `json:"agentTitle"`
	ChannelTitle  string            `json:"channelTitle"`
	Coding        string            `json:"coding"`
	CreateBy      string            `json:"createBy"`
	CreateTime    datetime.Datetime `json:"createTime"`
}

// AgentMDetailResponse 代理商户详情
type AgentMDetailResponse struct {
	Id           int64             `json:"id"`
	MId          int64             `json:"mId"`
	AId          int64             `json:"aId"`
	Currency     string            `json:"currency"`
	SysChannelId int64             `json:"sysChannelId"`
	DefaultRate  string            `json:"defaultRate"`
	SingleFee    string            `json:"singleFee"`
	Status       string            `json:"status"`
	CreateBy     string            `json:"createBy"`
	CreateTime   datetime.Datetime `json:"createTime"`
}
