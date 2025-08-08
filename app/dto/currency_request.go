package dto

// SaveCurrency 保存币种
type SaveCurrency struct {
	CurrencyId      int    `json:"currencyId" binding:"required"`
	PId             int    `json:"pId"`
	Currency        string `json:"currency"`
	Symbol          string `json:"symbol"`
	Logo            string `json:"logo"`
	ContractAddress string `json:"contractAddress"`
	ChainSymbol     string `json:"chainSymbol"`
	CurrencyType    string `json:"currencyType"`
	Protocol        string `json:"protocol"`
	Decimals        int    `json:"decimals"`
	Status          string `json:"status"`
	CreateBy        string `json:"createBy"`
	UpdateBy        string `json:"updateBy"`
	Remark          string `json:"remark"`
}

// CurrencyListRequest 币种列表
type CurrencyListRequest struct {
	PageRequest
	CurrencyName string `query:"currency" form:"currency"`
	Status       string `query:"status" form:"status"`
}

// CreateCurrencyRequest 新增币种
type CreateCurrencyRequest struct {
	PId             int    `json:"pId"`
	Currency        string `json:"currency" binding:"required"`
	Symbol          string `json:"symbol" binding:"required"`
	Logo            string `json:"logo"`
	ContractAddress string `json:"contractAddress"`
	ChainSymbol     string `json:"chainSymbol"`
	CurrencyType    string `json:"currencyType" binding:"required"`
	Protocol        string `json:"protocol" binding:"required"`
	Decimals        int    `json:"decimals"`
	Status          string `json:"status"`
	Remark          string `json:"remark"`
}

// UpdateCurrencyRequest 更新币种
type UpdateCurrencyRequest struct {
	CurrencyId      int    `json:"currencyId" binding:"required"`
	PId             int    `json:"pId" binding:"required"`
	Currency        string `json:"currency" binding:"required"`
	Symbol          string `json:"symbol" binding:"required"`
	Logo            string `json:"logo" binding:"required"`
	ContractAddress string `json:"contractAddress" binding:"required"`
	CurrencyType    string `json:"currencyType" binding:"required"`
	Protocol        string `json:"protocol" binding:"required"`
	Decimals        int    `json:"decimals"`
	Status          string `json:"status"`
	Remark          string `json:"remark"`
}
