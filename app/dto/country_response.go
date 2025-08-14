package dto

// 国际列表
type CountryListResponse struct {
	Id      int    `json:"id"`
	Code    string `json:"code"`
	NameEn  string `json:"nameEn"`
	NameZh  string `json:"nameZh"`
	Symbol  string `json:"symbol"`
	IsOpen  string `json:"isOpen"`
	Country string `json:"country"`
}

// 国家详情
type CountryDetailResponse struct {
	Id      int    `json:"id"`
	Code    string `json:"code"`
	NameEn  string `json:"nameEn"`
	NameZh  string `json:"nameZh"`
	Symbol  string `json:"symbol"`
	IsOpen  string `json:"isOpen"`
	Country string `json:"country"`
}
