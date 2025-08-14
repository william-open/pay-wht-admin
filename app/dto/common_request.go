package dto

// 保存状态
type SaveStatus struct {
	Id     int  `json:"id"`
	Status int8 `json:"status"`
}

// 更改状态
type UpdateStatusRequest struct {
	Id     int  `json:"id"`
	Status int8 `json:"status"`
}
