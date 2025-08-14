package model

type WCountry struct {
	Id      int `gorm:"primaryKey;autoIncrement"`
	Code    string
	NameEn  string
	NameZh  string
	Symbol  string
	Country string
	IsOpen  *int `gorm:"default:0"`
}

func (WCountry) TableName() string {
	return "w_currency_code"
}
