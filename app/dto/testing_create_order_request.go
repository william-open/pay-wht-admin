package dto

// 代收请求参数
type TestCreateOrderReq struct {
	Version      string `json:"version" binding:"required"`         //接口版本
	MerchantNo   string `json:"merchant_no" binding:"required"`     //商户号
	TranFlow     string `json:"tran_flow" binding:"required"`       //订单号
	TranDatetime string `json:"tran_datetime" binding:"required"`   //13位时间戳
	Amount       string `json:"amount" binding:"required"`          //订单金额
	PayType      string `json:"pay_type" binding:"required"`        //通道编码
	NotifyUrl    string `json:"notify_url" binding:"omitempty,url"` //回调地址
	RedirectUrl  string `json:"redirect_url"`                       //成功跳转地址
	ProductInfo  string `json:"product_info" binding:"required"`    //订单标题/内容
	AccNo        string `json:"acc_no"`                             // 付款人账号
	AccName      string `json:"acc_name"`                           //付款人姓名
	PayEmail     string `json:"pay_email"`                          //付款人邮箱
	PayPhone     string `json:"pay_phone"`                          //付款人手机号
	BankCode     string `json:"bank_code"`                          //付款人银行名
	BankName     string `json:"bank_name"`                          //付款人银行名
	Sign         string `json:"sign" binding:"required"`            //MD5 签名 32大写
}
