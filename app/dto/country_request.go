package dto

// 保存国家
type SaveCountry struct {
	Id      int    `json:"id"`
	Code    string `json:"code"`
	NameEn  string `json:"nameEn"`
	NameZh  string `json:"nameZh"`
	Symbol  string `json:"symbol"`
	IsOpen  *int   `json:"isOpen"`
	Country string `json:"country"`
}

// 国家列表
type CountryListRequest struct {
	PageRequest
	Code    string `query:"code" form:"code"`
	NameEn  string `query:"nameEn" form:"nameEn"`
	NameZh  string `query:"nameZh" form:"nameZh"`
	Symbol  string `query:"symbol" form:"symbol"`
	Country string `query:"country" form:"country"`
	IsOpen  string `query:"isOpen" form:"isOpen"`
}

// 新增国家
type CreateCountryRequest struct {
	Code    string `json:"code"`
	NameEn  string `json:"nameEn"`
	NameZh  string `json:"nameZh"`
	Symbol  string `json:"symbol"`
	IsOpen  string `json:"isOpen"`
	Country string `json:"country"`
}

// 更新国家
type UpdateCountryRequest struct {
	Id      int    `json:"id"`
	Code    string `json:"code"`
	NameEn  string `json:"nameEn"`
	NameZh  string `json:"nameZh"`
	Symbol  string `json:"symbol"`
	IsOpen  string `json:"isOpen"`
	Country string `json:"country"`
}

// 根据状态查询国家
type QueryCountryByStatusRequest struct {
	Status string `query:"status" form:"status"`
}
